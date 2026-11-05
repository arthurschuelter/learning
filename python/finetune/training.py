import torch
from datasets import load_dataset
from transformers import (
    AutoModelForCausalLM,
    AutoTokenizer,
    BitsAndBytesConfig,
    TrainingArguments,
)
from peft import LoraConfig, get_peft_model, prepare_model_for_kbit_training
from trl import SFTTrainer

# ── Config ──────────────────────────────────────────────────────────────────
BASE_MODEL   = "mistralai/Mistral-7B-v0.1"   # swap for any HF model
OUTPUT_DIR   = "./qlora-output"
DATASET_NAME = "timdettmers/openassistant-guanaco"  # or your own dataset
MAX_SEQ_LEN  = 512
EPOCHS       = 1

# ── 4-bit quantisation (QLoRA) ───────────────────────────────────────────────
bnb_config = BitsAndBytesConfig(
    load_in_4bit=True,
    bnb_4bit_quant_type="nf4",          # NormalFloat4 — best for QLoRA
    bnb_4bit_compute_dtype=torch.float16,
    bnb_4bit_use_double_quant=True,     # double quantisation saves ~0.4 bpp
)

# ── Load model & tokeniser ───────────────────────────────────────────────────
model = AutoModelForCausalLM.from_pretrained(
    BASE_MODEL,
    quantization_config=bnb_config,
    device_map="auto",
    trust_remote_code=True,
)
model = prepare_model_for_kbit_training(model)   # cast layernorms to fp32

tokenizer = AutoTokenizer.from_pretrained(BASE_MODEL, trust_remote_code=True)
tokenizer.pad_token = tokenizer.eos_token
tokenizer.padding_side = "right"

# ── LoRA adapter ─────────────────────────────────────────────────────────────
lora_config = LoraConfig(
    r=16,                    # rank — higher = more capacity, more VRAM
    lora_alpha=32,           # scaling factor (alpha/r = effective lr scale)
    target_modules=[         # which linear layers to adapt
        "q_proj", "k_proj", "v_proj", "o_proj",
        "gate_proj", "up_proj", "down_proj",
    ],
    lora_dropout=0.05,
    bias="none",
    task_type="CAUSAL_LM",
)
model = get_peft_model(model, lora_config)
model.print_trainable_parameters()   # should be ~1-2% of total params

# ── Dataset ──────────────────────────────────────────────────────────────────
dataset = load_dataset(DATASET_NAME, split="train")

# ── Training ─────────────────────────────────────────────────────────────────
training_args = TrainingArguments(
    output_dir=OUTPUT_DIR,
    num_train_epochs=EPOCHS,
    per_device_train_batch_size=4,
    gradient_accumulation_steps=4,   # effective batch = 16
    learning_rate=2e-4,
    fp16=True,
    logging_steps=50,
    save_strategy="epoch",
    optim="paged_adamw_8bit",        # paged optimiser saves VRAM
    lr_scheduler_type="cosine",
    warmup_ratio=0.03,
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

trainer.train()
trainer.save_model(OUTPUT_DIR)
print("✅ Training complete — adapter saved to", OUTPUT_DIR)