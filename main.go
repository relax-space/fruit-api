package main

import (
	"flag"
	"net/http"
	"os"

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
	v.PUT("/fruits", Fruit_Update)
	v.DELETE("/fruits/:Id", Fruit_Delete)
}

func InitDB(dialect, conn string) (newDb *xorm.Engine, err error) {
	newDb, err = xorm.NewEngine(dialect, conn)
	newDb.Sync(new(Fruit))
	newDb.ShowSQL(true)
	return
}
