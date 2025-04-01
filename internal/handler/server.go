package handler

import (
	"github.com/RakhimovAns/Person-Service/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	cfg     *config.Config
	handler *PersonHandler
	router  *gin.Engine
}

func NewServer(cfg *config.Config, handler *PersonHandler) *Server {
	router := gin.Default()
	router.Use(cors.Default())

	server := &Server{
		cfg:     cfg,
		handler: handler,
		router:  router,
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"), // Полный URL
		ginSwagger.DefaultModelsExpandDepth(-1),                  // Скрываем Models
	))
	server.configureRouter()
	return server
}

func (s *Server) configureRouter() {

	s.handler.RegisterRoutes(s.router)
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.Port)
}
