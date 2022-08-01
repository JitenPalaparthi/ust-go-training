package echo

import (
	"contacts/interfaces"
	"contacts/messagebroker"
	"contacts/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ContactHandler struct {
	IContact interfaces.IContact
	Brokers  []string
}

func (ch *ContactHandler) CreateContact() echo.HandlerFunc {
	return func(c echo.Context) error {
		contact := &models.Contact{}
		if err := c.Bind(contact); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})

		}
		if err := contact.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})

		}

		contact.Status = "created"
		contact.LastModified = fmt.Sprint(time.Now().Unix())
		result, err := ch.IContact.Create(contact)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		// call producer
		// downside of this implementation is if the below code fails the above db operations cannot be
		// rolled back
		jsonstr, _ := contact.ToJson()
		msg := &messagebroker.Messaging{Brokers: ch.Brokers, Topic: "contacts-create"}
		msg.Key = []byte(strconv.Itoa(contact.ID))
		msg.Value = []byte(jsonstr)
		err = msg.Produce(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, result)
	}
}

func (ch *ContactHandler) GetContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		contact, err := ch.IContact.GetBy(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, contact)
	}
}

func (ch *ContactHandler) DeleteContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		contact, err := ch.IContact.DeleteBy(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusNoContent, contact)
	}
}

func (ch *ContactHandler) UpdateContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		data := make(map[string]interface{})
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		rows, err := ch.IContact.UpdateBy(id, data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, rows)
	}
}
