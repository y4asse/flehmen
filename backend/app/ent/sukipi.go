// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flehmen-api/ent/sukipi"
	"flehmen-api/ent/user"
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
	// LikedAt holds the value of the "liked_at" field.
	LikedAt time.Time `json:"liked_at,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight *float64 `json:"weight,omitempty"`
	// Height holds the value of the "height" field.
	Height *float64 `json:"height,omitempty"`
	// XID holds the value of the "x_id" field.
	XID *string `json:"x_id,omitempty"`
	// Hobby holds the value of the "hobby" field.
	Hobby *string `json:"hobby,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday *time.Time `json:"birthday,omitempty"`
	// ShoesSize holds the value of the "shoesSize" field.
	ShoesSize *float64 `json:"shoesSize,omitempty"`
	// Family holds the value of the "family" field.
	Family *string `json:"family,omitempty"`
	// NearlyStation holds the value of the "nearly_station" field.
	NearlyStation *string `json:"nearly_station,omitempty"`
	// Mbti holds the value of the "mbti" field.
	Mbti *string `json:"mbti,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SukipiQuery when eager-loading is set.
	Edges        SukipiEdges `json:"edges"`
	sukipi_user  *int
	selectValues sql.SelectValues
}

// SukipiEdges holds the relations/edges for other nodes in the graph.
type SukipiEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SukipiEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sukipi) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sukipi.FieldWeight, sukipi.FieldHeight, sukipi.FieldShoesSize:
			values[i] = new(sql.NullFloat64)
		case sukipi.FieldID:
			values[i] = new(sql.NullInt64)
		case sukipi.FieldName, sukipi.FieldXID, sukipi.FieldHobby, sukipi.FieldFamily, sukipi.FieldNearlyStation, sukipi.FieldMbti:
			values[i] = new(sql.NullString)
		case sukipi.FieldLikedAt, sukipi.FieldBirthday, sukipi.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case sukipi.ForeignKeys[0]: // sukipi_user
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
		case sukipi.FieldLikedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field liked_at", values[i])
			} else if value.Valid {
				s.LikedAt = value.Time
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
		case sukipi.FieldHobby:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hobby", values[i])
			} else if value.Valid {
				s.Hobby = new(string)
				*s.Hobby = value.String
			}
		case sukipi.FieldBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				s.Birthday = new(time.Time)
				*s.Birthday = value.Time
			}
		case sukipi.FieldShoesSize:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field shoesSize", values[i])
			} else if value.Valid {
				s.ShoesSize = new(float64)
				*s.ShoesSize = value.Float64
			}
		case sukipi.FieldFamily:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family", values[i])
			} else if value.Valid {
				s.Family = new(string)
				*s.Family = value.String
			}
		case sukipi.FieldNearlyStation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nearly_station", values[i])
			} else if value.Valid {
				s.NearlyStation = new(string)
				*s.NearlyStation = value.String
			}
		case sukipi.FieldMbti:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mbti", values[i])
			} else if value.Valid {
				s.Mbti = new(string)
				*s.Mbti = value.String
			}
		case sukipi.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sukipi.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sukipi_user", value)
			} else if value.Valid {
				s.sukipi_user = new(int)
				*s.sukipi_user = int(value.Int64)
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

// QueryUser queries the "user" edge of the Sukipi entity.
func (s *Sukipi) QueryUser() *UserQuery {
	return NewSukipiClient(s.config).QueryUser(s)
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
	builder.WriteString("liked_at=")
	builder.WriteString(s.LikedAt.Format(time.ANSIC))
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
	if v := s.Hobby; v != nil {
		builder.WriteString("hobby=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.Birthday; v != nil {
		builder.WriteString("birthday=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.ShoesSize; v != nil {
		builder.WriteString("shoesSize=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := s.Family; v != nil {
		builder.WriteString("family=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.NearlyStation; v != nil {
		builder.WriteString("nearly_station=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.Mbti; v != nil {
		builder.WriteString("mbti=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sukipis is a parsable slice of Sukipi.
type Sukipis []*Sukipi
