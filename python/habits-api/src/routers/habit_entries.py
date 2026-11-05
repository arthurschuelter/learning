from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from schemas.habit_entry import HabitEntryCreate, HabitEntryResponse
from infrastructure.database.connection import get_db
from infrastructure.repositories.habit_entry_repository import HabitEntryRepository

habits_list = []
router = APIRouter(
    prefix="/habit_entries",
    tags=["habit_entries"],
    # dependencies=[Depends(get_token_header)],
    responses={404: {"description": "Not found"}},
)

def get_habit_entry_repository(db_session: Session = Depends(get_db)):
      return HabitEntryRepository(db_session)

@router.get("/")
async def habits(habit_entry_repository: HabitEntryRepository = Depends(get_habit_entry_repository)):
    habit_id = "08b8615d-2aca-4a45-8109-96066f912930"
    return await habit_entry_repository.get_by_habit_id(habit_id)

@router.post("/")
async def create_habit(habit_entry: HabitEntryCreate, habit_entry_repository: HabitEntryRepository = Depends(get_habit_entry_repository)):
    habit_entry_repository.create(habit_entry)
    # habit.id = len(habits_list) + 1
    # habit.print()
    # habits_list.append(habit)
    return habit_entry
