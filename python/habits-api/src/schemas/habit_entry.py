from pydantic import BaseModel, Field
from typing import Optional
from uuid import UUID

class HabitEntryCreate(BaseModel):
    habit_id: UUID = Field(..., example="123e4567-e89b-12d3-a456-426614174000")
    entry_date: str = Field(..., example="2024-06-15T00:00:00Z")
    duration: int = Field(..., example=30)

    def print(self):
        print(f"HabitEntryCreate(habit_id={self.habit_id}, entry_date={self.entry_date}, duration={self.duration})")

class HabitEntryResponse(BaseModel):
    habit_id: UUID = Field(..., example="123e4567-e89b-12d3-a456-426614174000")
    entry_date: str = Field(..., example="2024-06-15T00:00:00Z")
    duration: int = Field(..., example=30)