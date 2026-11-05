
import os
import logging
import ollama

from langchain_ollama import ChatOllama
from langchain_community.document_loaders import UnstructuredPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_chroma import Chroma
from langchain_ollama import OllamaEmbeddings

from .config import (
  OLLAMA_MODEL,
  EMBEDDING_MODEL,
  DOC_PATH,
  VECTOR_STORE_NAME,
  PERSIST_DIRECTORY,
)

class LlamaModel:
    def __init__(self, model_name=OLLAMA_MODEL):
        # print("Initializing LlamaModel:", model_name)
        self.model = ChatOllama(model=model_name)
        ollama.pull(EMBEDDING_MODEL)
        self.embedding_model = OllamaEmbeddings(model=EMBEDDING_MODEL)

    def load_vector_db(self):
        """Load or create the vector database."""
        if os.path.exists(PERSIST_DIRECTORY):
            vector_db = Chroma(
                embedding_function=self.embedding_model,
                collection_name=VECTOR_STORE_NAME,
                persist_directory=PERSIST_DIRECTORY,
            )
            logging.info("Loaded existing vector database.")
        else:
            # Load and process the PDF document
            data = self.selfingest_pdf(DOC_PATH)
            if data is None:
                return None

            # Split the documents into chunks
            chunks = self.split_documents(data)

            vector_db = Chroma.from_documents(
                documents=chunks,
                embedding=self.embedding_model,
                collection_name=VECTOR_STORE_NAME,
                persist_directory=PERSIST_DIRECTORY,
            )
            vector_db.persist()
            logging.info("Vector database created and persisted.")
        return vector_db
    
    def split_documents(self, documents):
        """Split documents into smaller chunks."""
        text_splitter = RecursiveCharacterTextSplitter(chunk_size=1200, chunk_overlap=300)
        chunks = text_splitter.split_documents(documents)
        logging.info("Documents split into chunks.")
        return chunks 

    def ingest_pdf(self, doc_path):
        """Load PDF documents."""
        if os.path.exists(doc_path):
            loader = UnstructuredPDFLoader(file_path=doc_path)
            data = loader.load()
            logging.info("PDF loaded successfully.")
            return data
        else:
            logging.error(f"PDF file not found at path: {doc_path}")
            print("PDF file not found.")
            return None 