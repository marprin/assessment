import logging
from datetime import datetime
from random import choice

import settings
from app.constant import ErrGeneratePassword
from app.repository.jwt_repo import generate_token, parse_token
from app.repository.user_repo import (
    create_user,
    find_user_by_phone_and_pwd,
    is_password_unique,
    is_username_unique,
)
from marshmallow.fields import DateTime
from utils.error import DuplicateValueError, NotFound

logger = logging.getLogger("app.service.user_service")


def register(payload: dict):
    # Check the username
    is_username_taken = is_username_unique(payload["username"])
    if is_username_taken:
        logger.error(dict(message=f"Duplicate username for {payload['username']}"))
        raise DuplicateValueError

    # Generate new password
    max_retry: int = 5
    generated_pwd = None
    while max_retry > 0:
        gen_pwd = generate_password(4)
        is_pwd_taken = is_password_unique(gen_pwd)
        if is_pwd_taken is None:
            generated_pwd = gen_pwd
            break

        max_retry = max_retry - 1

    # When the system failed to generate a password
    if generated_pwd is None:
        logger.error(dict(message="Failed to generate password"))
        raise Exception(ErrGeneratePassword)

    # Save to db
    create_user(
        username=payload["username"],
        name=payload["name"],
        phone=payload["phone"],
        role=payload["role"],
        pwd=generated_pwd,
    )

    return generated_pwd


def generate_password(length: int) -> str:
    options = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    return "".join(choice(options) for i in range(length))


def login(phone: str, pwd: str) -> str:
    user = find_user_by_phone_and_pwd(phone=phone, pwd=pwd)
    if not user:
        raise NotFound

    # When found, generate the token
    payload = {
        "name": user.name,
        "phone": user.phone,
        "role": user.role,
        "timestamp": int(datetime.utcnow().timestamp()),
    }

    return generate_token(
        payload=payload, expire_in=settings.JWT_EXPIRE_IN, secret=settings.JWT_SECRET
    )


def extract_token_info(token: str) -> dict:
    return parse_token(token=token, secret=settings.JWT_SECRET)
