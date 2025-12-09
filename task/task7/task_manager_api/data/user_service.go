package service

import (
	
	"taskManager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//jwt secret used to generate the jwt tokens. STORED HERE FOR EVALUATION PURPOSES
var Jwtsecret = []byte("secret_stuff")

func Register(c *gin.Context) {
	
	
	var user models.User
	err := c.ShouldBind(&user)

	//accept and validate user data
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid request payload"})
		return
	}

	//hash user password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	user.PasswordHash = string(hashedPwd)

	client := ConnectDB(c)

	userTable := client.Database("Task_manager").Collection("users")

	//check if username exists in the db
	userExists := userTable.FindOne(c, bson.D{{Key: "name", Value: user.Name}})
	var result models.User

	if userExists.Decode(&result) != mongo.ErrNoDocuments {
		c.IndentedJSON(400, gin.H{"error": "Username already registered"})
		return

	}

	count, err := userTable.CountDocuments(c, bson.D{})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Internal server error"})
		return
	}
	role := "user"

	//if there are no users assign the first user to the role of admin
	if count == 0 {
		role = "admin"
	}

	newUser := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "passwordHash", Value: user.PasswordHash},
		{Key: "role", Value: role},
	}

	_, err = userTable.InsertOne(c, newUser)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.IndentedJSON(200, gin.H{"success": "User registered successfully"})

}

func Login(c *gin.Context) {
	var user models.User
	var userCheck models.User


	//accept user info
	err := c.ShouldBind(&user)

	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "invalid request payload"})
		return
	}


	//connect to db
	client := ConnectDB(c)

	userTable := client.Database("Task_manager").Collection("users")

	//check if username exists in db
	filter := bson.D{
		{Key: "name", Value: user.Name},
	}

	userExist := userTable.FindOne(c, filter)
	
	err = userExist.Decode(&userCheck)

	
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(400, gin.H{"error": "User does not exist"})
		return
	}

	//check if password matches
	if bcrypt.CompareHashAndPassword([]byte(userCheck.PasswordHash), []byte(user.PasswordHash)) != nil {
		c.IndentedJSON(400, gin.H{"error": " passsword incorrect"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Internal server error"})
		return
	}

	user.PasswordHash = string(hashedPassword)

	// create a token with user role and expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"role": userCheck.Role,
		"expiration": time.Now().Add(5*time.Minute),
	})

	jwtToken, err := token.SignedString(Jwtsecret)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Internal server error"})

	}
	

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
}


func Promote(c *gin.Context, username string) {
	

	client := ConnectDB(c)
	defer DisconnectDB(c, client)

	collection := client.Database("Task_manager").Collection("users")

	filter := bson.D{{Key: "name",Value: username}}

	//check if user exists
	updated := bson.M{"$set":bson.M{"role": "admin"},}

	up, err := collection.UpdateOne(c,filter,updated)
	if err != nil{
		c.IndentedJSON(500, gin.H{"error":"Internal server error"})
		return
	}

	//check if the user role has been updated properly
	if up.MatchedCount== 1 &&  up.ModifiedCount == 0 {
		c.IndentedJSON(500, gin.H{"failure":"user is already promoted"})
		return
	}else if up.MatchedCount == 0 {
		c.IndentedJSON(500, gin.H{"failure":"user not found"})
		return
	}
	
	c.IndentedJSON(200, gin.H{"success":"user promoted successfully"})

}