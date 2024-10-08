package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/controllers/auth"
	"github.com/krystian-dajka/Mend-test-todolist/controllers/todo"
	"github.com/krystian-dajka/Mend-test-todolist/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter //
// creates gin router and defines routes
func SetupRouter(router *gin.Engine, client *mongo.Client) {
	// routes for todo CRUD
	todos := router.Group("/api/v1/todos")
	{
		todos.GET("", func(c *gin.Context) {
			todo.GetAllTodos(c, client)
		})
		todos.POST("", func(c *gin.Context) {
			todo.CreateTodo(c, client)
		})
		todos.DELETE("/:id", func(c *gin.Context) {
			todo.DeleteTodo(c, client)
		})
		todos.PUT("/:id", func(c *gin.Context) {
			todo.UpdateTodo(c, client)
		})
	}

	// routes for authorization and authentication
	authentication := router.Group("/api/v1/auth")
	{
		authentication.POST("/login", func(c *gin.Context) {
			auth.Login(c, client)
		})
		authentication.POST("/register", func(c *gin.Context) {
			auth.Register(c, client)
		})
		authentication.GET("/getMe", middleware.ProtectedRoute(), func(c *gin.Context) {
			auth.GetMe(c, client)
		})
		authentication.GET("/logout", auth.Logout)
	}

}
