// Code generated by ent, DO NOT EDIT.

package twitteruser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the twitteruser type in the database.
	Label = "twitter_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeReplies holds the string denoting the replies edge name in mutations.
	EdgeReplies = "replies"
	// Table holds the table name of the twitteruser in the database.
	Table = "twitter_users"
	// RepliesTable is the table that holds the replies relation/edge.
	RepliesTable = "tweets"
	// RepliesInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	RepliesInverseTable = "tweets"
	// RepliesColumn is the table column denoting the replies relation/edge.
	RepliesColumn = "reply_twitter_user_id"
)

// Columns holds all SQL columns for twitteruser fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUsername,
	FieldCreatedAt,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the TwitterUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByRepliesCount orders the results by replies count.
func ByRepliesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRepliesStep(), opts...)
	}
}

// ByReplies orders the results by replies terms.
func ByReplies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRepliesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRepliesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RepliesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
	)
}
