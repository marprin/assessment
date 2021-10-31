import logging
from datetime import datetime, timedelta

import jwt
from jwt.exceptions import DecodeError, ExpiredSignatureError

logger = logging.getLogger("app.repository.jwt_repo")


def generate_token(payload: dict, expire_in: int, secret: str) -> str:
    prep_jwt = {
        "iss": "authapp-http",
        "sub": "authapp",
        "iat": int(datetime.now().timestamp()),
        "exp": int((datetime.now() + timedelta(minutes=expire_in)).timestamp()),
        "data": payload,
    }

    return jwt.encode(prep_jwt, secret, algorithm="HS256")


def parse_token(token: str, secret: str):
    try:
        token_data = jwt.decode(token, secret, algorithms=["HS256"])
    except ExpiredSignatureError as e:
        logger.error(dict(message="Token expired", err=repr(e)))
        raise
    except DecodeError as e:
        logger.error(dict(message="Decode failed", err=repr(e)))
        raise

    return token_data
