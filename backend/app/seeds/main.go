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

	if err := seedTweets(ctx, client); err != nil {
		log.Fatalf("failed seeding tweets: %v", err)
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
