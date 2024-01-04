package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Age       string `json:"age" bson:"age"`
}

var seedUsers = []User{
	{
		UserName:  "74e51f57-2f90-41e4-b6a0-16fda2e86fad",
		Email:     "theresamckenzie@slofast.com",
		FirstName: "Lilly",
		LastName:  "Sanford",
		Age:       "30",
	},
	{
		UserName:  "26fc9416-c2f4-4085-a99c-82fe7306929e",
		Email:     "lillysanford@slofast.com",
		FirstName: "Avery",
		LastName:  "Mclaughlin",
		Age:       "28",
	},
	{
		UserName:  "9918fa1f-a772-42ed-8543-080311286a54",
		Email:     "averymclaughlin@slofast.com",
		FirstName: "Reed",
		LastName:  "Martin",
		Age:       "29",
	},
	{
		UserName:  "c4e238fc-6d4b-473a-b676-11b0bb51d696",
		Email:     "reedmartin@slofast.com",
		FirstName: "Mia",
		LastName:  "Wallace",
		Age:       "34",
	},
	{
		UserName:  "22834a3f-4e42-489a-bf45-1f5892e0a6be",
		Email:     "miawallace@slofast.com",
		FirstName: "Lesley",
		LastName:  "Thomas",
		Age:       "37",
	},
	{
		UserName:  "d0cc7c30-b353-43cf-bbb6-68f81605a712",
		Email:     "lesleythomas@slofast.com",
		FirstName: "Therese",
		LastName:  "Nelson",
		Age:       "23",
	},
}

func getAllUsers(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, seedUsers)
}

func getUser(context *gin.Context) {

	userName := context.Param("username")
	for _, val := range seedUsers {
		if val.UserName == userName {
			context.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	context.IndentedJSON(http.StatusOK, "User not found with given id :"+userName)
}

func postUser(context *gin.Context) {
	var newUser User
	if err := context.BindJSON(&newUser); err != nil {

		fmt.Println(err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	fmt.Println(newUser)
	seedUsers = append(seedUsers, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func main() {

	var router *gin.Engine = gin.Default()
	router.GET("/users", getAllUsers)
	router.POST("/users", postUser)
	router.GET("/users/:username", getUser)
	router.Run(":5000")
}
