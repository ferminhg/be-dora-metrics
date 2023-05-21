package metrics

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/ferminhg/be-dora-metrics/internal/domain"
)

func FindAllHandler(metricRepository domain.MetricRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		metrics, err := metricRepository.FindAll(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, metrics)
	}
}
