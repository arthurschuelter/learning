from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from schemas.habit import HabitCreate, HabitResponse
from infrastructure.database.connection import get_db
from infrastructure.repositories.habit_repository import HabitRepository

habits_list = []
router = APIRouter(
    prefix="/habits",
    tags=["habits"],
    # dependencies=[Depends(get_token_header)],
    responses={404: {"description": "Not found"}},
)

def get_habit_repository(db_session: Session = Depends(get_db)):
      return HabitRepository(db_session)

@router.get("/")
async def habits(habit_repository: HabitRepository = Depends(get_habit_repository)):
    user_id = "a4956161-30d2-4792-924f-79d9e5260b05"
    return await habit_repository.get_by_user_id(user_id)

@router.post("/")
async def create_habit(habit: HabitCreate, habit_repository: HabitRepository = Depends(get_habit_repository)):
    habit_repository.create(habit)
    # habit.id = len(habits_list) + 1
    # habit.print()
    # habits_list.append(habit)
    return habit
