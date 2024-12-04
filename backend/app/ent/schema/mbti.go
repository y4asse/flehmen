package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Mbti holds the schema definition for the Mbti entity.
type Mbti struct {
	ent.Schema
}

// Fields of the Mbti.
func (Mbti) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.Time("created_at").Default(time.Now),
	}

}

// Edges of the Mbti.
func (Mbti) Edges() []ent.Edge {
	return nil
}
