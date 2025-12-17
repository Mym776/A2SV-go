package repositories

import (
	
	entities "taskManager/Domain"

	services "taskManager/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

var MyUrl = "mongodb://localhost:27017"

type UserRepository struct {
	Context *gin.Context
}



func (u *UserRepository)Register(user entities.User)bool{
	c:=u.Context
	Users := ConnectDB(c,"Users")
	defer DisconnectDB(c,Users.Database().Client())

	userExists := Users.FindOne(c, bson.D{{Key: "name", Value: user.Name}})
	var result entities.User
	
	if userExists.Decode(&result) != mongo.ErrNoDocuments {
		return false

	}

	count, err := Users.CountDocuments(c, bson.D{})
	if err != nil {
		return false
	}
	role := "user"

	//if there are no users assign the first user to the role of admin
	if count == 0 {
		role = "admin"
	}

	newUser := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "passwordHash", Value: services.HashPassword(user.PasswordHash)},
		{Key: "role", Value: role},
	}

	_, err = Users.InsertOne(c, newUser)

	if err != nil {
		return false
	}

	return true

}

func(u *UserRepository) Login(user entities.User) bool{
	c:=u.Context
	Users := ConnectDB(c,"Users")
	defer DisconnectDB(c,Users.Database().Client())


	var userCheck entities.User

	filter := bson.D{
		{Key: "name", Value: user.Name},
	}

	userExist := Users.FindOne(c, filter)
	
	err := userExist.Decode(&userCheck)
	
	if err == mongo.ErrNoDocuments {
		return false
	}

	//check if password matches
	// if bcrypt.CompareHashAndPassword([]byte(userCheck.PasswordHash), []byte(user.PasswordHash)) != nil {
	// 	c.IndentedJSON(400, gin.H{"error": " passsword incorrect"})
	// 	return
	// }
	

	if !services.ComparePassword(user.PasswordHash,userCheck.PasswordHash){
		return false
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.IndentedJSON(500, gin.H{"error": "Internal server error"})
	// 	return
	// }

	// user.PasswordHash = string(hashedPassword)

	// create a token with user role and expiration time
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"name": user.Name,
	// 	"role": userCheck.Role,
	// 	"expiration": time.Now().Add(5*time.Minute),
	// })

	// jwtToken, err := token.SignedString(Jwtsecret)

	jwtToken := services.GenerateToken(userCheck)
	if err != nil {
		return false
	}
	

	c.JSON(200, gin.H{"message": "User logged in successfully","token":jwtToken})
	return true
}

func(u *UserRepository) Promote(username string) bool{
	c:=u.Context
	Users := ConnectDB(c,"Users")
	defer DisconnectDB(c,Users.Database().Client())

	
	filter := bson.D{{Key: "name",Value: username}}

	//check if user exists
	updated := bson.M{"$set":bson.M{"role": "admin"},}

	up, err := Users.UpdateOne(c,filter,updated)
	if err != nil{
		return false
	}

	//check if the user role has been updated properly
	if up.MatchedCount== 1 &&  up.ModifiedCount == 0 {
		return false
	}else if up.MatchedCount == 0 {
		return false
	}
	
	return true
}