package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radifan9/minitask-w10/internal/models"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// Setup Routing
	InitUserRegisterRouter(router, db)
	InitUserLoginRouter(router, db)

	// Catch all route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Wrong route",
			Status:  "Route not found",
		})
	})

	return router
}
