package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SpecialEvent holds the schema definition for the SpecialEvent entity.
type SpecialEvent struct {
	ent.Schema
}

// Fields of the SpecialEvent.
func (SpecialEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Time("occured_at"),
		field.String("title"),
		field.String("detail_comment"),
	}
}

// Edges of the SpecialEvent.
func (SpecialEvent) Edges() []ent.Edge {
	return nil
}
