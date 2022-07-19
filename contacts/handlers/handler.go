package handlers

import (
	"contacts/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseStatus = struct {
	Status  string
	Code    string
	Message string
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := ResponseStatus{Message: "pong", Status: "Success", Code: "200"}
		c.JSON(http.StatusOK, response)
	}
}

func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := ResponseStatus{Message: "OK", Status: "Success", Code: "200"}
		c.JSON(http.StatusOK, response)
	}
}

func Greet(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, msg)
	}
}

func SamplePost() gin.HandlerFunc {
	return func(c *gin.Context) {

		contact := &models.Contact{}
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = json.Unmarshal(buf, contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = contact.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, contact)
	}
}
