// Code generated by ent, DO NOT EDIT.

package nextaction

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the nextaction type in the database.
	Label = "next_action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScoreMin holds the string denoting the score_min field in the database.
	FieldScoreMin = "score_min"
	// FieldScoreMax holds the string denoting the score_max field in the database.
	FieldScoreMax = "score_max"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// Table holds the table name of the nextaction in the database.
	Table = "next_actions"
)

// Columns holds all SQL columns for nextaction fields.
var Columns = []string{
	FieldID,
	FieldScoreMin,
	FieldScoreMax,
	FieldAction,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the NextAction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScoreMin orders the results by the score_min field.
func ByScoreMin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScoreMin, opts...).ToFunc()
}

// ByScoreMax orders the results by the score_max field.
func ByScoreMax(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScoreMax, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}
