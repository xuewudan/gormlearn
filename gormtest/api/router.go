package api

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
	r.GET("/get", GetUserById)
	r.GET("/all", GetAll)
	r.GET("/update", UpdateUser)
	r.GET("/delete", DeleteUser)
}
