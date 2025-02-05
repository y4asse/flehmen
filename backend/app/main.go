package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flehmen-api/ent"
	"flehmen-api/ent/nextaction"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"context"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/openai/openai-go/option"
	"google.golang.org/api/option"
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
	text := c.FormValue("text")

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

	if res.StatusCode != http.StatusOK {
		// jsonの内容を返す
		body, _ := io.ReadAll(res.Body)
		fmt.Println(string(body))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "音声ファイルのアップロードに失敗しました"})
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read response"})
	}

	voiceResponse := new(SukipiVoiceResponse)
	err = json.Unmarshal(body, &voiceResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to parse ElevenLabs response"})
	}

	fmt.Println(voiceResponse)

	// ここから音声合成
	url = fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", voiceResponse.VoiceID)

	data := map[string]string{
		"text":     text,
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

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "音声合成に失敗しました"})
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read response"})
	}

	// 音声ファイルを消す
	url = fmt.Sprintf("https://api.elevenlabs.io/v1/voices/%s", voiceResponse.VoiceID)

	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create request"})
	}

	req.Header.Set("xi-api-key", os.Getenv("ELEVENLABS_API_KEY"))

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to send request"})
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "音声ファイルの削除に失敗しました"})
	}
	return c.Blob(http.StatusOK, "audio/mpeg", body)
}

type ClosenessScore struct {
	Time    int `json:"time"`
	Balance int `json:"balance"`
	Rhythm  int `json:"rhythm"`
	Type    int `json:"type"`
	Words   int `json:"words"`
	Total   int `json:"total"`
}

func ReadRecentChat(chatFile *multipart.FileHeader) (string, error) {
	file, err := chatFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rowCount := 50

	recentChat := strings.Join(lines[len(lines)-rowCount:], "\n")
	fmt.Println("recentChat: ")

	return recentChat, nil
}

// func GetClosenessScore(chatText string) (*ClosenessScore, error) {
// 	apiKey := os.Getenv("OPENAI_API_KEY")

// 	prompt := fmt.Sprintf(`
// 		以下の会話ログを基に、親密度を5つの観点で評価してください。
// 		各項目は20点満点で、合計100点とします。評価結果はJSON形式で出力してください。  JSON以外は出力しないでください．

// 		{
// 			"time": number,
// 			"balance": number,
// 			"rhythm": number,
// 			"type": number,
// 			"words": number,
// 		}

// 		### 評価項目:
// 		1. **時間帯 (time)**
// 		- 昼・夜などの時間帯による親密度の影響
// 		- 深夜の会話が多いほど親密度が高い傾向がある

// 		2. **バランス (balance)**
// 		- どちらか一方が話しすぎていないか
// 		- 双方が均等に会話しているか

// 		3. **テンポ (rhythm)**
// 		- メッセージのやりとりの間隔が短いか
// 		- リアルタイムな会話が多いほどスコアが高い

// 		4. **タイプ (type)**
// 		- 会話の種類（相談、冗談、情報共有など）
// 		- 相談や深い話が多いほど親密度が高い

// 		5. **言葉 (words)**
// 		- 使われる言葉の親密度（愛称、タメ口、感情表現）
// 		- 「ありがとう」「好き」「嬉しい」などのポジティブ表現が多いほどスコアが高い

// 		### 会話ログ:
// 		%s
// 	`, chatText)

// 	client := openai.NewClient(
// 		option.WithAPIKey(apiKey),
// 	)

// 	// var HistoricalComputerResponseSchema = GenerateSchema[HistoricalComputer]()

// 	// schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
// 	// 	Name:        openai.F("biography"),
// 	// 	Description: openai.F("Notable information about a person"),
// 	// 	Schema:      openai.F(HistoricalComputerResponseSchema),
// 	// 	Strict:      openai.Bool(true),
// 	// }

// 	resp, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
// 			openai.UserMessage(prompt),
// 		}),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 		// ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
// 		// 	openai.ResponseFormatJSONSchemaParam{
// 		// 		Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
// 		// 		JSONSchema: openai.F(schemaParam),
// 		// 	},
// 		// ),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(resp.Choices[0].Message.Content)
// 	// ```jsonを削除する
// 	validResult := strings.ReplaceAll(resp.Choices[0].Message.Content, "```json", "")
// 	validResult = strings.ReplaceAll(validResult, "```", "")
// 	// レスポンスをパース
// 	var score ClosenessScore
// 	err = json.Unmarshal([]byte(validResult), &score)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &score, nil
// }

