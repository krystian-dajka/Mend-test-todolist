package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/krystian-dajka/Mend-test-todolist/models"
	"github.com/krystian-dajka/Mend-test-todolist/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Login //
// @desc Login in user with provided credentials
// @route POST /api/v1/auth/login
// @access Public
func Login(c *gin.Context, client *mongo.Client) {

	credentials := models.UserCred{}
	bindErr := c.ShouldBindJSON(&credentials)
	if bindErr != nil {
		log.Fatal(bindErr)
	}

	result := models.UserDB{}

	dbName := os.Getenv("DB_NAME")
	usersCollection := client.Database(dbName).Collection("users")

	// query for the user email because that is unique
	findOneErr := usersCollection.FindOne(c.Request.Context(), bson.M{
		"email": credentials.Email,
	}).Decode(&result)
	// if query error respond with wrong email
	if findOneErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That email does not exist",
		})
		return
	}

	compErr := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(credentials.Password))
	// if there is an error the provided password was incorrect else sign in the user and respond with cookie
	if compErr != nil {
		c.JSON(401, util.ResMessage{
			Success: false,
			Message: "Incorrect password",
		})
		return
	}

	token, getSignedErr := credentials.GetSignedJWT(result.ID.Hex())
	// jwt errror, should be rare but needs to return
	if getSignedErr != nil {
		c.JSON(400, util.ResError{
			Success: false,
			Error:   getSignedErr,
		})
		return
	}

	// secure cookie unless in development env
	secure := true
	if os.Getenv("GIN_ENV") == "development" {
		secure = false
	}

	// strict for csrf safety
	c.SetSameSite(http.SameSiteStrictMode)

	c.SetCookie("token", token, 2000, "/", "", secure, true)

	c.JSON(200, util.ResUser{
		Success: true,
		Message: models.UserRes{
			Name:  result.Name,
			Email: result.Email,
		},
	})
}
