package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/radifan9/minitask-w10/internal/configs"
	"github.com/radifan9/minitask-w10/internal/models"
)

var users = make(map[string]models.User)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env\nCause: ", err.Error())
		return
	}

	// DB initialization

	db, err := configs.InitDB()
	if err != nil {
		log.Println("failed to connect to database\nCause: ", err.Error())
		return
	}
	defer db.Close()

	// Test DB connection
	if err := configs.TestDBCon(db); err != nil {
		log.Println("ping to DB failed\nCause: ", err.Error())
		return
	}

	// Inisiasi engine gin
	router := gin.Default()

	// Catch all route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Wrong route",
			Status:  "Route not found",
		})
	})

	// Get all registered user
	router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"users":   users,
		})
	})

	// --- --- Register --- ---
	router.POST("/register", func(ctx *gin.Context) {
		user := models.User{}

		// Data-binding, masukkan
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// If there's no error in data-binding, validate the password
		if err := ValidatePassword(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// Build Query
		sql := "insert into users (email, password) values ($1, $2) returning id, email"
		values := []any{user.Email, user.Password}

		// Returning variable
		var returnedUser models.User
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&returnedUser.Id, &returnedUser.Email); err != nil {
			// Scanning error
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"body":    []any{},
			})
			return
		}

		// If Register successful
		// users[user.Email] = user // Store user with email as key
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"body":    returnedUser,
		})

	})

	// --- --- Login --- --
	router.POST("/login", func(ctx *gin.Context) {
		var loginUser models.User
		if err := ctx.ShouldBind(&loginUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// Query for checking if user with an email is exist
		sql := "select id, email from users where email = $1"
		values := []any{loginUser.Email}

		// Variable return
		var returnedUser models.User
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&returnedUser.Id, &returnedUser.Email); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				// When email is not registered
				"success": false,
				"body":    []any{},
			})
			return
		}

		sqlGetPass := "select email, password from users where id = $1"
		valuesGetPass := []any{returnedUser.Id}

		var userDB models.User
		if err := db.QueryRow(ctx.Request.Context(), sqlGetPass, valuesGetPass...).Scan(&userDB.Email, &userDB.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"body":    []any{},
			})
			return
		}

		// Check if inputed email pass, match the DB
		if loginUser.Password == userDB.Password {
			// Login successful
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login successful",
			})
		} else {
			// Login unsuccesful
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Login failed",
			})
		}

	})

	// Run Engine Gin
	router.Run(":8080")
}

var (
	reMin8           = regexp.MustCompile(`^.{8,}$`)
	reMinSmall       = regexp.MustCompile(`[a-z]`)
	reMinLarge       = regexp.MustCompile(`[A-Z]`)
	reMinSpecialChar = regexp.MustCompile(`[!@#$%^&*/()]`)
)

func ValidatePassword(user models.User) error {
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
