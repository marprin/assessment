import os

BASE_DIR = os.path.abspath(os.path.dirname(__file__))

JWT_SECRET = os.getenv("JWT_SECRET")
JWT_EXPIRE_IN = os.getenv("JWT_EXPIRE_IN", 3600)

LOGGING_CONFIG = {
    "version": 1,
    "disable_existing_loggers": False,
    "formatters": {
        "default": {
            "format": "%(asctime)s - %(process)s - %(name)s - %(levelname)s - %(message)s"
        },
        "json_formatter": {
            "class": "pythonjsonlogger.jsonlogger.JsonFormatter",
            "format": "%(asctime)s %(process)d %(threadName)s "
            "%(name)s %(levelname)s %(pathname)s %(lineno)s %(message)s",
            "datefmt": "%Y-%m-%dT%H:%M:%S%z",
        },
    },
    "handlers": {
        "wsgi": {
            "class": "logging.StreamHandler",
            "stream": "ext://flask.logging.wsgi_errors_stream",
            "formatter": "default",
        },
        "console": {
            "class": "logging.StreamHandler",
            "level": "ERROR",
            "formatter": "json_formatter",
            "stream": "ext://sys.stdout",
        },
    },
    "loggers": {
        "flask.app": {
            "level": "DEBUG",
            "propagate": False,
            "handlers": ["wsgi", "console"],
        }
    },
    "root": {"level": "ERROR", "handlers": ["wsgi", "console"], "propagate": False},
}
