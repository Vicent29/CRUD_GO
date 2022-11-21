package Routes

import (
	"CRUD-api/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)

		grp1.GET("table", Controllers.GetTables)
		grp1.POST("table", Controllers.CreateTable)
		grp1.GET("table/:id", Controllers.GetTableByID)
		grp1.PUT("table/:id", Controllers.UpdateTable)
		grp1.DELETE("table/:id", Controllers.DeleteTable)
	}
	return r
}
