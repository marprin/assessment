from app.handler import user_route


def register_router(app):
    app.register_blueprint(user_route, url_prefix="/user")
