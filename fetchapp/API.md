# API Docs

## User Profile
Method: GET
Endpoint: /user/profile
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
    "data": {
        "name": "Hello",
        "phone": "0819381918191",
        "role": "admin"
    }
}
```

## Storage List
Method: GET
Endpoint: /storage/list
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
    "data": [
        {
            "uuid": "8a91e0cc-51f2-4c7b-b7c6-947571a0fa3a",
            "komoditas": "Udang Vannake",
            "area_provinsi": "BANTEN",
            "area_kota": "ACEH KOTA",
            "size": "110",
            "price": "90978",
            "tgl_parsed": "2021-10-29T16:57:14+07:00",
            "timestamp": "1635501434",
            "usd_price": 6.394045774165323
        },
        ...

    ]
}
```

## Storage Filter
Method: GET
Endpoint: /storage/filter
Header:
```
{
    "Content-Type": "application/json",
    "Authorization": "Bearer {changewithyourtoken}"
}
```

Query Params
```json
{
    "area_provinsi": "",
    "start_date": "2021-08-08T11:20:30",
    "end_date": "2021-08-010T11:20:30",
}
```

Response
```json
{
    "data": {
        "min": 500,
        "max": 500000,
        "avg": 15000,
        "median": 15000
    }
}
```