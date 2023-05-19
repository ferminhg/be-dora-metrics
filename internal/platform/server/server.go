package server

import (
	"fmt"
	"github.com/ferminhg/be-dora-metrics/internal/domain"
	"github.com/ferminhg/be-dora-metrics/internal/platform/server/metrics"
	"log"

	"github.com/ferminhg/be-dora-metrics/internal/platform/server/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr         string
	engine           *gin.Engine
	metricRepository domain.MetricRepository
}

func New(host string, port uint, metricRepository domain.MetricRepository) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		metricRepository: metricRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.GET("/metrics", metrics.FindAllHandler(s.metricRepository))
}
