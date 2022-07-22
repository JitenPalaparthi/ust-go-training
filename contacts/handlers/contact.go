package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	IContact interfaces.IContact
}

func (ch *ContactHandler) CreateContact() gin.HandlerFunc {
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
		contact.Status = "created"
		contact.LastModified = fmt.Sprint(time.Now().Unix())
		result, err := ch.IContact.Create(contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func (ch *ContactHandler) GetContactByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		contact, err := ch.IContact.GetBy(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, contact)

	}
}

func (ch *ContactHandler) UpdateContactByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data := make(map[string]interface{})

		err = json.Unmarshal(buf, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rows, err := ch.IContact.UpdateBy(id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rows)

	}
}

func (ch *ContactHandler) DeleteContactByID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		rows, err := ch.IContact.DeleteBy(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, rows)
	}
}
