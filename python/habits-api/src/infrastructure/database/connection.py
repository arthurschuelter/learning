from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import os
# from dotenv import load_dotenv

# load_dotenv()

engine = create_engine(f"postgresql://admin:adminpassword@localhost:5432/habitsdb")
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()