package auth_test

import (
	"os"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/krystian-dajka/Mend-test-todolist/config"
	"github.com/krystian-dajka/Mend-test-todolist/controllers/auth"
	"github.com/krystian-dajka/Mend-test-todolist/models"
)

func TestRegister(t *testing.T) {
	os.Setenv("DB_NAME", "TodosDB")
	os.Setenv("MONGO_URI", "mongodb://root:krystian@localhost:27017")

	router := gin.Default()

	// Mock request data
	newUser := models.UserCred{
        Name:     "Test_name",
        Email:    "test@example.com",
        Password: "password",
    }

	jsonValue, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	client := config.ConnectDB()

	router.POST("/api/v1/auth/register", func(c *gin.Context) {
		auth.Register(c, client)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}