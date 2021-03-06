from marshmallow import Schema, fields


class CreateUserSchema(Schema):
    phone = fields.String(required=True)
    name = fields.String(required=True)
    username = fields.String(required=True)
    role = fields.String(required=True)


class LoginSchema(Schema):
    phone = fields.String(required=True)
    password = fields.String(required=True)
