package controllers

import (
	"strconv"

	"github.com/daintree-henry/msa-containerapp-go-itemapi/domain"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "OK")
}

func ListItem(c *gin.Context) {
	var item domain.Item
	items, err := item.List()
	if err != "" {
		c.JSON(500, err)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(201, items)
}

func GetItem(c *gin.Context) {
	var item domain.Item

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	item.GetById(int64(id))
	c.JSON(200, item)
}

func CreateItem(c *gin.Context) {
	var item domain.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, err)
		return
	}
	created, err := item.Create()
	if err != "" {
		c.JSON(500, err)
		return
	}
	c.JSON(200, created)
}

func DeleteItem(c *gin.Context) {
	var item domain.Item

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	result := item.DeleteById(int64(id))
	c.JSON(200, result)
}
