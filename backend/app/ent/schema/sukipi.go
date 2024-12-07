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
		field.Float("weight").Nillable().Optional(),
		field.Float("height").Nillable().Optional(),
		field.String("x_id").Nillable().Optional(),
		field.String("instagram_id").Nillable().Optional(),
		field.Bool("is_male"),
		field.Time("start_at"),
		field.Time("birthday").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Sukipi.
func (Sukipi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mbti", Mbti.Type).Unique(),
		edge.To("tweets", Tweet.Type),
	}
}
