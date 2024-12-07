package main

import (
	"flehmen-api/ent"
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	client *ent.Client
}

func (controller *Controller) Ok(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func main() {
	e := echo.New()
	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_HOST"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
	}

	client, err := ent.Open(dialect.MySQL, c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	controller := &Controller{client: client}

	e.GET("/ok", controller.Ok)

	e.Logger.Fatal(e.Start(":8080"))
}
