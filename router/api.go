package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建一个路由组
func NewRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", listUsers)
		// api.POST("/users", createUser)
		// api.PUT("/users/:id", updateUser)
		// api.DELETE("/users/:id", deleteUser)
	}

}

func listUsers(c *gin.Context) {
	// 获取用户列表逻辑
	c.JSON(http.StatusOK, gin.H{"users": []string{"Alice", "Bob"}})
}
