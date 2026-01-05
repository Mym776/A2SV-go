
package repositories

import (
	
	"log"
	entities "taskManager/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	Context *gin.Context
	
}

var MyURL = "mongodb://localhost:27017"


func ConnectDB(c *gin.Context, collection string) *mongo.Collection {
	
	clientOptions := options.Client().ApplyURI(MyURL)

	client, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	table :=client.Database("Task_manager").Collection(collection)

	
	return table 

}
func DisconnectDB(c *gin.Context, db *mongo.Client) {
	err := db.Disconnect(c)

	if err != nil {
		log.Fatal(err)
	}
}


func (t *TaskRepository)GetAllTask() []entities.Task{
	c:= t.Context
	tasks := ConnectDB(c,"Tasks")
	defer DisconnectDB(c,tasks.Database().Client())


	var taskList []entities.Task


	cursor, err := tasks.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}


	defer cursor.Close(c)


	// loops through the results in the cursor and appends it to an array 
	for cursor.Next(c) {
		var result entities.Task
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		taskList = append(taskList, result)
	}

	return taskList
}

func (t *TaskRepository)GetTaskByID(id string) entities.Task{
	c := t.Context
	tasks := ConnectDB(c,"Tasks")
	defer DisconnectDB(c,tasks.Database().Client())

	
	var result entities.Task
	err := tasks.FindOne(c, bson.D{{Key: "id",Value: id}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			//if not found return an empty task instance
			return entities.Task{}
		}
		log.Fatal(err) // Return other errors
	}

	return result

}

func(t *TaskRepository) AddTask(task entities.Task) bool{
	c := t.Context
	tasks := ConnectDB(c,"Tasks")
	defer DisconnectDB(c,tasks.Database().Client())

	
	filter := bson.D{
		{Key: "id",Value: task.ID},
	}

	var result entities.Task
	err := tasks.FindOne(c, filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		newTask := bson.D{
			{Key: "id",Value: task.ID},
			{Key: "title", Value: task.Title},
			{Key: "description", Value: task.Description},
			{Key: "due_date", Value: task.DueDate},
			{Key: "status", Value: task.Status},
		}

		_ ,err := tasks.InsertOne(c,newTask)
		if err != nil{
			log.Fatal(err)
		}
		return true
	}else{
		return false
	}
}


func(t *TaskRepository) UpdateTask(update entities.Task, id string)bool{
	c := t.Context
	tasks := ConnectDB(c,"Tasks")
	defer DisconnectDB(c,tasks.Database().Client())



	filter := bson.D{
		{Key: "id",Value: id},
	}
	updated := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "id",Value: update.ID},
			{Key: "title", Value: update.Title},
			{Key: "description", Value: update.Description},
			{Key: "due_date", Value: update.DueDate},
			{Key: "status", Value: update.Status},
		}},
	}

	up, err := tasks.UpdateOne(c, filter, updated)
	if err != nil {
		return false
	}
	

	if up.ModifiedCount > 0 {
		return true

	}
	return false
}

func(t *TaskRepository) DeleteTask(id string) bool{
	c := t.Context
	tasks := ConnectDB(c,"Tasks")
	defer DisconnectDB(c,tasks.Database().Client())
	
	filter := bson.D{
		{Key: "id",Value: id},
	}

	del, err := tasks.DeleteOne(c, filter)

	if err != nil {
		log.Fatal(err)
	}
	if del.DeletedCount >0{
		return true
	}

	return false
}