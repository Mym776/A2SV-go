# Task Manager
Before starting make sure to install all necessary packages and dependencies using the go mod tidy command
This collection is made for interacting with the task manager api's various functionalities. The collection includes requests for:


## Registration and logging in 

### Users

Users can register by sending their username and password. If the username is unique then the registration will be successful and the user name and password hash will be stored in the database. If there is no user registered in the database the first user will be assigned to be an admin.

After registration the user can use their username and password to login, if the credentials match then the response will include a success message and a token

The token is valid for 5 minutes and will allow users to access features based on their role.

#### Users 
1. getting the entire list of tasks
2. getting a task by its ID

#### Admins
1. Adding new tasks
2. Updating existing tasks
3. Deleting tasks
4. Promoting users to admins

if any user tries to access the endpoints with no authorization or with an expired token, they will get an error message stating to try with a proper token

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
