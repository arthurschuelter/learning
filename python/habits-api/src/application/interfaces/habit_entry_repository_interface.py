from abc import abstractmethod
from uuid import UUID
from typing import Optional, List
from domain.entities.habit_entry import HabitEntry

class HabitEntryRepositoryInterface:
    @abstractmethod
    async def create(self, habit_entry: HabitEntry):
        pass

    @abstractmethod
    async def get_by_id(self, habit_entry_id: UUID) -> Optional[HabitEntry]:
        pass

    @abstractmethod
    async def get_by_habit_id(self, habit_id: UUID) -> List[HabitEntry]:
        pass

    @abstractmethod
    async def update(self, habit_entry: HabitEntry) -> HabitEntry:
        pass

    @abstractmethod
    async def delete(self, habit_entry_id: UUID) -> bool:
        pass