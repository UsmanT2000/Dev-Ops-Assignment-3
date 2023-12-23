package middleware

import (
	"fmt"
	"github.com/UsmanT2000/ginAPIs/models"
	"github.com/UsmanT2000/ginAPIs/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	users = []models.User{
		{ID: 1, Username: "Usman", Password: "password1"},
		//{ID: 2, Username: "jane_smith", Password: "password2"},
		// Add more users as needed.
	}

	secretKey = []byte("12345")
	mutex     sync.Mutex
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &models.CustomClaims{}

		tokenString := c.GetHeader("Authorization")
		log.Println("The token is: ", tokenString)
		// Check if the token is present.
		if tokenString == "" {
			utils.SendJSONResponse(c, gin.H{"error": "No Token Present"}, http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Split the token string if it includes the "Bearer " prefix.
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		} else {
			utils.SendJSONResponse(c, gin.H{"error": "Invalid token format"}, http.StatusUnauthorized)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Invalid signing method")
			}
			return secretKey, nil
		})

		log.Println("The Secret Key is: ", secretKey)
		log.Println("Parsed token:", token)
		log.Println("Token parsing error:", err)

		if err != nil || !token.Valid {
			utils.SendJSONResponse(c, gin.H{"error": "Unauthorized token"}, http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Set the user ID and username in the context for further use.
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

func GenerateJWTToken(userID int64, username string) (string, error) {
	claims := models.CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day).
			Issuer:    "EMI",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	log.Println("Generated token:", tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func AuthenticateUser(username, password string) (*models.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, u := range users {
		if u.Username == username && u.Password == password {
			return &models.User{
				ID:       u.ID,
				Username: u.Username,
				Password: u.Password,
			}, nil
		}
	}
	return nil, fmt.Errorf("Invalid credentials")
}

func LoginEndpoint(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		utils.SendJSONResponse(c, gin.H{"error": "Bad Request"}, http.StatusBadRequest)
		return
	}

	// Authenticate the user based on the provided username and password.
	matchedUser, err := AuthenticateUser(u.Username, u.Password)
	if err != nil {
		utils.SendJSONResponse(c, gin.H{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}

	// If the user is authenticated, generate a JWT token for the user.
	token, err := GenerateJWTToken(matchedUser.ID, matchedUser.Username)
	if err != nil {
		utils.SendJSONResponse(c, gin.H{"error": "Failed to generate token"}, http.StatusInternalServerError)
		return
	}

	// Set the token in the Authorization header of the response.
	c.Header("Authorization", "Bearer "+token)

	// Send a success response.
	utils.SendJSONResponse(c, gin.H{"message": "Token generated and set in Authorization header: " + token}, http.StatusOK)
}

// Function for testing Endpoints
func AnalyticsEndpoint(c *gin.Context) {
	// Access the user ID and username from the context (set during JWT authentication).
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	// Generate sample analytics data (you can replace this with your actual analytics logic).
	analyticsData := gin.H{
		"user_id":   userID,
		"username":  username,
		"pageviews": 1000,
		"clicks":    500,
		"visitors":  300,
	}

	// Send the analytics data as the JSON response.
	utils.SendJSONResponse(c, analyticsData, http.StatusOK)
}
