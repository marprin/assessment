# API Docs

## User Profile
Method: GET
Endpoint: /user
Header:
```
{
    "Content-Type": "application/json",
    "Authorization": "Bearer {changewithyourtoken}"
}
```

Response
```json
{
    "name": "Hello",
    "phone": "0819381918191",
    "role": "admin",
    "timestamp": 1635634763
}
```

## Login
Method: POST
Endpoint: /user/login
Header:
```
{
    "Content-Type": "application/json",
}
```

Body
```json
{
    "phone": "0819381918",
    "password": "pdk8"
}
```

Response
```json
{
    "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJhdXRoYXBwLWh0dHAiLCJzdWIiOiJhdXRoYXBwIiwiaWF0IjoxNjM1NjkyODI2LCJleHAiOjE2MzU5MDg4MjYsImRhdGEiOnsibmFtZSI6IkhlbGxvIiwicGhvbmUiOiIwODE4MzkxODEiLCJyb2xlIjoiYWRtaW4iLCJ0aW1lc3RhbXAiOjE2MzU2OTI4MjZ9fQ.TSSs1vvWAMyX77JdQAjyyHQADey9kEg-VAgJA2XoQh4"
}
```

## Register
Method: POST
Endpoint: /user
Header:
```
{
    "Content-Type": "application/json",
}
```

Body
```json
{
    "name": "Hello",
    "phone": "081839181",
    "role": "admin",
    "username": "admin2"
}
```

Response
```json
{
    "password": "OclK"
}
```