package controller

import (
	"net/http"

	"github.com/fazarrahman/contact-app/model"
	"github.com/fazarrahman/contact-app/service"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	service *service.Svc
}

// New ...
func New(_svc *service.Svc) *Rest {
	return &Rest{service: _svc}
}

// Register ...
func (r *Rest) Register(g *gin.RouterGroup) {
	g.GET("/contact", r.GetContacts)
	g.POST("/contact", r.InsertContacts)
}

func (r *Rest) InsertContacts(c *gin.Context) {
	var req model.Contacts
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	errLib := r.service.InsertContact(c, &req)
	if errLib != nil {
		c.JSON(errLib.StatusCode, gin.H{"success": false, "error": errLib})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "OK"})
}

func (r *Rest) GetContacts(c *gin.Context) {
	contacts, err := r.service.GetContacts(c)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"success": false, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": contacts})
}
