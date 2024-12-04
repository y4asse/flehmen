// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flehmen-api/ent/mbti"
	"flehmen-api/ent/sukipi"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Sukipi is the model entity for the Sukipi schema.
type Sukipi struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight *float64 `json:"weight,omitempty"`
	// Height holds the value of the "height" field.
	Height *float64 `json:"height,omitempty"`
	// XID holds the value of the "x_id" field.
	XID *string `json:"x_id,omitempty"`
	// InstagramID holds the value of the "instagram_id" field.
	InstagramID *string `json:"instagram_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// IsMale holds the value of the "is_male" field.
	IsMale bool `json:"is_male,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SukipiQuery when eager-loading is set.
	Edges        SukipiEdges `json:"edges"`
	sukipi_mbti  *int
	selectValues sql.SelectValues
}

// SukipiEdges holds the relations/edges for other nodes in the graph.
type SukipiEdges struct {
	// Mbti holds the value of the mbti edge.
	Mbti *Mbti `json:"mbti,omitempty"`
	// Tweets holds the value of the tweets edge.
	Tweets []*Tweet `json:"tweets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MbtiOrErr returns the Mbti value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SukipiEdges) MbtiOrErr() (*Mbti, error) {
	if e.Mbti != nil {
		return e.Mbti, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: mbti.Label}
	}
	return nil, &NotLoadedError{edge: "mbti"}
}

// TweetsOrErr returns the Tweets value or an error if the edge
// was not loaded in eager-loading.
func (e SukipiEdges) TweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[1] {
		return e.Tweets, nil
	}
	return nil, &NotLoadedError{edge: "tweets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sukipi) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sukipi.FieldIsMale:
			values[i] = new(sql.NullBool)
		case sukipi.FieldWeight, sukipi.FieldHeight:
			values[i] = new(sql.NullFloat64)
		case sukipi.FieldID:
			values[i] = new(sql.NullInt64)
		case sukipi.FieldName, sukipi.FieldXID, sukipi.FieldInstagramID:
			values[i] = new(sql.NullString)
		case sukipi.FieldCreatedAt, sukipi.FieldStartAt:
			values[i] = new(sql.NullTime)
		case sukipi.ForeignKeys[0]: // sukipi_mbti
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sukipi fields.
func (s *Sukipi) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sukipi.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sukipi.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case sukipi.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				s.Weight = new(float64)
				*s.Weight = value.Float64
			}
		case sukipi.FieldHeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				s.Height = new(float64)
				*s.Height = value.Float64
			}
		case sukipi.FieldXID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field x_id", values[i])
			} else if value.Valid {
				s.XID = new(string)
				*s.XID = value.String
			}
		case sukipi.FieldInstagramID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instagram_id", values[i])
			} else if value.Valid {
				s.InstagramID = new(string)
				*s.InstagramID = value.String
			}
		case sukipi.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sukipi.FieldIsMale:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_male", values[i])
			} else if value.Valid {
				s.IsMale = value.Bool
			}
		case sukipi.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				s.StartAt = value.Time
			}
		case sukipi.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sukipi_mbti", value)
			} else if value.Valid {
				s.sukipi_mbti = new(int)
				*s.sukipi_mbti = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Sukipi.
// This includes values selected through modifiers, order, etc.
func (s *Sukipi) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryMbti queries the "mbti" edge of the Sukipi entity.
func (s *Sukipi) QueryMbti() *MbtiQuery {
	return NewSukipiClient(s.config).QueryMbti(s)
}

// QueryTweets queries the "tweets" edge of the Sukipi entity.
func (s *Sukipi) QueryTweets() *TweetQuery {
	return NewSukipiClient(s.config).QueryTweets(s)
}

// Update returns a builder for updating this Sukipi.
// Note that you need to call Sukipi.Unwrap() before calling this method if this Sukipi
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sukipi) Update() *SukipiUpdateOne {
	return NewSukipiClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Sukipi entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sukipi) Unwrap() *Sukipi {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sukipi is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sukipi) String() string {
	var builder strings.Builder
	builder.WriteString("Sukipi(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	if v := s.Weight; v != nil {
		builder.WriteString("weight=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := s.Height; v != nil {
		builder.WriteString("height=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := s.XID; v != nil {
		builder.WriteString("x_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.InstagramID; v != nil {
		builder.WriteString("instagram_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_male=")
	builder.WriteString(fmt.Sprintf("%v", s.IsMale))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(s.StartAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sukipis is a parsable slice of Sukipi.
type Sukipis []*Sukipi