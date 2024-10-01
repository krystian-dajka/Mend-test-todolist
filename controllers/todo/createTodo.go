package todo

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/models"
	"github.com/krystian-dajka/Mend-test-todolist/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTodo //
// @desc Create a new todo from todo struc and add it to the database
// @route POST /api/v1/todos
// @access Private
func CreateTodo(c *gin.Context, client *mongo.Client) {

	newTodo := models.NewTodo{}
	newTodo.SetCreatedAt()
	newTodo.SetUpdatedAt()

	bindErr := c.ShouldBindJSON(&newTodo)
	if bindErr != nil {
		c.JSON(400, util.ResError{
			Success: false,
			Error:   bindErr,
		})
		return
	}

	dbName := os.Getenv("DB_NAME")
	todosCollection := client.Database(dbName).Collection("todos")

	// put todo in the collection, returns _id
	todosRes, insertErr := todosCollection.InsertOne(c.Request.Context(), newTodo)
	if insertErr != nil {
		c.JSON(401, util.ResError{
			Success: false,
			Error:   insertErr,
		})
		return
	}

	// uses struct to format the response body. ID is from InserOne
	todo := models.Todo{
		ID:          todosRes.InsertedID.(primitive.ObjectID),
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Done:        newTodo.Done,
		CreatedAt:   newTodo.CreatedAt,
		UpdatedAt:   newTodo.UpdatedAt,
		User:        newTodo.User,
	}

	c.JSON(200, util.ResTodo{
		Success: true,
		Message: todo,
	})
}
