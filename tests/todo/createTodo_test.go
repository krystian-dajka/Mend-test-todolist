package todo_test

import (
	"os"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/controllers/todo"
	"github.com/krystian-dajka/Mend-test-todolist/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateTodo(t *testing.T) {
	DB_CONN_URL := "mongodb://root:krystian@localhost:27017"

	os.Setenv("DB_NAME", "TodosDB")
	os.Setenv("MONGO_URI", DB_CONN_URL)

	// Set up a test Gin router
	router := gin.Default()

	// Mock MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(DB_CONN_URL))
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	assert.NoError(t, err)
	defer client.Disconnect(ctx)

	// Mock request data
	newTodo := models.NewTodo{
		Title:       "Test Todo",
		Description: "Test Description",
		Done:        false,
	}

	jsonValue, _ := json.Marshal(newTodo)
	req, _ := http.NewRequest("POST", "/api/v1/todos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Register the CreateTodo route
	router.POST("/api/v1/todos", func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["id"] = "66f82efbfc6ccd38a00b9115"  // Mocking the user ID in the test
		todo.CreateTodo(c, client)
	})

	// Perform the request
	router.ServeHTTP(w, req)

	// // Check assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.True(t, response["success"].(bool))
	assert.Equal(t, newTodo.Title, response["message"].(map[string]interface{})["title"])
	assert.Equal(t, newTodo.Description, response["message"].(map[string]interface{})["description"])
}
