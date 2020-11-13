package infrastructure

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} method=${method} path="${uri}" host=${host} status=${status}` + "\n",
	}))

	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)

	e.GET("/", hello)

	return &Server{echo: e}
}

func (s *Server) Run() {
	s.echo.Logger.Fatal(s.echo.Start(":" + os.Getenv("PORT")))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is Jehanne Task Board.")
}
