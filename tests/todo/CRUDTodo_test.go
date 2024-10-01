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
	"github.com/krystian-dajka/Mend-test-todolist/util"
	"github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// mock MongoDB client for testing

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestUpdateTodo(t *testing.T) {
	DB_CONN_URL := "mongodb://root:krystian@localhost:27017"

	os.Setenv("DB_NAME", "TodosDB")
	os.Setenv("MONGO_URI", DB_CONN_URL)

	router := setupRouter()

	// Mock MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(DB_CONN_URL))
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	assert.NoError(t, err)
	defer client.Disconnect(ctx)

	todoCollection := client.Database("TodosDB").Collection("todos")

	// Mock insert newTodo (old)
	newTodo := models.NewTodo{
		Title:       "Test Old Todo",
		Description: "Test Old Description",
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

	router.PUT("/api/v1/todos/:id", func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["id"] = "66f82efbfc6ccd38a00b9115"  // Mocking the user ID in the test
		todo.UpdateTodo(c, client)
	})

	router.DELETE("/api/v1/todos/:id", func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["id"] = "66f82efbfc6ccd38a00b9115"  // Mocking the user ID in the test
		todo.DeleteTodo(c, client)
	})

	// Perform the request
	router.ServeHTTP(w, req)

	// // Check assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var insert_resp util.ResTodo
	err = json.Unmarshal(w.Body.Bytes(), &insert_resp)
	assert.NoError(t, err)

	assert.True(t, insert_resp.Success)
	assert.Equal(t, newTodo.Title, insert_resp.Message.Title)
	assert.Equal(t, newTodo.Description, insert_resp.Message.Description)

	// Mock todoData updating
	insertedID := insert_resp.Message.ID

	// Define the update payload
	updatePayload := models.NewTodo{
		Title:     "Test New Title",
		Description:   "Test New Content",
		Done: true,
	}
	body, _ := json.Marshal(updatePayload)

	// Create a request to update the todo
	req, _ = http.NewRequest("PUT", "/api/v1/todos/"+insertedID.Hex(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Assert the status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode the response body
	var response util.ResTodo
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Validate the returned updated todo fields
	assert.True(t, response.Success)
	assert.Equal(t, "Test New Title", response.Message.Title)
	assert.Equal(t, "Test New Content", response.Message.Description)
	assert.Equal(t, true, response.Message.Done)

	// Check if the document is updated in the database
	var updatedTodo models.Todo
	err = todoCollection.FindOne(nil, bson.M{"_id": insertedID}).Decode(&updatedTodo)
	assert.NoError(t, err)
	assert.Equal(t, "Test New Title", updatedTodo.Title)
	assert.Equal(t, "Test New Content", updatedTodo.Description)
	assert.Equal(t, true, updatedTodo.Done)

	// Create a request to delete the todo
	req, _ = http.NewRequest("DELETE", "/api/v1/todos/"+insertedID.Hex(), nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Assert the status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}
