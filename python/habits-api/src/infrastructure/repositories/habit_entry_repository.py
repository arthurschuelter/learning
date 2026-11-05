from typing import List, Optional
from uuid import UUID
from application.interfaces.habit_entry_repository_interface import HabitEntryRepositoryInterface
from domain.entities.habit_entry import HabitEntry
from infrastructure.database.connection import get_db
import datetime
from libs.utils.timing import measure_time
from sqlalchemy import text

class HabitEntryRepository(HabitEntryRepositoryInterface):
    def __init__(self, db_session):
        self.db_session = db_session

    def create(self, habit_entry: HabitEntry):
        print("Creating habit entry in the database")
        try:
            query = text("INSERT INTO habit_entries (habit_id, entry_date, duration, created_at) VALUES (:habit_id, :entry_date, :duration, :created_at) RETURNING id")
            result = self.db_session.execute(
                query,
                {"habit_id": habit_entry.habit_id, "entry_date": habit_entry.entry_date, "duration": habit_entry.duration, "created_at": datetime.datetime.now()}
            )
            self.db_session.commit()
            habit_entry_id = result.fetchone()[0]
            print(f"Habit entry created with id: {habit_entry_id}")
            return habit_entry
        except Exception as e:
            print(f"Error creating habit entry: {e}")
            self.db_session.rollback()
            # raise

    # async def get_by_id(self, habit_id: int) -> Optional[Habit]:
    #     return self.db_session.query(Habit).filter(Habit.id == habit_id).first()
    @measure_time
    async def get_by_habit_id(self, habit_id: UUID) -> List[HabitEntry]:
        query = text("SELECT id, habit_id, entry_date, created_at, updated_at FROM habit_entries WHERE habit_id = :habit_id")
        result = self.db_session.execute(query, {"habit_id": str(habit_id)})
        rows = result.fetchall()

        # Convert rows to Habit domain entities
        habits = []
        for row in rows:
            habit = HabitEntry(
                habit_id=row[1],
                entry_date=row[2],
                created_at=row[3],
                updated_at=row[4]
            )
            habits.append(habit)
        return habits

    # async def update(self, habit: Habit) -> Habit:
    #     self.db_session.merge(habit)
    #     self.db_session.commit()
    #     return habit

    # async def delete(self, habit_id: int) -> bool:
    #     habit = self.db_session.query(Habit).filter(Habit.id == habit_id).first()
    #     if habit:
    #         self.db_session.delete(habit)
    #         self.db_session.commit()
    #         return True
    #     return False