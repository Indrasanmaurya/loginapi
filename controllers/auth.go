package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"loginapi/models"
	"loginapi/utils"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 🟡 Step 1: JSON body ko bind karo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 🔍 Step 2: User match karo
	for _, user := range models.Users {
		if user.Email == input.Email {
			// 🔐 Step 3: Password compare karo
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
			if err == nil {
				// ✅ Step 4: JWT token banao
				token, _ := utils.GenerateJWT(user.Email)
				c.JSON(http.StatusOK, gin.H{
					"message": "Login successful",
					"token":   token,
				})
				return
			}
		}
	}

	// ❌ Step 5: Agar user ya password galat ho
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}
