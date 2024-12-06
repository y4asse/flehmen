package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// University holds the schema definition for the University entity.
type University struct {
	ent.Schema
}

// Fields of the University.
func (University) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("deviationLowerValue"),
		field.Int("deviationUpperValue"),
		field.String("abbreviation"),
		field.String("prefecture"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the University.
func (University) Edges() []ent.Edge {
	return nil
}

func (University) Indexes() []ent.Index {
	return []ent.Index{
		// 複合インデックスの作成
		index.Fields("name", "abbreviation").
			Unique(),
	}
}
