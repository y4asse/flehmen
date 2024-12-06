package main

import (
	"context"
	"encoding/json"
	"flehmen-api/ent"
	"fmt"
	"log"
	"os"

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

	err = seedUniversity(ctx, client)
	if err != nil {
		log.Fatalf("failed seeding university: %v", err)
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
