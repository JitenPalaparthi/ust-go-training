package gin

import (
	"contacts/interfaces"
	"contacts/messagebroker"
	"contacts/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	IContact interfaces.IContact
	Brokers  []string
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
		// call producer
		// downside of this implementation is if the below code fails the above db operations cannot be
		// rolled back
		jsonstr, _ := contact.ToJson()
		msg := &messagebroker.Messaging{Brokers: ch.Brokers, Topic: "contacts-create"}
		msg.Key = []byte(strconv.Itoa(contact.ID))
		msg.Value = []byte(jsonstr)
		err = msg.Produce(c.Request.Context())
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
