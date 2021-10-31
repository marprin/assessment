from datetime import datetime
from enum import unique

from pony.orm import Optional, Required
from utils.db import db


class User(db.Entity):
    name = Required(str)
    username = Required(str, unique=True)
    password = Required(str, unique=True)
    phone = Required(str)
    role = Required(str)
    created_at = Optional(datetime)
    updated_at = Optional(datetime, nullable=True)


db.generate_mapping(create_tables=True)
