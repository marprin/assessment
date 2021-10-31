from functools import wraps

from app.constant import ErrTokenNotGiven
from flask import g, request


def token_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        header = request.headers.get("Authorization") or None
        if header is None:
            return {"errors": {"message": ErrTokenNotGiven}}, 401

        token = header.split("Bearer ")[-1]
        if token is None:
            return {"errors": {"message": ErrTokenNotGiven}}, 401

        g.token = token

        return f(*args, **kwargs)

    return decorated_function
