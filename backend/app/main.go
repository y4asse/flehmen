package main

import (
	"bytes"
	"encoding/json"
	"flehmen-api/ent"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	XID         *string    `json:"twitterId"`
	Hobby       *string    `json:"hobby"`
	Birthday    *time.Time `json:"birthday"`
	ShoesSize   *float64   `json:"shoesSize"`
	Family      *string    `json:"family"`
	LikedAt     time.Time  `json:"likedAt"`
	MbtiId      *int       `json:"mbtiId"`
	NearStation *string    `json:"nearStation"`
}

func (controller *Controller) GetSukipiById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sukipi, err := controller.entClient.Sukipi.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sukipi)
}

func (controller *Controller) SaveSukipi(c echo.Context) error {
	req := new(SukipiRequest)
	if err := c.Bind(req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	sukipi, err := controller.entClient.Sukipi.
		Create().
		SetName(req.Name).
		SetNillableWeight(req.Weight).
		SetNillableHeight(req.Height).
		SetNillableXID(req.XID).
		SetNillableHobby(req.Hobby).
		SetNillableBirthday(req.Birthday).
		SetNillableShoesSize(req.ShoesSize).
		SetNillableFamily(req.Family).
		SetLikedAt(req.LikedAt).
		SetNillableMbtiID(req.MbtiId).
		SetNillableNearlyStation(req.NearStation).
		Save(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sukipi)
}

type SukipiVoiceResponse struct {
	VoiceID              string `json:"voice_id"`
	RequiresVerification bool   `json:"requires_verification"`
}

func (controller *Controller) CreateSukipiVoice(c echo.Context) error {
	voiceFile, err := c.FormFile("voice")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "voice file is required"})
	}

	file, err := voiceFile.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to open the file"})
	}
	defer file.Close()

	url := "https://api.elevenlabs.io/v1/voices/add"

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	now := time.Now().Format("2006-01-02 15:04:05")
	_ = writer.WriteField("name", now)
	_ = writer.WriteField("description", "from flehmen-api")
	_ = writer.WriteField("remove_background_noise", "true")

	part, err := writer.CreateFormFile("files", voiceFile.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create form file"})
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to copy file content"})
	}

	err = writer.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to close writer"})
	}

	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create request"})
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("xi-api-key", os.Getenv("ELEVENLABS_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to send request"})
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read response"})
	}

	voiceResponse := new(SukipiVoiceResponse)
	err = json.Unmarshal(body, &voiceResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(voiceResponse)

	url = fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", voiceResponse.VoiceID)

	data := map[string]string{
		"text":     "やせ、今日もいちにちお疲れ様。今日もがんばったね。おやすみ",
		"model_id": "eleven_multilingual_v2",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to marshal json"})
	}
	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("xi-api-key", os.Getenv("ELEVENLABS_API_KEY"))

	res, err = client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to send request"})
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read response"})
	}

	return c.Blob(http.StatusOK, "audio/mpeg", body)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_HOST"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
	}

	entClient, err := ent.Open(dialect.MySQL, c.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer entClient.Close()

	controller := &Controller{entClient: entClient}

	e.GET("/ok", controller.Ok)
	e.GET("/universities", controller.GetUniversities)
	e.GET("/sukipi/:id", controller.GetSukipiById)
	e.POST("/sukipi", controller.SaveSukipi)
	e.POST("/sukipi_voice", controller.CreateSukipiVoice)

	e.Logger.Fatal(e.Start(":8080"))
}
