package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Int("tweet_id").Unique(),
		field.Time("tweet_created_at"),
		field.Int("reply_twitter_user_id").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Int("author_id"),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", TwitterUser.Type).
			Ref("replies").
			Unique().
			Field("reply_twitter_user_id"),
	}
}
