package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"test.com/gormtest/dao"
)

func SaveUser(c *gin.Context) {

	user := &dao.User{
		Username:   "test",
		Password:   "test",
		CreateTime: time.Now().UnixMilli(),
	}

	dao.SaveUser(user)
	c.JSON(200, user)
}

func GetUserById(c *gin.Context) {
	user := dao.GetUserById(1)
	c.JSON(200, user)
}

func GetAll(c *gin.Context) {
	user := dao.GetAll()
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateUser(1)
	c.JSON(200, gin.H{})
}

func DeleteUser(c *gin.Context) {
	dao.DeleteUser(1)
	c.JSON(200, gin.H{})
}
