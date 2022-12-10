# cake-store-api

A RestFul API build using Go to demonstrate basic CRUD. Code organized using Clean Architecture.

## How to run

Clone the repository

```
git clone git@github.com:guntoroyk/cake-store-api.git
```

Run on docker

```
docker-compose up --build -d
```

The API will available at http://localhost:8000/cakes

## API Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/6929584-5922eef1-a2f0-44af-bab8-ca43e188999d?action=collection%2Ffork&collection-url=entityId%3D6929584-5922eef1-a2f0-44af-bab8-ca43e188999d%26entityType%3Dcollection%26workspaceId%3D376ac51e-7371-4129-b917-abb587ed642f)

Or click this URL https://elements.getpostman.com/redirect?entityId=6929584-5922eef1-a2f0-44af-bab8-ca43e188999d&entityType=collection

The API has 5 endpoints:

1. [Create a Cake](#create-a-cake)
2. [Get List of Cakes](#get-list-cakes)
3. [Get a Cake](#get-a-cake)
4. [Update a Cake](#update-a-cake)
5. [Delete a Cake](#delete-a-cake)

The API will return response with the following format:

```
{
  "code": number,
  "data": any,
  "error": string
}
```

### Create a Cake

#### Request

`POST /cakes`

```
curl -X POST \
  'localhost:8000/cakes' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "title": "Lemon cheesecake 3",
  "description": "A cheesecake made of lemon",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}'
```

#### Response

```
{
  "code": 201,
  "data": {
    "id": 5,
    "title": "Lemon cheesecake 3",
    "description": "A cheesecake made of lemon",
    "rating": 7,
    "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    "created_at": "2022-12-06 09:26:50",
    "updated_at": "2022-12-06 09:26:50"
  }
}
```

### Get List Cakes

#### Request

`GET /cakes`

```
curl -X GET \
  'localhost:8000/cakes' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

#### Response

```
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "title": "Lemon cheesecake 1",
      "description": "A cheesecake made of lemon",
      "rating": 7,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2022-12-05 16:23:43",
      "updated_at": "2022-12-05 16:59:13"
    },
    {
      "id": 2,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2022-12-05 16:44:52",
      "updated_at": "2022-12-05 16:44:52"
    },
    {
      "id": 5,
      "title": "Lemon cheesecake 3",
      "description": "A cheesecake made of lemon",
      "rating": 7,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2022-12-06 09:26:50",
      "updated_at": "2022-12-06 09:26:50"
    }
  ]
}
```

### Get a Cake

#### Request

`GET /cakes/:id`

```
curl -X GET \
  'localhost:8000/cakes/1' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

#### Success Response

```
{
  "code": 200,
  "data": {
    "id": 1,
    "title": "Lemon cheesecake 1",
    "description": "A cheesecake made of lemon",
    "rating": 7,
    "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    "created_at": "2022-12-05 16:23:43",
    "updated_at": "2022-12-05 16:59:13"
  }
}
```

#### Error Response

```
{
  "code": 404,
  "error": "cake not found"
}
```

### Update a Cake

#### Request

`PATC /cakes/:id`

```
curl -X PATCH \
  'localhost:8000/cakes/1' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "title": "Lemon cheesecake updated",
  "description": "A cheesecake made of lemon updated",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}'
```

#### Success Response

```
{
  "code": 200,
  "data": {
    "id": 1,
    "title": "Lemon cheesecake updated",
    "description": "A cheesecake made of lemon updated",
    "rating": 7,
    "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    "created_at": "",
    "updated_at": ""
  }
}
```

#### Error Response

```
{
  "code": 404,
  "error": "cake not found"
}
```

### Delete a Cake

#### Request

`DELETE /cakes/:id`

```

curl -X DELETE \
 'localhost:8000/cakes/5' \
 --header 'Accept: _/_' \
 --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'

```

#### Success Response

Status Code: 204 No Content

```

```

#### Error Response

```
{
  "code": 404,
  "error": "cake not found"
}

```
