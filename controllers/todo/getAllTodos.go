package todo

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/models"
	"github.com/krystian-dajka/Mend-test-todolist/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTodos //
// @desc Get all todos
// @route GET /api/v1/todos
// @access Private
func GetAllTodos(c *gin.Context, client *mongo.Client) {

	todos := []models.Todo{}

	dbName := os.Getenv("DB_NAME")
	todosCollection := client.Database(dbName).Collection("todos")

	// query db and filter by user id
	cursor, findErr := todosCollection.Find(c.Request.Context(), bson.M{
		"user": c.Keys["id"],
	})
	if findErr != nil {
		c.JSON(500, util.ResError{
			Success: false,
			Error:   findErr,
		})
		return
	}

	// loop through cursor and put todos in the todos slice of todos
	cursorErr := cursor.All(c.Request.Context(), &todos)
	if cursorErr != nil {
		c.JSON(500, util.ResError{
			Success: false,
			Error:   cursorErr,
		})
		return
	}

	c.JSON(200, util.ResTodos{
		Success: true,
		Message: todos,
	})
}
