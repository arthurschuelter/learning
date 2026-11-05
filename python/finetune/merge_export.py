# merge_and_export.py — run AFTER training

from transformers import AutoModelForCausalLM, AutoTokenizer
from peft import PeftModel
import torch

BASE_MODEL  = "mistralai/Mistral-7B-v0.1"
ADAPTER_DIR = "./qlora-output"
MERGED_DIR  = "./merged-model"

tokenizer = AutoTokenizer.from_pretrained(BASE_MODEL)
base = AutoModelForCausalLM.from_pretrained(
    BASE_MODEL, torch_dtype=torch.float16, device_map="cpu"
)
model = PeftModel.from_pretrained(base, ADAPTER_DIR)
model = model.merge_and_unload()          # fuse LoRA weights into base
model.save_pretrained(MERGED_DIR)
tokenizer.save_pretrained(MERGED_DIR)
print("✅ Merged model saved to", MERGED_DIR)