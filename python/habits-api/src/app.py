import os
from fastapi import FastAPI
from routers import habits, habit_entries

app = FastAPI()

app.include_router(habits.router, prefix="/api")
app.include_router(habit_entries.router, prefix="/api")

@app.get("/")
async def root():
    return {"message": "Hello Habits API!"}

def create_app():
    print("Creating FastAPI app...")

if __name__ == '__main__':
    create_app()