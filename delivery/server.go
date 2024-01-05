package delivery

import (
	"database/sql"
	"gotodo-app/config"
	"gotodo-app/delivery/controller"
	"gotodo-app/delivery/middleware"
	"gotodo-app/repository"
	"gotodo-app/shared/service"
	"gotodo-app/usecase"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	authorUc   usecase.AuthorUseCase
	authUsc    usecase.AuthUseCase
	taskUc     usecase.TaskUseCase
	engine     *gin.Engine
	jwtService service.JwtService
	host       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthorController(s.authorUc, rg, authMiddleware).Route()
	controller.NewTaskController(s.taskUc, rg).Route()
	controller.NewAuthController(s.authUsc, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}
	// Inject DB ke -> repository
	authorRepo := repository.NewAuthorRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	// Inject REPO ke -> useCase
	authorUC := usecase.NewAuthorUseCase(authorRepo)
	taskUC := usecase.NewTaskUseCase(taskRepo, authorUC)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUc := usecase.NewAuthUseCase(authorUC, jwtService)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		authorUc:   authorUC,
		authUsc:    authUc,
		taskUc:     taskUC,
		engine:     engine,
		host:       host,
		jwtService: jwtService,
	}
}
