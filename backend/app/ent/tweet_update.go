// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flehmen-api/ent/predicate"
	"flehmen-api/ent/tweet"
	"flehmen-api/ent/twitteruser"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetText sets the "text" field.
func (tu *TweetUpdate) SetText(s string) *TweetUpdate {
	tu.mutation.SetText(s)
	return tu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableText(s *string) *TweetUpdate {
	if s != nil {
		tu.SetText(*s)
	}
	return tu
}

// SetTweetID sets the "tweet_id" field.
func (tu *TweetUpdate) SetTweetID(i int) *TweetUpdate {
	tu.mutation.ResetTweetID()
	tu.mutation.SetTweetID(i)
	return tu
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableTweetID(i *int) *TweetUpdate {
	if i != nil {
		tu.SetTweetID(*i)
	}
	return tu
}

// AddTweetID adds i to the "tweet_id" field.
func (tu *TweetUpdate) AddTweetID(i int) *TweetUpdate {
	tu.mutation.AddTweetID(i)
	return tu
}

// SetTweetCreatedAt sets the "tweet_created_at" field.
func (tu *TweetUpdate) SetTweetCreatedAt(t time.Time) *TweetUpdate {
	tu.mutation.SetTweetCreatedAt(t)
	return tu
}

// SetNillableTweetCreatedAt sets the "tweet_created_at" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableTweetCreatedAt(t *time.Time) *TweetUpdate {
	if t != nil {
		tu.SetTweetCreatedAt(*t)
	}
	return tu
}

// SetReplyTwitterUserID sets the "reply_twitter_user_id" field.
func (tu *TweetUpdate) SetReplyTwitterUserID(i int) *TweetUpdate {
	tu.mutation.SetReplyTwitterUserID(i)
	return tu
}

// SetNillableReplyTwitterUserID sets the "reply_twitter_user_id" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableReplyTwitterUserID(i *int) *TweetUpdate {
	if i != nil {
		tu.SetReplyTwitterUserID(*i)
	}
	return tu
}

// ClearReplyTwitterUserID clears the value of the "reply_twitter_user_id" field.
func (tu *TweetUpdate) ClearReplyTwitterUserID() *TweetUpdate {
	tu.mutation.ClearReplyTwitterUserID()
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TweetUpdate) SetCreatedAt(t time.Time) *TweetUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableCreatedAt(t *time.Time) *TweetUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUserID sets the "user" edge to the TwitterUser entity by ID.
func (tu *TweetUpdate) SetUserID(id int) *TweetUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetNillableUserID sets the "user" edge to the TwitterUser entity by ID if the given value is not nil.
func (tu *TweetUpdate) SetNillableUserID(id *int) *TweetUpdate {
	if id != nil {
		tu = tu.SetUserID(*id)
	}
	return tu
}

// SetUser sets the "user" edge to the TwitterUser entity.
func (tu *TweetUpdate) SetUser(t *TwitterUser) *TweetUpdate {
	return tu.SetUserID(t.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearUser clears the "user" edge to the TwitterUser entity.
func (tu *TweetUpdate) ClearUser() *TweetUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if value, ok := tu.mutation.TweetID(); ok {
		_spec.SetField(tweet.FieldTweetID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedTweetID(); ok {
		_spec.AddField(tweet.FieldTweetID, field.TypeInt, value)
	}
	if value, ok := tu.mutation.TweetCreatedAt(); ok {
		_spec.SetField(tweet.FieldTweetCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(tweet.FieldCreatedAt, field.TypeTime, value)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: []string{tweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(twitteruser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: []string{tweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(twitteruser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetText sets the "text" field.
func (tuo *TweetUpdateOne) SetText(s string) *TweetUpdateOne {
	tuo.mutation.SetText(s)
	return tuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableText(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetText(*s)
	}
	return tuo
}

// SetTweetID sets the "tweet_id" field.
func (tuo *TweetUpdateOne) SetTweetID(i int) *TweetUpdateOne {
	tuo.mutation.ResetTweetID()
	tuo.mutation.SetTweetID(i)
	return tuo
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableTweetID(i *int) *TweetUpdateOne {
	if i != nil {
		tuo.SetTweetID(*i)
	}
	return tuo
}

// AddTweetID adds i to the "tweet_id" field.
func (tuo *TweetUpdateOne) AddTweetID(i int) *TweetUpdateOne {
	tuo.mutation.AddTweetID(i)
	return tuo
}

// SetTweetCreatedAt sets the "tweet_created_at" field.
func (tuo *TweetUpdateOne) SetTweetCreatedAt(t time.Time) *TweetUpdateOne {
	tuo.mutation.SetTweetCreatedAt(t)
	return tuo
}

// SetNillableTweetCreatedAt sets the "tweet_created_at" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableTweetCreatedAt(t *time.Time) *TweetUpdateOne {
	if t != nil {
		tuo.SetTweetCreatedAt(*t)
	}
	return tuo
}

// SetReplyTwitterUserID sets the "reply_twitter_user_id" field.
func (tuo *TweetUpdateOne) SetReplyTwitterUserID(i int) *TweetUpdateOne {
	tuo.mutation.SetReplyTwitterUserID(i)
	return tuo
}

// SetNillableReplyTwitterUserID sets the "reply_twitter_user_id" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableReplyTwitterUserID(i *int) *TweetUpdateOne {
	if i != nil {
		tuo.SetReplyTwitterUserID(*i)
	}
	return tuo
}

// ClearReplyTwitterUserID clears the value of the "reply_twitter_user_id" field.
func (tuo *TweetUpdateOne) ClearReplyTwitterUserID() *TweetUpdateOne {
	tuo.mutation.ClearReplyTwitterUserID()
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TweetUpdateOne) SetCreatedAt(t time.Time) *TweetUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableCreatedAt(t *time.Time) *TweetUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUserID sets the "user" edge to the TwitterUser entity by ID.
func (tuo *TweetUpdateOne) SetUserID(id int) *TweetUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetNillableUserID sets the "user" edge to the TwitterUser entity by ID if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableUserID(id *int) *TweetUpdateOne {
	if id != nil {
		tuo = tuo.SetUserID(*id)
	}
	return tuo
}

// SetUser sets the "user" edge to the TwitterUser entity.
func (tuo *TweetUpdateOne) SetUser(t *TwitterUser) *TweetUpdateOne {
	return tuo.SetUserID(t.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearUser clears the "user" edge to the TwitterUser entity.
func (tuo *TweetUpdateOne) ClearUser() *TweetUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Where appends a list predicates to the TweetUpdate builder.
func (tuo *TweetUpdateOne) Where(ps ...predicate.Tweet) *TweetUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if value, ok := tuo.mutation.TweetID(); ok {
		_spec.SetField(tweet.FieldTweetID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedTweetID(); ok {
		_spec.AddField(tweet.FieldTweetID, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.TweetCreatedAt(); ok {
		_spec.SetField(tweet.FieldTweetCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(tweet.FieldCreatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: []string{tweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(twitteruser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: []string{tweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(twitteruser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
