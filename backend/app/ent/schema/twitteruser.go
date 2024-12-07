package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TwitterUser holds the schema definition for the TwitterUser entity.
type TwitterUser struct {
	ent.Schema
}

// Fields of the TwitterUser.
func (TwitterUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.String("username"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the TwitterUser.
func (TwitterUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("replies", Tweet.Type),
	}
}
