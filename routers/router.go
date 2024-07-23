package routers

import (
	"crud-golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Auth
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/", handlers.AuthMiddleware())
	{
		// Users
		protected.POST("/users", handlers.CreateUser)
		protected.GET("/users", handlers.GetUsers)
		r.GET("/users/:id", handlers.GetUserByID)
		protected.PUT("/users-update/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)

		// Employee
		protected.POST("/employee", handlers.CreateEmployee)
		protected.GET("/employee", handlers.GetEmployee)
		protected.GET("/employee/:id", handlers.GetEmployeeByID)
		protected.PUT("/employee/:id", handlers.UpdateEmployee)
		protected.DELETE("/employee/:id", handlers.DeleteEmployee)

		// Role
		protected.POST("/role", handlers.CreateRole)
		protected.GET("/role", handlers.GetRole)
		protected.GET("/role/:id", handlers.GetRoleByID)
		protected.PUT("/role/:id", handlers.UpdateRole)
		protected.DELETE("/role/:id", handlers.DeleteRole)

		// Position
		protected.POST("/position", handlers.CreatePosition)
		protected.GET("/position", handlers.GetPosition)
		protected.GET("/position/:id", handlers.GetPositionByID)
		protected.PUT("/position/:id", handlers.UpdatePosition)
		protected.DELETE("/position/:id", handlers.DeletePosition)
	}
}
