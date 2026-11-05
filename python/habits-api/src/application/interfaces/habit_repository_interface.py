from abc import abstractmethod
from uuid import UUID
from typing import Optional, List
from domain.entities.habit import Habit
# from schemas.habit import HabitCreate

class HabitRepositoryInterface:
    @abstractmethod
    async def create(self, habit: Habit):
        pass

    @abstractmethod
    async def get_by_id(self, habit_id: int) -> Optional[Habit]:
        pass

    @abstractmethod
    async def get_by_user_id(self, user_id: UUID) -> List[Habit]:
        pass

    @abstractmethod
    async def update(self, habit: Habit) -> Habit:
        pass

    @abstractmethod
    async def delete(self, habit_id: int) -> bool:
        pass