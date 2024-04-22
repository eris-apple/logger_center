package app

import (
	config2 "github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/transport/rest"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redis2 "github.com/redis/go-redis/v9"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Router *gin.Engine
	Store  store.Store
	Redis  *redis2.Client
}

func newServer(config *config2.Config, store store.Store) *Server {
	s := &Server{
		Router: gin.New(),
		Store:  store,
		Redis: redis2.NewClient(&redis2.Options{
			Addr:     config.RedisURL,
			Password: config.RedisPwd,
		}),
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.Router.Use(cors.New(cors.Config{AllowAllOrigins: true}))
	s.Router.Use(s.setRequestID())
	s.Router.Use(s.logger())
	s.Router.Handle("GET", "/", HealthCheckHandler)

	ug := s.Router.Group("/users")
	{
		ug.Use(s.AuthRequired())

		userStore := s.Store.User()
		userService := services.NewUserService(userStore)
		userHandler := rest.NewUserHandler(userService)

		ug.GET("/", userHandler.FindAll)
		ug.GET("/:id", userHandler.FindById)
		ug.PUT("/:id", userHandler.Update)
		ug.DELETE("/:id", userHandler.Delete)
	}

	ig := s.Router.Group("/identity")
	{

		userStore := s.Store.User()
		sessionStore := s.Store.Session()
		identityService := services.NewIdentityService(userStore, sessionStore)
		identityHandler := rest.NewIdentityHandler(identityService)

		ig.POST("/sign-up", identityHandler.SignUp)
		ig.POST("/sign-in", identityHandler.SignIn)
		ig.GET("/check", s.AuthRequired(), identityHandler.Check)
	}

	pg := s.Router.Group("/projects")
	{
		pg.Use(s.AuthRequired())
		pg.Use(s.RoleRequired(enums.Admin, enums.Moderator))
		projectStore := s.Store.Project()
		projectService := services.NewProjectService(projectStore)
		projectHandler := rest.NewProjectHandler(projectService)

		pg.GET("/", projectHandler.FindAll)
		pg.GET("/:id", projectHandler.FindById)
		pg.POST("/", projectHandler.Create)
		pg.PUT("/:id", projectHandler.Update)
		//pg.DELETE("/:id", projectHandler.Create)
	}
}

func (s *Server) setRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.NewV4().String()
		ctx.Header("X-Request-ID", id)

		ctx.Next()
	}
}

func (s *Server) logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		// access the status we are sending
		url := ctx.Request.URL
		method := ctx.Request.Method
		body := ctx.Request.Body
		status := ctx.Writer.Status()
		log.Println(status, "—", method, url, body, "—", latency)
	}
}

type healthCheck struct {
	status string
}

func HealthCheckHandler(ctx *gin.Context) {
	response := healthCheck{
		status: "ok",
	}

	ctx.JSON(200, response)
}

func (s *Server) AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponseHandler(ctx, http.StatusUnauthorized, config2.ErrForbiddenAccess, nil)
			ctx.Abort()
			return
		}

		authHeaderArray := strings.Split(authHeader, " ")
		if len(authHeaderArray) != 2 {
			utils.ErrorResponseHandler(ctx, http.StatusUnauthorized, config2.ErrForbiddenAccess, nil)
			ctx.Abort()
			return
		}

		ss, SSErr := s.Store.Session().FindByToken(authHeaderArray[1])
		if SSErr != nil {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config2.ErrForbiddenAccess, nil)
			ctx.Abort()
			return
		}

		if ss.IsActive == false {
			utils.ErrorResponseHandler(ctx, http.StatusUnauthorized, config2.ErrSessionExpired, nil)
			ctx.Abort()
			return
		}

		user, UErr := s.Store.User().FindById(ss.UserID)
		if UErr != nil {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config2.ErrForbiddenAccess, nil)
			ctx.Abort()
			return
		}

		user.Sanitize()

		ctx.Set("user", user)
		ctx.Next()
	}
}

func (s *Server) RoleRequired(roles ...enums.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.Value("user").(*models.User)
		if user == nil {
			utils.ErrorResponseHandler(ctx, http.StatusForbidden, config2.ErrForbiddenAccess, nil)
		}

		for _, role := range roles {
			if user.Role == role.String() {
				ctx.Next()
				return
			}
		}

		utils.ErrorResponseHandler(ctx, http.StatusForbidden, config2.ErrForbiddenAccess, nil)
	}
}
