// README.md
# Library Book Management System
### Introduction
Wanted to learn Go & Gin so made an attempt to build this very very simple API.
### Building & Running 
* Clone this repository [here](https://github.com/DrewK11/Go-Gin-Simple-REST-API.git).
* Cd into the project root folder.
* Then we can use this command to run:
```
go run main.go
```
### Usage
* Once the app is running, you can make requests to the API. Here are some examples:
### API Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | /books/{id} | To retrieve a book with the specified {id} |
| GET | /books | To get all books |
| POST | /books | To create a new book |
| PATCH | /return?id={id} | To return a book with the specified {id} |
| PATCH | /checkout?id={id} | To check out a book with the specified {id} |

* Please include a JSON body when creating a book. Please see below for an example:
``` 
{
    "id": "4",
    "title": "Hamlet",
    "author": "William Shakespeare",
    "quantity": 2
}
```
### Technologies Used
* Go
* Gin