// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldClerkID holds the string denoting the clerk_id field in the database.
	FieldClerkID = "clerk_id"
	// FieldIsMale holds the string denoting the is_male field in the database.
	FieldIsMale = "is_male"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeMbti holds the string denoting the mbti edge name in mutations.
	EdgeMbti = "mbti"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MbtiTable is the table that holds the mbti relation/edge.
	MbtiTable = "users"
	// MbtiInverseTable is the table name for the Mbti entity.
	// It exists in this package in order to avoid circular dependency with the "mbti" package.
	MbtiInverseTable = "mbtis"
	// MbtiColumn is the table column denoting the mbti relation/edge.
	MbtiColumn = "user_mbti"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWeight,
	FieldHeight,
	FieldClerkID,
	FieldIsMale,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_mbti",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByClerkID orders the results by the clerk_id field.
func ByClerkID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClerkID, opts...).ToFunc()
}

// ByIsMale orders the results by the is_male field.
func ByIsMale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsMale, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByMbtiField orders the results by mbti field.
func ByMbtiField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMbtiStep(), sql.OrderByField(field, opts...))
	}
}
func newMbtiStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MbtiInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MbtiTable, MbtiColumn),
	)
}
