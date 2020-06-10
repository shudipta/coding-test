### Sample CRUD API for the mysql database root:root@tcp(localhost:3306)/test

## Example
The project is a RESTful api for accessing the mysql database root:root@tcp(localhost:3306)/test.

## Run

```bash
go run app/server/main.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.16
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:80

```

So the server will run at localhost:80.

## API

Use Postman or `curl` to test the api.

- set rating
  method: PATCH
  url: http://localhost:80/set-rating
  body:
  ```json
  {
  	"product_id": 2,
  	"order_id": 8,
  	"rating": 4.7
  }
  ```
  response:
  
  `"saved successfully"`

- top categories
  method: GET
  url: http://localhost:80/top-categories
  response:
  ```json
  [
      {
          "id": 1,
          "category_name": "",
          "description": {
              "String": "",
              "Valid": false
          }
      }
  ]
  ```