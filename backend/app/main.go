package main

import (
	"flehmen-api/ent"
	"log"
	"net/http"
	"os"
	"time"

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

type SukipiRequest struct {
	Name        string     `json:"name" validate:"required"`
	Weight      *float64   `json:"weight"`
	Height      *float64   `json:"height"`
	XID         *string    `json:"x_id"`
	InstagramID *string    `json:"instagram_id"`
	Hobby       *string    `json:"hobby"`
	Birthday    *time.Time `json:"birthday"`
	Family      *string    `json:"family"`
	IsMale      bool       `json:"is_male" validate:"required"`
	StartAt     time.Time  `json:"start_at" validate:"required"`
	MbtiId      *int       `json:"mbti_id"`
}

func (controller *Controller) SaveSukipi(c echo.Context) error {
	req := new(SukipiRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sukipi, err := controller.entClient.Sukipi.
		Create().
		SetName(req.Name).
		SetNillableWeight(req.Weight).
		SetNillableHeight(req.Height).
		SetNillableXID(req.XID).
		SetNillableInstagramID(req.InstagramID).
		SetNillableHobby(req.Hobby).
		SetNillableBirthday(req.Birthday).
		SetNillableFamily(req.Family).
		SetIsMale(req.IsMale).
		SetStartAt(req.StartAt).
		SetNillableMbtiID(req.MbtiId).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sukipi)
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
	e.POST("/sukipi", controller.SaveSukipi)

	e.Logger.Fatal(e.Start(":8080"))
}
