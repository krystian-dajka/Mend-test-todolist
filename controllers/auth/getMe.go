package auth

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/models"
	"github.com/krystian-dajka/Mend-test-todolist/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetMe //
// @desc Used to tell the client if they are logged in, since the credencials cookies cannot be read
// @route GET /api/v1/auth/getMe
// @access Private
func GetMe(c *gin.Context, client *mongo.Client) {
	// This is a protected route by design, so if the request hits this func it is authenticated and should respond with 200. else it will get "not logged in" or "invalid token"

	id, objErr := primitive.ObjectIDFromHex(c.Keys["id"].(string))
	if objErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Bad cookie",
		})
	}

	// holds the result of the db query
	result := models.UserDB{}

	dbName := os.Getenv("DB_NAME")
	usersCollection := client.Database(dbName).Collection("users")

	findOneErr := usersCollection.FindOne(c.Request.Context(), bson.M{
		"_id": id,
	}).Decode(&result)
	if findOneErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Query error for id",
		})
		return
	}

	c.JSON(200, util.ResUser{
		Success: true,
		Message: models.UserRes{
			Name:  result.Name,
			Email: result.Email,
		},
	})
}
