import os
import logging
import ollama

from src.model import LlamaModel

from ollama import chat
from ollama import ChatResponse

from .config import (
  OLLAMA_MODEL
)

class LlamaChat:
  def __init__(self):
    # print("Initializing LlamaChat with model:", OLLAMA_MODEL)
    llama_model = LlamaModel(OLLAMA_MODEL)
    self.model = llama_model.model
    self.vector_db = llama_model.load_vector_db()

  def get_response(self, messages):
    response : ChatResponse = chat(
        model=OLLAMA_MODEL,
        messages=messages,
    )
    return response.message.content
