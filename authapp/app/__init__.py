import logging.config

import settings
from flask import Flask
from route import register_router
from werkzeug.utils import import_string


def create_app():
    logging.config.dictConfig(import_string("settings.LOGGING_CONFIG"))
    app = Flask(__name__, instance_relative_config=True)
    app.url_map.strict_slashes = False
    app.config.from_object(settings)

    register_router(app)

    return app


app = create_app()
