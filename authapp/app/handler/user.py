from app.constant import ErrInternalServer, ErrTokenNotValid, SuccessCreateUser
from app.decorator import token_required
from app.entity import CreateUserSchema, LoginSchema
from app.service.user_service import extract_token_info, login, register
from flask import Blueprint, g, jsonify, request
from jwt.exceptions import DecodeError
from marshmallow import ValidationError
from utils.error import DuplicateValueError, NotFound

user_route = Blueprint("user", __name__)


@user_route.route("/", methods=["POST"])
def create_user_handler():
    try:
        payload = CreateUserSchema().load(request.get_json())
    except ValidationError as e:
        return jsonify({"errors": e.messages}), 422

    try:
        generated_pwd = register(
            {
                "name": payload["name"],
                "phone": payload["phone"],
                "role": payload["role"],
                "username": payload["username"],
            }
        )
        return jsonify({"password": generated_pwd}), 201
    except DuplicateValueError as e:
        return jsonify({"errors": {"message": "Username is taken"}}), 400
    except Exception as e:
        return (
            jsonify({"errors": {"message": ErrInternalServer}}),
            500,
        )


@user_route.route("/login", methods=["POST"])
def login_handler():
    try:
        payload = LoginSchema().load(request.get_json())
    except ValidationError as e:
        return jsonify({"errros": e.messages}), 422

    try:
        token = login(payload["phone"], payload["password"])
        return jsonify({"token": token}), 200
    except NotFound as e:
        return (
            jsonify({"errors": {"message": "Please check your phone and password"}}),
            400,
        )
    except Exception as e:
        return (
            jsonify({"errors": {"message": ErrInternalServer}}),
            500,
        )


@user_route.route("/", methods=["GET"])
@token_required
def user():
    token = g.get("token")

    try:
        return jsonify(extract_token_info(token)["data"]), 200
    except DecodeError as e:
        return (
            jsonify({"errors": {"message": ErrTokenNotValid}}),
            400,
        )
    except Exception as e:
        return (
            jsonify({"errors": {"message": ErrInternalServer}}),
            500,
        )
