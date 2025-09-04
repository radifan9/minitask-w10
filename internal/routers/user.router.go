package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radifan9/minitask-w10/internal/models"
	"github.com/radifan9/minitask-w10/internal/utils"
)

func InitUserRegisterRouter(router *gin.Engine, db *pgxpool.Pool) {
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
		if err := utils.ValidatePassword(user); err != nil {
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
}

func InitUserLoginRouter(router *gin.Engine, db *pgxpool.Pool) {
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
}