func GetClosenessScore(chatText string) (*ClosenessScore, error) {
	prompt := fmt.Sprintf(`
		以下の会話ログを基に、親密度を5つの観点で評価してください。
		各項目は20点満点で、合計100点とします。評価結果はJSON形式で出力してください。  JSON以外は出力しないでください．


		### 評価項目:
		1. **時間帯 (time)**
		- 昼・夜などの時間帯による親密度の影響
		- 深夜の会話が多いほど親密度が高い傾向がある

		2. **バランス (balance)**
		- どちらか一方が話しすぎていないか
		- 双方が均等に会話しているか

		3. **テンポ (rhythm)**
		- メッセージのやりとりの間隔が短いか
		- リアルタイムな会話が多いほどスコアが高い

		4. **タイプ (type)**
		- 会話の種類（相談、冗談、情報共有など）
		- 相談や深い話が多いほど親密度が高い

		5. **言葉 (words)**
		- 使われる言葉の親密度（愛称、タメ口、感情表現）
		- 「ありがとう」「好き」「嬉しい」などのポジティブ表現が多いほどスコアが高い

		

		### 会話ログ:
		%s
	`, chatText)

	fmt.Println(prompt)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer client.Close()

	schema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"time": {
				Type:        genai.TypeInteger,
				Description: "昼・夜などの時間帯による親密度の影響",
			},
			"balance": {
				Type:        genai.TypeInteger,
				Description: "どちらか一方が話しすぎていないか",
			},
			"rhythm": {
				Type:        genai.TypeInteger,
				Description: "メッセージのやりとりの間隔が短いか",
			},
			"type": {
				Type:        genai.TypeInteger,
				Description: "会話の種類（相談、冗談、情報共有など）",
			},
			"words": {
				Type:        genai.TypeInteger,
				Description: "使われる言葉の親密度（愛称、タメ口、感情表現）",
			},
		},
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = schema
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var score ClosenessScore
	if err := json.Unmarshal([]byte(resp.Candidates[0].Content.Parts[0].(genai.Text)), &score); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	score.Total = score.Time + score.Balance + score.Rhythm + score.Type + score.Words
	fmt.Println(score.Balance)

	return &score, nil
}

//	func GetNextAction(total int) (string, error) {
//		switch {
//		case total >= 80:
//			return "親密度が非常に高い！次は、一緒に遊ぶ計画を立てたり、直接会う約束をしよう。", nil
//		case total >= 60:
//			return "かなり親密な関係！もっと深い話をしたり、通話を試みるのがおすすめ。", nil
//		case total >= 40:
//			return "普通の親しさ。会話の話題を増やし、相手に質問を多く投げかけよう。", nil
//		case total >= 20:
//			return "まだ距離があるかも。ポジティブな話題を増やして、相手の興味を引こう！", nil
//		default:
//			return "距離が遠い状態。無理に距離を詰めず、相手のペースに合わせてゆっくり関係を築こう。", nil
//		}
//	}
func GetNextAction(ctx context.Context, client *ent.Client, total int) (string, error) {
	nextaction, err := client.NextAction.
		Query().
		Where(nextaction.ScoreMinLTE(total), nextaction.ScoreMaxGTE(total)).
		Only(ctx)
	if err != nil {
		return "データベースから取り出せないよ", err
	}
	return nextaction.Action, nil
}

func (controller *Controller) GetBetween(c echo.Context) error {

	// LINEの会話データを読み込む
	chatFile, err := c.FormFile("text")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "text file is required"})
	}
	lineText, err := ReadRecentChat(chatFile) // 直近50行を取得
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to read chat file"})
	}

	// AIで親密度スコアを取得
	score, err := GetClosenessScore(lineText)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//スコアからnextactionを決定
	nextaction, err := GetNextAction(c.Request().Context(), controller.entClient, score.Total)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := struct {
		Score      *ClosenessScore `json:"score"`
		NextAction string          `json:"next_action"`
	}{
		Score:      score,
		NextAction: nextaction,
	}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 全オリジンを許可
		AllowHeaders: []string{"*"}, // 全ヘッダーを許可
		AllowMethods: []string{"*"}, // 全HTTPメソッドを許可
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
	e.POST("/between", controller.GetBetween)

	e.Logger.Fatal(e.Start(":8080"))
}
