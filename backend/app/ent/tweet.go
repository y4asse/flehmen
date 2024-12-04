// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flehmen-api/ent/tweet"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Tweet is the model entity for the Tweet schema.
type Tweet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// TweetID holds the value of the "tweet_id" field.
	TweetID int `json:"tweet_id,omitempty"`
	// TweetCreatedAt holds the value of the "tweet_created_at" field.
	TweetCreatedAt time.Time `json:"tweet_created_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt     time.Time `json:"created_at,omitempty"`
	sukipi_tweets *int
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tweet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID, tweet.FieldTweetID:
			values[i] = new(sql.NullInt64)
		case tweet.FieldText:
			values[i] = new(sql.NullString)
		case tweet.FieldTweetCreatedAt, tweet.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case tweet.ForeignKeys[0]: // sukipi_tweets
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tweet fields.
func (t *Tweet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tweet.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case tweet.FieldTweetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				t.TweetID = int(value.Int64)
			}
		case tweet.FieldTweetCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_created_at", values[i])
			} else if value.Valid {
				t.TweetCreatedAt = value.Time
			}
		case tweet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tweet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sukipi_tweets", value)
			} else if value.Valid {
				t.sukipi_tweets = new(int)
				*t.sukipi_tweets = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tweet.
// This includes values selected through modifiers, order, etc.
func (t *Tweet) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Tweet.
// Note that you need to call Tweet.Unwrap() before calling this method if this Tweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tweet) Update() *TweetUpdateOne {
	return NewTweetClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tweet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tweet) Unwrap() *Tweet {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tweet is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tweet) String() string {
	var builder strings.Builder
	builder.WriteString("Tweet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteString(", ")
	builder.WriteString("tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TweetID))
	builder.WriteString(", ")
	builder.WriteString("tweet_created_at=")
	builder.WriteString(t.TweetCreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tweets is a parsable slice of Tweet.
type Tweets []*Tweet
