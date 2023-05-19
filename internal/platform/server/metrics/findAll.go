package metrics

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, []int{})
	}
}
