package controller

import (
	"net/http"
	"strconv"

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
	g.POST("/contact", r.InsertContact)
	g.PUT("/contact/:id", r.UpdateContact)
	g.DELETE("/contact/:id", r.DeleteContact)
}

func (r *Rest) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	errLib := r.service.DeleteContact(c, id)
	if errLib != nil {
		c.JSON(errLib.StatusCode, gin.H{"success": false, "error": errLib})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "OK"})
}

func (r *Rest) InsertContact(c *gin.Context) {
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

func (r *Rest) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var req model.Contacts
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	req.Id = id
	errLib := r.service.UpdateContact(c, &req)
	if errLib != nil {
		c.JSON(errLib.StatusCode, gin.H{"success": false, "error": errLib})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "OK"})
}

func (r *Rest) GetContacts(c *gin.Context) {
	var limit int = 5
	var offset int = 0
	var er error
	limitStr, ok := c.GetQuery("limit")
	if ok {
		limit, er = strconv.Atoi(limitStr)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid limit value"})
		}
	}
	offsetStr, ok := c.GetQuery("")
	if ok {
		offset, er = strconv.Atoi(offsetStr)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid offset value"})
		}
	}
	contacts, err := r.service.GetContacts(c, limit, offset)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"success": false, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": contacts})
}
