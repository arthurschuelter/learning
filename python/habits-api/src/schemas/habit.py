from pydantic import BaseModel, Field
from typing import Optional
from uuid import UUID

class HabitCreate(BaseModel):
    name: str = Field(..., example="Drink Water")
    description: Optional[str] = Field(None, example="Drink at least 8 glasses of water daily")
    user_id: UUID = Field(..., example="123e4567-e89b-12d3-a456-426614174000")

    def print(self):
        print(f"HabitCreate(name={self.name}, description={self.description}, user_id={self.user_id})")

class HabitResponse(BaseModel):
    id: UUID = Field(..., example=1)
    name: str = Field(..., example="Drink Water")
    description: Optional[str] = Field(None, example="Drink at least 8 glasses of water daily")
    user_id: UUID = Field(..., example="123e4567-e89b-12d3-a456-426614174000")