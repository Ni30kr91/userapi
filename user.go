package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserID   string `json:"user_id"`
	City     string `json:"city"`
	Password string `json:"password" binding:"required"`
}

var Data map[string]User

func main() {
	Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/user/:user_id", GetUserByUserID)
	r.GET("/user", GetAllUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:user_id", UpdateUser)
	r.DELETE("/user/:user_id", deleteuser)
}

//GetUserByUserID function
func GetUserByUserID(c *gin.Context) {
	//records := readCsvFile("./movies.csv")
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	user := getUserByID(userID)
	res := gin.H{
		"user": user,
	}
	c.JSON(http.StatusOK, res)
}

//GetAllUser function
func GetAllUser(c *gin.Context) {
	res := gin.H{
		"user": Data,
	}
	c.JSON(http.StatusOK, res)
}

//CreateUser POST
func CreateUser(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if UniqueEmail(reqBody.Email) {
		res := gin.H{
			"error": "Email is already Exist",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if isUnique(reqBody.UserID) {
		res := gin.H{
			"error": "User already Exist",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if len(reqBody.Phone) != 13 {
		res := gin.H{
			"error": "phone no must be 13 digits",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	Data[reqBody.UserID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
		"length":  len(reqBody.Phone),
	}
	c.JSON(http.StatusOK, res)
	return
}

//Update User PUT
func UpdateUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	reqBody := User{}

	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID == "" {
		res := gin.H{
			"error": "UserId must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if reqBody.UserID != userID {
		res := gin.H{
			"error": "UserId cannot be updated",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	Password := c.GetHeader("Password")
	if !cheakpassword(userID, Password) {
		res := gin.H{
			"error": "invalid password",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Data[userID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}

//Delete function
func deleteuser(c *gin.Context) {
	UserID, ok := c.Params.Get("user_id")
	fmt.Println(UserID)
	fmt.Println(Data)
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	reqBody := User{}

	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := delete_user(UserID)
	res := gin.H{
		"success": true,
		"message": result,
	}
	c.JSON(http.StatusOK, res)
	return
}
func cheakpassword(username, password string) bool {
	if Data[username].Password == password {
		return true
	} else {
		return false
	}

}
