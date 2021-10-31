import logging
from random import choice

from app.constant import ErrGeneratePassword
from app.repository.user_repo import create_user, is_password_unique, is_username_unique
from utils.error import DuplicateValueError

logger = logging.getLogger("app.service.user_service")


def register(payload):
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

    return None


def generate_password(length: int) -> str:
    options = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    return "".join(choice(options) for i in range(length))
