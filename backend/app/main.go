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
	entClient *ent.Client
}

func (controller *Controller) Ok(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (controller *Controller) GetUniversities(c echo.Context) error {
	universities, err := controller.entClient.University.Query().All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err) // TODO: errを返すのはセキュリティ的に問題がありそう
	}
	return c.JSON(http.StatusOK, universities)
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

	controller := &Controller{entClient: client}

	e.GET("/ok", controller.Ok)
	e.GET("/universities", controller.GetUniversities)

	e.Logger.Fatal(e.Start(":8080"))
}
