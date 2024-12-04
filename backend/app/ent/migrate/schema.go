// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MbtisColumns holds the columns for the "mbtis" table.
	MbtisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// MbtisTable holds the schema information for the "mbtis" table.
	MbtisTable = &schema.Table{
		Name:       "mbtis",
		Columns:    MbtisColumns,
		PrimaryKey: []*schema.Column{MbtisColumns[0]},
	}
	// SpecialEventsColumns holds the columns for the "special_events" table.
	SpecialEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "occured_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "detail_comment", Type: field.TypeString},
		{Name: "user_special_events", Type: field.TypeInt, Nullable: true},
	}
	// SpecialEventsTable holds the schema information for the "special_events" table.
	SpecialEventsTable = &schema.Table{
		Name:       "special_events",
		Columns:    SpecialEventsColumns,
		PrimaryKey: []*schema.Column{SpecialEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "special_events_users_special_events",
				Columns:    []*schema.Column{SpecialEventsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SukipisColumns holds the columns for the "sukipis" table.
	SukipisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "weight", Type: field.TypeFloat64},
		{Name: "height", Type: field.TypeFloat64},
		{Name: "x_id", Type: field.TypeString},
		{Name: "instagram_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_male", Type: field.TypeBool},
		{Name: "start_at", Type: field.TypeTime, Nullable: true},
		{Name: "sukipi_mbti", Type: field.TypeInt, Nullable: true},
	}
	// SukipisTable holds the schema information for the "sukipis" table.
	SukipisTable = &schema.Table{
		Name:       "sukipis",
		Columns:    SukipisColumns,
		PrimaryKey: []*schema.Column{SukipisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sukipis_mbtis_mbti",
				Columns:    []*schema.Column{SukipisColumns[9]},
				RefColumns: []*schema.Column{MbtisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TweetsColumns holds the columns for the "tweets" table.
	TweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
		{Name: "tweet_id", Type: field.TypeInt},
		{Name: "tweet_created_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "sukipi_tweets", Type: field.TypeInt, Nullable: true},
	}
	// TweetsTable holds the schema information for the "tweets" table.
	TweetsTable = &schema.Table{
		Name:       "tweets",
		Columns:    TweetsColumns,
		PrimaryKey: []*schema.Column{TweetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweets_sukipis_tweets",
				Columns:    []*schema.Column{TweetsColumns[5]},
				RefColumns: []*schema.Column{SukipisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "weight", Type: field.TypeFloat64, Nullable: true},
		{Name: "height", Type: field.TypeFloat64, Nullable: true},
		{Name: "clerk_id", Type: field.TypeString},
		{Name: "is_male", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_mbti", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_mbtis_mbti",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{MbtisColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MbtisTable,
		SpecialEventsTable,
		SukipisTable,
		TweetsTable,
		UsersTable,
	}
)

func init() {
	SpecialEventsTable.ForeignKeys[0].RefTable = UsersTable
	SukipisTable.ForeignKeys[0].RefTable = MbtisTable
	TweetsTable.ForeignKeys[0].RefTable = SukipisTable
	UsersTable.ForeignKeys[0].RefTable = MbtisTable
}