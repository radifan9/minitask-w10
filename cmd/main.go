package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]User)

func main() {
	fmt.Println("--- --- Week 10 - Beginner Backend --- ---")

	// Inisiasi engine gin
	router := gin.Default()

	// Catch all route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Message: "Wrong route",
			Status:  "Route not found",
		})
	})

	// Register
	router.POST("/register", func(ctx *gin.Context) {
		user := User{}

		// Data-binding, masukkan
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// If there's no error in data-binding
		if err := ValidatePassword(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// If Register successful
		users[user.Email] = user // Store user with email as key
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"body":    user,
		})

	})

	router.POST("/login", func(ctx *gin.Context) {
		var loginUser User
		if err := ctx.ShouldBind(&loginUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// Check if user exists
		storedUser, exists := users[loginUser.Email]
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "invalid email or password",
				"success": false,
			})
			return
		}

		// Check password (in a real app, use bcrypt or similar)
		if storedUser.Password != loginUser.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "invalid email or password",
				"success": false,
			})
			return
		}

		// Login successful
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Login successful",
		})
	})

	// Run Engine Gin
	router.Run(":8080")
}

type Response struct {
	Message string
	Status  string
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	reMin8           = regexp.MustCompile(`^.{8,}$`)
	reMinSmall       = regexp.MustCompile(`[a-z]`)
	reMinLarge       = regexp.MustCompile(`[A-Z]`)
	reMinSpecialChar = regexp.MustCompile(`[!@#$%^&*/()]`)
)

func ValidatePassword(user User) error {
	fmt.Println(user.Email)
	fmt.Println(user.Password)

	if !reMin8.MatchString(user.Password) {
		return errors.New("password harus minimal 8 karakter")
	}
	if !reMinSmall.MatchString(user.Password) {
		return errors.New("password minimal harus 1 karakter kecil")
	}
	if !reMinLarge.MatchString(user.Password) {
		return errors.New("password minimal harus 1 karakter besar")
	}
	if !reMinSpecialChar.MatchString(user.Password) {
		return errors.New("password harus ada karakter spesial")
	}
	return nil
}
