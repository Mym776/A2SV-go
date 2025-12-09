package service

import (
	// "fmt"
	// "context"
	// "fmt"
	"log"
	"taskManager/models"
	

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Task data
var MyURL = "mongodb://localhost:27017"

func Tasks(c *gin.Context) []models.Task {
	client := ConnectDB(c)
	// defer statement closes the connection with the db after the function is done
	defer DisconnectDB(c, client)
	collection := client.Database("Task_manager").Collection("Tasks")

	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Fatal(err)
	}


	defer cursor.Close(c)

	var taskList = []models.Task{}

	// loops through the results in the cursor and appends it to an array 
	for cursor.Next(c) {
		var result models.Task
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		taskList = append(taskList, result)
	}

	return taskList
}

func TaskId(c *gin.Context, id string) models.Task {
	client := ConnectDB(c)
	defer DisconnectDB(c, client)

	collection := client.Database("Task_manager").Collection("Tasks")

	var result models.Task
	err := collection.FindOne(c, bson.D{{"id", id}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			//if not found return an empty task instance
			return models.Task{}
		}
		log.Fatal(err) // Return other errors
	}

	return result
}

func UpdateTask(c *gin.Context, id string, update models.Task) bool {

	client := ConnectDB(c)
	defer DisconnectDB(c, client)

	collection := client.Database("Task_manager").Collection("Tasks")

	filter := bson.D{
		{"id", id},
	}
	updated := bson.D{
		{"$set", bson.D{
			{"id", update.ID},
			{"title", update.Title},
			{"description", update.Description},
			{"due_date", update.DueDate},
			{"status", update.Status},
		}},
	}

	up, err := collection.UpdateOne(c, filter, updated)
	if err != nil {
		return false
	}
	

	if up.ModifiedCount > 0 {
		return true

	}
	return false

}

func DeleteTask(c *gin.Context,id string) bool {
	client := ConnectDB(c)
	defer DisconnectDB(c,client)
	collection := client.Database("Task_manager").Collection("Tasks")

	filter := bson.D{
		{"id",id},
	}

	
	_, err := collection.DeleteOne(c, filter)

	if err != nil {
		log.Fatal(err)
	}

	return true
	
}

func AddTask(c *gin.Context, task models.Task) bool {
	client := ConnectDB(c)
	defer DisconnectDB(c,client)
	collection := client.Database("Task_manager").Collection("Tasks")

	filter := bson.D{
		{"id",task.ID},
	}

	var result models.Task
	err := collection.FindOne(c, filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		newTask := bson.D{
			{"id",task.ID},
			{"title", task.Title},
			{"description", task.Description},
			{"due_date", task.DueDate},
			{"status", task.Status},
		}

		_ ,err := collection.InsertOne(c,newTask)
		if err != nil{
			log.Fatal(err)
		}
		return true
	}else{
		return false
	}
}


//connects to the database and returns the client
func ConnectDB(c *gin.Context) *mongo.Client {

	clientOptions := options.Client().ApplyURI(MyURL)

	client, err := mongo.Connect(c, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client

}

// disconnects from the database
func DisconnectDB(c *gin.Context, db *mongo.Client) {
	err := db.Disconnect(c)

	if err != nil {
		log.Fatal(err)
	}
}
