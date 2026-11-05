"""
1_finetune_qlora.py
────────────────────────────────────────────────────────────────
QLoRA fine-tuning on CodeAlpaca-20k + CyberNative Cybersecurity dataset.
Target model: mistralai/Mistral-7B-v0.1  (swap freely)

Install dependencies first:
    pip install transformers peft bitsandbytes datasets accelerate trl
────────────────────────────────────────────────────────────────
"""

import torch
from datasets import load_dataset, concatenate_datasets
from transformers import (
    AutoModelForCausalLM,
    AutoTokenizer,
    BitsAndBytesConfig,
    TrainingArguments,
)
from peft import LoraConfig, prepare_model_for_kbit_training
from trl import SFTTrainer

# ─── Configuration ────────────────────────────────────────────────────────────

BASE_MODEL  = "mistralai/Mistral-7B-v0.1"   # change to any HF causal LM
OUTPUT_DIR  = "./qlora-output"
MAX_SEQ_LEN = 1024    # code examples can be long; increase if VRAM allows
EPOCHS      = 3
BATCH_SIZE  = 4       # per-device; lower to 2 if you hit OOM
GRAD_ACCUM  = 4       # effective batch = BATCH_SIZE * GRAD_ACCUM = 16


# ─── 1. Dataset ───────────────────────────────────────────────────────────────

def load_and_format_datasets():
    print("📦 Loading datasets...")

    # --- CodeAlpaca-20k (instruction / input / output columns) ---
    coding = load_dataset("sahil2801/CodeAlpaca-20k", split="train")

    def format_coding(row):
        instruction = row.get("instruction", "").strip()
        ctx         = row.get("input", "").strip()
        output      = row.get("output", "").strip()
        prompt = f"{instruction}\n{ctx}" if ctx else instruction
        return {"text": f"<s>[INST] {prompt} [/INST] {output}</s>"}

    coding_fmt = coding.map(
        format_coding,
        remove_columns=coding.column_names,
        desc="Formatting CodeAlpaca",
    )

    # --- CyberNative Cybersecurity dataset ---
    try:
        vulns = load_dataset(
            "CyberNative-AI/Cybersecurity_Specialized_Dataset", split="train"
        )

        # Inspect available columns and adapt
        print(f"   Cybersecurity columns: {vulns.column_names}")

        def format_vulns(row):
            # Common column names — the function picks whichever exist
            instruction = (
                row.get("instruction")
                or row.get("prompt")
                or row.get("question")
                or ""
            ).strip()
            ctx = (row.get("input") or row.get("context") or "").strip()
            output = (
                row.get("output")
                or row.get("response")
                or row.get("answer")
                or ""
            ).strip()
            prompt = f"{instruction}\n{ctx}" if ctx else instruction
            return {"text": f"<s>[INST] {prompt} [/INST] {output}</s>"}

        vulns_fmt = vulns.map(
            format_vulns,
            remove_columns=vulns.column_names,
            desc="Formatting CyberNative",
        )
    except Exception as e:
        print(f"⚠️  Could not load CyberNative dataset ({e}). Using coding only.")
        vulns_fmt = None

    # --- CVE explanations (secondary vuln source) ---
    try:
        cve = load_dataset("detomo/cve-explain-openai", split="train")
        print(f"   CVE columns: {cve.column_names}")

        def format_cve(row):
            prompt = (row.get("prompt") or row.get("question") or "").strip()
            answer = (
                row.get("completion") or row.get("answer") or row.get("response") or ""
            ).strip()
            return {"text": f"<s>[INST] {prompt} [/INST] {answer}</s>"}

        cve_fmt = cve.map(
            format_cve,
            remove_columns=cve.column_names,
            desc="Formatting CVE",
        )
    except Exception as e:
        print(f"⚠️  Could not load CVE dataset ({e}). Skipping.")
        cve_fmt = None

    # --- Combine ---
    parts = [coding_fmt]
    if vulns_fmt:
        parts.append(vulns_fmt)
    if cve_fmt:
        parts.append(cve_fmt)

    combined = concatenate_datasets(parts).shuffle(seed=42)

    # Drop empty rows
    combined = combined.filter(lambda x: len(x["text"].strip()) > 20)

    print(f"✅ Total training examples: {len(combined):,}")
    print(f"   Sample:\n{combined[0]['text'][:300]}\n")
    return combined


# ─── 2. Model & Tokeniser ─────────────────────────────────────────────────────

def load_model_and_tokenizer():
    print("🔧 Loading model in 4-bit (QLoRA)...")

    bnb_config = BitsAndBytesConfig(
        load_in_4bit=True,
        bnb_4bit_quant_type="nf4",           # NormalFloat4 — optimal for QLoRA
        bnb_4bit_compute_dtype=torch.float16,
        bnb_4bit_use_double_quant=True,      # saves ~0.4 bits per param extra
    )

    model = AutoModelForCausalLM.from_pretrained(
        BASE_MODEL,
        quantization_config=bnb_config,
        device_map="auto",
        trust_remote_code=True,
    )
    model = prepare_model_for_kbit_training(model)  # cast layernorms → fp32

    tokenizer = AutoTokenizer.from_pretrained(BASE_MODEL, trust_remote_code=True)
    tokenizer.pad_token    = tokenizer.eos_token
    tokenizer.padding_side = "right"   # avoid warnings with causal LMs

    return model, tokenizer


# ─── 3. LoRA Config ───────────────────────────────────────────────────────────

def get_lora_config():
    return LoraConfig(
        r=16,              # adapter rank — 8–64 typical; higher = more capacity
        lora_alpha=32,     # scaling: effective_lr_scale = alpha / r
        target_modules=[   # all projection layers for Mistral / Llama architectures
            "q_proj", "k_proj", "v_proj", "o_proj",
            "gate_proj", "up_proj", "down_proj",
        ],
        lora_dropout=0.05,
        bias="none",
        task_type="CAUSAL_LM",
    )


# ─── 4. Training ──────────────────────────────────────────────────────────────

def train():
    dataset           = load_and_format_datasets()
    model, tokenizer  = load_model_and_tokenizer()
    lora_config       = get_lora_config()

    training_args = TrainingArguments(
        output_dir=OUTPUT_DIR,
        num_train_epochs=EPOCHS,
        per_device_train_batch_size=BATCH_SIZE,
        gradient_accumulation_steps=GRAD_ACCUM,
        learning_rate=2e-4,
        fp16=True,
        logging_steps=50,
        save_strategy="epoch",
        save_total_limit=2,
        optim="paged_adamw_8bit",    # paged optimiser offloads states to CPU RAM
        lr_scheduler_type="cosine",
        warmup_ratio=0.03,
        group_by_length=True,        # speeds up training by batching similar lengths
        report_to="none",
    )

    trainer = SFTTrainer(
        model=model,
        train_dataset=dataset,
        dataset_text_field="text",
        max_seq_length=MAX_SEQ_LEN,
        tokenizer=tokenizer,
        args=training_args,
        peft_config=lora_config,
    )

    print("🚀 Starting training...")
    trainer.train()
    trainer.save_model(OUTPUT_DIR)
    tokenizer.save_pretrained(OUTPUT_DIR)
    print(f"✅ Training complete — adapter saved to {OUTPUT_DIR}")

def main():
    train()
if __name__ == "__main__":
    main()
