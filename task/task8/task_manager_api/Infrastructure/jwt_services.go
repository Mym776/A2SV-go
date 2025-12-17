package services

import (
	"fmt"
	"log"
	entities "taskManager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtsecret = []byte("jwtsecret")

func GenerateToken(user entities.User) string{
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":       user.Name,
		"role":       user.Role,
		"expiration": time.Now().Add(5 * time.Minute),
	})
	jwtToken, err := token.SignedString(jwtsecret)
	if err != nil{
		log.Fatal("ERROR: JWT_GENERATION",err.Error())
	}
	return jwtToken

}

func VerifyToken(t *jwt.Token) (interface{}, error){
		
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		
		return jwtsecret, nil
}
