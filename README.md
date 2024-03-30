# book_management_crud_api_sql (WIP)

This is a simple CRUD API written in Go using the [gorilla/mux](https://github.com/gorilla/mux) package and [GORM](https://gorm.io/)

It is based on [this](https://www.youtube.com/watch?v=jFfo23yIWac&t=4034s) tutorial

| Endpoint    | Method |
| ----------- | ------ |
| /book       | GET    |
| /book/model | GET    |
| /book       | POST   |
| /book/{id}  | GET    |
| /book/{id}  | PUT    |
| /book/{id}  | DELETE |

Here is the JSON representation of a "Book" object:

```json
"Book": {
    "name": {
        "type": "string"
    },
    "author": {
        "type": "string"
    },
    "publication": {
        "type": "string"
    }
}
```

## Running the program

```console
go run cmd/main/main.go
```
