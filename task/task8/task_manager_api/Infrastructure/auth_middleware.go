package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) (jwt.MapClaims, bool) {
	//middleware for aut
	// checks if the auth is in the proper format
	authHandler := c.GetHeader("Authorization")
	if authHandler == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return nil, false
	}

	authParts := strings.Split(authHandler, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.IndentedJSON(401, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return nil, false
	}
	//checks if the token is signed using the HMAC_SHA family of signing and returns the jwt secret/
	token, err := jwt.Parse(authParts[1], func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		
		return jwtsecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return nil, false
	}

	mps := token.Claims.(jwt.MapClaims)

	return mps, true

}

// auth method for low clearance access methods
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mps, ok := middleware(c)
		if ok != true {
			return
		}

		layout := "2006-01-02T15:04:05.999999999-07:00"
		lt := mps["expiration"].(string)
		tokenTime, err := time.Parse(layout, lt)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid time format"})
			c.Abort()
			return
		}

		if time.Now().After(tokenTime) {
			c.IndentedJSON(401, gin.H{"error": "expired token"})
			c.Abort()
			return
		}

		role := mps["role"].(string)
		if role != "user" && role != "admin" {
			c.IndentedJSON(401, gin.H{"error": "User not qualified to access this feature"})
			c.Abort()
			return
		}
	}

}

// auth method for admin level clearance methods
func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mps, ok := middleware(c)
		if ok != true {
			c.Abort()
			return
		}

		layout := "2006-01-02T15:04:05.999999999-07:00"
		lt := mps["expiration"].(string)
		tokenTime, err := time.Parse(layout, lt)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid time format"})
			c.Abort()
			return
		}

		if time.Now().After(tokenTime) {
			c.IndentedJSON(401, gin.H{"error": "expired token"})
			c.Abort()
			return
		}

		role := mps["role"].(string)
		if role != "admin" {
			c.IndentedJSON(401, gin.H{"error": "User not qualified to access this feature"})
			c.Abort()
			return
		}
	}

}
