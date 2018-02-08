package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/relax-space/go-kitt/echomiddleware"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appEnv  = flag.String("APP_ENV", os.Getenv("APP_ENV"), "APP_ENV")
	connEnv = flag.String("FRUIT_CONN", os.Getenv("FRUIT_CONN"), "FRUIT_CONN")
)

func main() {
	flag.Parse()
	envParam := &EnvParamDto{
		AppEnv:  *appEnv,
		ConnEnv: *connEnv,
	}
	InitEnv(envParam)
	db, err := InitDB("mysql", envParam.ConnEnv)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.CORS())
	RegisterApi(e)
	e.Use(echomiddleware.ContextDB(db))
	e.Start(":5000")

}
func RegisterApi(e *echo.Echo) {
	jwtEnv := flag.String("JWT_SECRET", os.Getenv("JWT_SECRET"), "JWT_SECRET")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "fruit-api")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	v := e.Group("/v1")
	v.GET("/fruits/:id", Fruit_Get)
	v.GET("/fruits", Fruit_Find)
	v.POST("/fruits", Fruit_Create)
	v.PUT("/fruits/:id", Fruit_Update)
	v.DELETE("/fruits/:id", Fruit_Delete)

	v2 := e.Group("/v2")
	v2.GET("/fruits/:id", Fruit_Get)
	v2.GET("/fruits", Fruit_Find)
	v2.POST("/fruits", Fruit_Create)
	v2.PUT("/fruits/:id", Fruit_Update)
	v2.DELETE("/fruits/:id", Fruit_Delete)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(*jwtEnv),
		Skipper: func(c echo.Context) bool {
			ignore := []string{
				"/ping",
				"/v1",
			}
			for _, i := range ignore {
				if strings.HasPrefix(c.Request().URL.Path, i) {
					return true
				}
			}

			return false
		},
	}))

}

func InitDB(dialect, conn string) (newDb *xorm.Engine, err error) {
	newDb, err = xorm.NewEngine(dialect, conn)
	newDb.Sync(new(Fruit))
	newDb.ShowSQL(true)
	return
}
