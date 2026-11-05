from uuid import UUID, uuid4
from typing import Optional, List
from datetime import datetime, timedelta

class Habit:
    def __init__(
        self,
        id: UUID, 
        name: str,
        description: str,
        user_id: UUID,
        created_at: Optional[datetime] = None,
        updated_at: Optional[datetime] = None,
        entries: Optional[List] = None,
    ):
        self.id = id
        self.name = name
        self.description = description
        self.user_id = user_id
        self.created_at = created_at or datetime.now()
        self.updated_at = updated_at or datetime.now()
        self.entries = entries or []
        self.streak = self.calculate_streak()
    
    def calculate_streak(self) -> int:
        if not self.entries:
            return 0
        
        # Sort entries by date in descending order
        sorted_entries = sorted(self.entries, key=lambda x: x.entry_date, reverse=True)
        
        today = datetime.now().date()
        most_recent_entry = sorted_entries[0].entry_date.date()
        days_difference = (today - most_recent_entry).days
        if days_difference > 1:
            return 0  # Streak broken

        today = today if days_difference == 0 else today - timedelta(days=1)
        streak = 0
        for entry in sorted_entries:
            entry_date = entry.entry_date.date()
            if entry_date == today - timedelta(days=streak):
                streak += 1
            else:
                break
        
        return streak