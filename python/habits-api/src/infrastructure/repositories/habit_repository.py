from typing import List, Optional
from uuid import UUID
from application.interfaces.habit_repository_interface import HabitRepositoryInterface
from domain.entities.habit import Habit
from domain.entities.habit_entry import HabitEntry
from infrastructure.database.connection import get_db
import datetime

from sqlalchemy import text

class HabitRepository(HabitRepositoryInterface):
    def __init__(self, db_session):
        self.db_session = db_session

    def create(self, habit: Habit):
        print("Creating habit in the database")
        try:
            query = text("INSERT INTO habits (name, description, user_id, created_at) VALUES (:name, :description, :user_id, :created_at) RETURNING id")
            result = self.db_session.execute(
                query,
                {"name": habit.name, "description": habit.description, "user_id": habit.user_id, "created_at": datetime.datetime.now()}
            )
            self.db_session.commit()
            habit_id = result.fetchone()[0]
            print(f"Habit created with id: {habit_id}")
            return habit
        except Exception as e:
            print(f"Error creating habit: {e}")
            self.db_session.rollback()
            raise

    # async def get_by_id(self, habit_id: int) -> Optional[Habit]:
    #     return self.db_session.query(Habit).filter(Habit.id == habit_id).first()

    async def get_by_user_id(self, user_id: UUID) -> List[Habit]:

        # Fetch habits
        query = text("SELECT id, name, description, user_id, created_at, updated_at FROM habits WHERE user_id = :user_id")
        result = self.db_session.execute(query, {"user_id": str(user_id)})
        rows = result.fetchall()

        # Convert rows to Habit domain entities with their entries
        habits = []
        for row in rows:
            habit_id = row[0]

            # Fetch habit entries for this habit
            entries_query = text("SELECT id, habit_id, entry_date, duration, created_at, updated_at FROM habit_entries WHERE habit_id = :habit_id ORDER BY entry_date DESC")
            entries_result = self.db_session.execute(entries_query, {"habit_id": str(habit_id)})
            entries_rows = entries_result.fetchall()

            # Convert entry rows to HabitEntry domain entities
            entries = []
            for entry_row in entries_rows:
                entry = HabitEntry(
                    habit_id=entry_row[1],
                    entry_date=entry_row[2],
                    duration=entry_row[3],
                    created_at=entry_row[4],
                    updated_at=entry_row[5]
                )
                entries.append(entry)

            habit = Habit(
                id=row[0],
                name=row[1],
                description=row[2],
                user_id=row[3],
                created_at=row[4],
                updated_at=row[5],
                entries=entries
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