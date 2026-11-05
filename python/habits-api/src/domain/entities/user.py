from uuid import UUID, uuid4
from typing import Optional
from datetime import datetime

class User:
    def __init__(
        self, 
        user_id: UUID, 
        username: str, 
        email: str,
        created_at: Optional[datetime] = None,
    ):
        self.user_id = user_id
        self.username = username
        self.email = email
        self.created_at = created_at or datetime.now()
