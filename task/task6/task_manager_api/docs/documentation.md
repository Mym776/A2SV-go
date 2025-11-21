# Task Manager
Before starting make sure to install all necessary packages and dependencies using the go mod tidy command
This collection is made for interacting with the task manager api's various functionalities. The collection includes requests for:



###  Getting all tasks

### Getting a task by ID

### Updating a task field by ID

### Deleting a task by ID

### ADD a task 


## MongoDB


URL: mongodb://localhost:27017

to connect to your database change the MyURL variable inside the data/task_service.go to your db url


## GET Get all task

localhost:3000/tasks
Fetches all stored tasks

Endpoint: /tasks

Method: GET

Parameters: None



## GET Get task by ID
localhost:3000/tasks/:id
Fetches task of specified id

Endpoint: /tasks/"id"

Method: GET

Parameters: id

Path Variables
id


## PUT Update task detail
localhost:3000/tasks/:id
Updates task of specified id

Endpoint: /tasks/"id"

Method: PUT

Parameters: id



Path Variables
id



## DELETE Delete task
localhost:3000/tasks/:id
Delete task of specified id

Endpoint: /tasks/"id"

Method: delete

Parameters: id



Path Variables
id

## Add a task
localhost:3000/tasks/new
Adds a new task

Endpoint: /tasks/"id"

Method: POST

Parameters: id



Path Variables
id
