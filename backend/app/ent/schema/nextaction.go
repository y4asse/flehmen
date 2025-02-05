package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NextAction holds the schema definition for the NextAction entity.
type NextAction struct {
	ent.Schema
}

func (NextAction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("score_min"),
		field.Int("score_max"),
		field.String("action"),
	}
}
