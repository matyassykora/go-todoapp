# go-todoapp

> golang, htmx & postgres todoapp

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Technologies used
- [Go](https://go.dev/)
- [HTMX](https://htmx.org/)
- [PostgreSQL](https://www.postgresql.org/)

## Running

### Create the database
```sh
psql -U postgres -f todo-db.pgsql
```

### Install required packages
```sh
go install .
```

### Run the server
(By default, the requests are slowed down by one second, this can be changed by setting the *demoSleepTime* variable in *handlers/todo.go*)
```sh
go run main.go
```
or
```sh
go build && ./todoapp
```
or (if you have [Air](https://github.com/cosmtrek/air) intalled)
```sh
air
```

### Endpoints
(Endpoints return HTML)

Return a page with some performance data:
```
GET /metrics
```

Get a list of all todos:
```
GET /
GET /todos
```

Get a filtered list of todos:
```
GET /todos?filter=all
GET /todos?filter=done
GET /todos?filter=notdone
```

Add a todo and render it at the end of the list:
```
POST /todos
```

Delete a todo and remove it from the DOM:
```
DELETE /todos/:id
```

Get a form to edit a todo:
```
GET /todos/edit/:id
```

Update a todo and render the edited todo item:
```
PATCH /todos/update/:id
```

Toggle a todo:
```
PATCH /todos/toggle/:id
```
