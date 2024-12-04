package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sukipi holds the schema definition for the Sukipi entity.
type Sukipi struct {
	ent.Schema
}

// Fields of the Sukipi.
func (Sukipi) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("weight"),
		field.Float("height"),
		field.String("x_id").Optional(),
		field.String("instagram_id").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_male"),
		field.Time("start_at").Optional(),
	}
}

// Edges of the Sukipi.
func (Sukipi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mbti", Mbti.Type).Unique(),
		edge.To("tweets", Tweet.Type),
	}
}
