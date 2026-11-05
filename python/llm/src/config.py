from dotenv import load_dotenv
import os

load_dotenv()

# OLLAMA_MODEL = os.getenv("OLLAMA_MODEL", "gemma3")
OLLAMA_MODEL = os.getenv("OLLAMA_MODEL", "llama3.2")
EMBEDDING_MODEL = os.getenv("EMBEDDING_MODEL", "nomic-embed-text")
DOC_PATH = os.getenv("DOC_PATH", "./data/data.pdf")
VECTOR_STORE_NAME = os.getenv("VECTOR_STORE_NAME", "simple-rag")
PERSIST_DIRECTORY = os.getenv("PERSIST_DIRECTORY", "./chroma_db")