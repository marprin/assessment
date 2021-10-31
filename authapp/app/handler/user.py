from app.constant import ErrInternalServer, SuccessCreateUser
from app.entity.user import CreateUserSchema
from app.service.user_service import register
from flask import Blueprint, jsonify, request
from marshmallow import ValidationError
from utils.error import DuplicateValueError

user_route = Blueprint("user", __name__)


@user_route.route("/", methods=["POST"])
def create_user():
    try:
        payload = CreateUserSchema().load(request.get_json())
    except ValidationError as e:
        return jsonify({"errors": e.messages}), 422

    try:
        register(
            {
                "name": payload["name"],
                "phone": payload["phone"],
                "role": payload["role"],
                "username": payload["username"],
            }
        )
        return jsonify({"message": SuccessCreateUser}), 201
    except DuplicateValueError as e:
        return jsonify({"errors": {"message": "Username is taken"}}), 400
    except Exception as e:
        return (
            jsonify({"errors": {"message": ErrInternalServer}}),
            500,
        )
