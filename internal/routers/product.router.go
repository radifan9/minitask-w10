package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/radifan9/minitask-w10/internal/models"
)

func InitProductUpdateRouter(router *gin.Engine, db *pgxpool.Pool) {
	router.PATCH("/product/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		// Bind the body
		var productprice models.UpdatePriceRequest
		if err := ctx.ShouldBind(&productprice); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		// Build query
		query := "update products set price = $1, updated_at = current_timestamp where id = $2 returning id, price"
		values := []any{productprice.Price, id}

		var newProductPrice models.UpdatePriceRequest
		if err := db.QueryRow(ctx.Request.Context(), query, values...).Scan(&newProductPrice.Id, &newProductPrice.Price); err != nil {
			log.Println("Internal Server Error : ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    newProductPrice,
		})
	})
}
