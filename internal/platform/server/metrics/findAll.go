package metrics

import (
	"github.com/ferminhg/be-dora-metrics/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		metricsRepository := domain.NewInMemoryMetricRepository([]domain.Metric{})
		metrics, err := metricsRepository.FindAll()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, metrics)
	}
}
