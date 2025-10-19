package router

import (
	"github.com/gin-gonic/gin"
	"test.com/gormtest/api"
)

func InitRouter(r *gin.Engine) {

	api.RegisRouter(r)

}
