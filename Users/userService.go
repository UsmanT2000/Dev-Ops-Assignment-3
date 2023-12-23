package Users

import (
	"github.com/UsmanT2000/ginAPIs/db"
	"github.com/UsmanT2000/ginAPIs/models"
	"github.com/UsmanT2000/ginAPIs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsers Get Users Stored in the Database
func GetUsers(ctx *gin.Context) {
	var users []models.UserInfo

	if err := db.DB.Find(&users).Error; err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": err.Error()}, http.StatusNotFound)
		return
	}

	utils.SendJSONResponse(ctx, gin.H{"user": users}, http.StatusOK)
}
func GetUserById(ctx *gin.Context) {
	var users models.UserInfo
	if err := db.DB.Where("id= ?", ctx.Param("id")).First(&users).Error; err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": "User not found"}, http.StatusNotFound)
		return
	}

	utils.SendJSONResponse(ctx, gin.H{"user": users}, http.StatusOK)
}

// Create New User
func CreateUser(ctx *gin.Context) {
	var user models.UserInfo
	if err := ctx.BindJSON(&user); err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	newUser := models.UserInfo{Name: user.Name, Age: user.Age, Email: user.Email}

	if err := db.DB.Create(&newUser).Error; err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": err.Error()}, http.StatusBadRequest)
	}

	utils.SendJSONResponse(ctx, gin.H{"Newly Added": newUser}, http.StatusOK)
}
func UpdateUser(ctx *gin.Context) {
	var user models.UserInfo

	if err := db.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": "User not found"}, http.StatusNotFound)
		return
	}

	var updateUser models.UserInfo

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only the fields that are provided in the updateUser object
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}

	if updateUser.Age != 0 {
		user.Age = updateUser.Age
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	if err := db.DB.Save(&user).Error; err != nil {
		utils.SendJSONResponse(ctx, gin.H{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(ctx, user, http.StatusOK)
}
func DeleteUser(c *gin.Context) {
	var users models.UserInfo

	if err := db.DB.Delete(&users, c.Param("id")).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) //If the id is found, delete the user
}
