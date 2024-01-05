# go-todoapp

> golang, htmx & postgres todoapp

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
By default, the requests are slowed down by one second, this can be changed by setting the *demoSleepTime* variable in *handlers/todo.go*
```sh
go run main.go
```

### Endpoints
Endpoints return HTML

Return a page with some performance data:
```
GET /metrics
```

Get a list of todos:
```
GET /
GET /todos
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
