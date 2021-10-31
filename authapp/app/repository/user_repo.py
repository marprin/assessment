from datetime import datetime

from app.model import User
from pony.orm import db_session


@db_session
def is_username_unique(username: str) -> User:
    return User.get(username=username)


@db_session
def is_password_unique(pwd: str) -> User:
    return User.get(password=pwd)


@db_session
def create_user(username: str, name: str, phone: str, pwd: str, role: str):
    return User(
        name=name,
        username=username,
        password=pwd,
        phone=phone,
        role=role,
        created_at=datetime.utcnow(),
    )


@db_session
def find_user_by_phone_and_pwd(phone: str, pwd: str) -> User:
    return User.get(phone=phone, password=pwd)
