package server

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"userapi/handlers"
)

type Server struct {
	httpServer *http.Server
}

type Container struct {
	Config *Config
	DB     *sqlx.DB
}

//func (s *Server) Run(port string, handler http.Handler) error {
//	s.httpServer = &http.Server{
//		Addr:           ":" + port,
//		MaxHeaderBytes: 1 << 20,
//		Handler:        handler,
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//	}
//
//	logrus.Infof("listen and serve on port %s", port)
//
//	return s.httpServer.ListenAndServe()
//}

func Run(addr string) error {
	e := echo.New()

	routes(e)

	return e.Start(addr)
}

func routes(e *echo.Echo) {
	config := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := NewPostgresDB(config)

	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	container := Container{
		Config: &config,
		DB:     db,
	}

	handlers.NewAuthHandler(&container).SetRoutes(e.Group("/auth"))
	handlers.NewUsersHandler(&container).SetRoutes(e.Group("/users"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
