package main

import (
	"context"
	"encoding/json"
	"flehmen-api/ent"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
)

type University struct {
	Name                string `json:"name"`
	Abbreviation        string `json:"abbreviation"`
	DeviationLowerValue int    `json:"deviationLowerValue"`
	DeviationUpperValue int    `json:"deviationUpperValue"`
	Prefecture          string `json:"prefecture"`
}

type NextAction struct {
	MinScore   int    `json:"score_min"`
	MaxScore   int    `json:"score_max"`
	NextAction string `json:"next_action"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Tweet struct {
	ID              string `json:"id"`
	InReplyToUserID string `json:"in_reply_to_user_id"`
	Text            string `json:"text"`
	CreatedAt       string `json:"created_at"`
}

type Includes struct {
	Users []User `json:"users"`
}

type Metadeta struct {
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	NextToken   string `json:"next_token"`
}

type TweetData struct {
	Data     []Tweet  `json:"data"`
	Includes Includes `json:"includes"`
	Metadeta Metadeta `json:"metadeta"`
}

type UserData struct {
	ID      string `json:"id"`
	ClerkID string `json:"clerk_id"`
}

func main() {
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

	ctx := context.Background()

	if err := seedUniversity(ctx, client); err != nil {
		log.Fatalf("failed seeding university: %v", err)
	}
	if err := seedNextAction(ctx, client); err != nil {
		log.Fatalf("failed seeding next action: %v", err)
	}

	if err := seedTweets(ctx, client); err != nil {
		log.Fatalf("failed seeding tweets: %v", err)
	}

	if err := seedUser(ctx, client); err != nil {
		log.Fatalf("failed seeding user: %v", err)
	}

	if err := seedSukipi(ctx, client); err != nil {
		log.Fatalf("failed seeding sukipi: %v", err)
	}

	fmt.Println("シーディングが完了しました")
}

func seedUniversity(ctx context.Context, client *ent.Client) error {
	file, err := os.Open("seeds/universities.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var universities []University
	if err := json.NewDecoder(file).Decode(&universities); err != nil {
		return err
	}
	var bulk []*ent.UniversityCreate
	for _, u := range universities {
		bulk = append(bulk, client.University.Create().
			SetName(u.Name).
			SetAbbreviation(u.Abbreviation).
			SetDeviationLowerValue(u.DeviationLowerValue).
			SetDeviationUpperValue(u.DeviationUpperValue).
			SetPrefecture(u.Prefecture),
		)
	}
	if err := client.University.CreateBulk(bulk...).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func seedNextAction(ctx context.Context, client *ent.Client) error {
	file, err := os.Open("seeds/nextaction.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var nextaction []NextAction
	if err := json.NewDecoder(file).Decode(&nextaction); err != nil {
		return err
	}
	var bulk []*ent.NextActionCreate
	for _, n := range nextaction {
		bulk = append(bulk, client.NextAction.Create().
			SetScoreMin(n.MinScore).
			SetScoreMax(n.MaxScore).
			SetAction(n.NextAction),
		)
	}
	if err := client.NextAction.CreateBulk(bulk...).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func seedTweets(ctx context.Context, client *ent.Client) error {
	file, err := os.Open("seeds/tweets.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var tweetData TweetData
	if err := json.NewDecoder(file).Decode(&tweetData); err != nil {
		return err
	}

	var bulkUsers []*ent.TwitterUserCreate
	for _, u := range tweetData.Includes.Users {
		userId, err := strconv.Atoi(u.ID)
		if err != nil {
			return err
		}
		bulkUsers = append(bulkUsers, client.TwitterUser.Create().
			SetID(userId).
			SetName(u.Name).
			SetUsername(u.Username),
		)
	}
	if err := client.TwitterUser.CreateBulk(bulkUsers...).Exec(ctx); err != nil {
		return err
	}

	var bulk []*ent.TweetCreate
	for _, t := range tweetData.Data {
		tweetID, err := strconv.Atoi(t.ID)
		if err != nil {
			return err
		}
		var replyToUserID *int
		if t.InReplyToUserID != "" {
			id, err := strconv.Atoi(t.InReplyToUserID)
			if err == nil {
				replyToUserID = &id
			}
		}
		createdAt, err := time.Parse(time.RFC3339, t.CreatedAt)
		if err != nil {
			return err
		}

		bulk = append(bulk, client.Tweet.Create().
			SetText(t.Text).
			SetTweetID(tweetID).
			SetTweetCreatedAt(createdAt).
			SetNillableReplyTwitterUserID(replyToUserID),
		)
	}
	if err := client.Tweet.CreateBulk(bulk...).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func seedUser(ctx context.Context, client *ent.Client) error {
	err := client.User.Create().
		SetClerkID("user_2sfTghKeM7uB97iqooGRkN6fAd8").
		SetName("test").
		SetIsMale(true).
		SetWeight(173).
		SetHeight(173).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func seedSukipi(ctx context.Context, client *ent.Client) error {
	err := client.Sukipi.Create().
		SetUserID(1).
		SetName("やせ").
		SetLikedAt(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)).
		SetWeight(66).
		SetHeight(174).
		SetXID("y4isse").
		SetHobby("Apex").
		SetBirthday(time.Date(2003, 2, 18, 0, 0, 0, 0, time.UTC)).
		SetShoesSize(173).
		SetFamily("片親").
		SetNearlyStation("日比野").
		SetMbti("INTP").
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
