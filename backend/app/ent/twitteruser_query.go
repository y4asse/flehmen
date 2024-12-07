// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"flehmen-api/ent/predicate"
	"flehmen-api/ent/tweet"
	"flehmen-api/ent/twitteruser"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TwitterUserQuery is the builder for querying TwitterUser entities.
type TwitterUserQuery struct {
	config
	ctx         *QueryContext
	order       []twitteruser.OrderOption
	inters      []Interceptor
	predicates  []predicate.TwitterUser
	withReplies *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwitterUserQuery builder.
func (tuq *TwitterUserQuery) Where(ps ...predicate.TwitterUser) *TwitterUserQuery {
	tuq.predicates = append(tuq.predicates, ps...)
	return tuq
}

// Limit the number of records to be returned by this query.
func (tuq *TwitterUserQuery) Limit(limit int) *TwitterUserQuery {
	tuq.ctx.Limit = &limit
	return tuq
}

// Offset to start from.
func (tuq *TwitterUserQuery) Offset(offset int) *TwitterUserQuery {
	tuq.ctx.Offset = &offset
	return tuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tuq *TwitterUserQuery) Unique(unique bool) *TwitterUserQuery {
	tuq.ctx.Unique = &unique
	return tuq
}

// Order specifies how the records should be ordered.
func (tuq *TwitterUserQuery) Order(o ...twitteruser.OrderOption) *TwitterUserQuery {
	tuq.order = append(tuq.order, o...)
	return tuq
}

// QueryReplies chains the current query on the "replies" edge.
func (tuq *TwitterUserQuery) QueryReplies() *TweetQuery {
	query := (&TweetClient{config: tuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(twitteruser.Table, twitteruser.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, twitteruser.RepliesTable, twitteruser.RepliesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TwitterUser entity from the query.
// Returns a *NotFoundError when no TwitterUser was found.
func (tuq *TwitterUserQuery) First(ctx context.Context) (*TwitterUser, error) {
	nodes, err := tuq.Limit(1).All(setContextOp(ctx, tuq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twitteruser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tuq *TwitterUserQuery) FirstX(ctx context.Context) *TwitterUser {
	node, err := tuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwitterUser ID from the query.
// Returns a *NotFoundError when no TwitterUser ID was found.
func (tuq *TwitterUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tuq.Limit(1).IDs(setContextOp(ctx, tuq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twitteruser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tuq *TwitterUserQuery) FirstIDX(ctx context.Context) int {
	id, err := tuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwitterUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwitterUser entity is found.
// Returns a *NotFoundError when no TwitterUser entities are found.
func (tuq *TwitterUserQuery) Only(ctx context.Context) (*TwitterUser, error) {
	nodes, err := tuq.Limit(2).All(setContextOp(ctx, tuq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twitteruser.Label}
	default:
		return nil, &NotSingularError{twitteruser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tuq *TwitterUserQuery) OnlyX(ctx context.Context) *TwitterUser {
	node, err := tuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwitterUser ID in the query.
// Returns a *NotSingularError when more than one TwitterUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (tuq *TwitterUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tuq.Limit(2).IDs(setContextOp(ctx, tuq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twitteruser.Label}
	default:
		err = &NotSingularError{twitteruser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tuq *TwitterUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := tuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwitterUsers.
func (tuq *TwitterUserQuery) All(ctx context.Context) ([]*TwitterUser, error) {
	ctx = setContextOp(ctx, tuq.ctx, ent.OpQueryAll)
	if err := tuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwitterUser, *TwitterUserQuery]()
	return withInterceptors[[]*TwitterUser](ctx, tuq, qr, tuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tuq *TwitterUserQuery) AllX(ctx context.Context) []*TwitterUser {
	nodes, err := tuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwitterUser IDs.
func (tuq *TwitterUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tuq.ctx.Unique == nil && tuq.path != nil {
		tuq.Unique(true)
	}
	ctx = setContextOp(ctx, tuq.ctx, ent.OpQueryIDs)
	if err = tuq.Select(twitteruser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tuq *TwitterUserQuery) IDsX(ctx context.Context) []int {
	ids, err := tuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tuq *TwitterUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tuq.ctx, ent.OpQueryCount)
	if err := tuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tuq, querierCount[*TwitterUserQuery](), tuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tuq *TwitterUserQuery) CountX(ctx context.Context) int {
	count, err := tuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tuq *TwitterUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tuq.ctx, ent.OpQueryExist)
	switch _, err := tuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tuq *TwitterUserQuery) ExistX(ctx context.Context) bool {
	exist, err := tuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwitterUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tuq *TwitterUserQuery) Clone() *TwitterUserQuery {
	if tuq == nil {
		return nil
	}
	return &TwitterUserQuery{
		config:      tuq.config,
		ctx:         tuq.ctx.Clone(),
		order:       append([]twitteruser.OrderOption{}, tuq.order...),
		inters:      append([]Interceptor{}, tuq.inters...),
		predicates:  append([]predicate.TwitterUser{}, tuq.predicates...),
		withReplies: tuq.withReplies.Clone(),
		// clone intermediate query.
		sql:  tuq.sql.Clone(),
		path: tuq.path,
	}
}

// WithReplies tells the query-builder to eager-load the nodes that are connected to
// the "replies" edge. The optional arguments are used to configure the query builder of the edge.
func (tuq *TwitterUserQuery) WithReplies(opts ...func(*TweetQuery)) *TwitterUserQuery {
	query := (&TweetClient{config: tuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tuq.withReplies = query
	return tuq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TwitterUser.Query().
//		GroupBy(twitteruser.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tuq *TwitterUserQuery) GroupBy(field string, fields ...string) *TwitterUserGroupBy {
	tuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwitterUserGroupBy{build: tuq}
	grbuild.flds = &tuq.ctx.Fields
	grbuild.label = twitteruser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TwitterUser.Query().
//		Select(twitteruser.FieldName).
//		Scan(ctx, &v)
func (tuq *TwitterUserQuery) Select(fields ...string) *TwitterUserSelect {
	tuq.ctx.Fields = append(tuq.ctx.Fields, fields...)
	sbuild := &TwitterUserSelect{TwitterUserQuery: tuq}
	sbuild.label = twitteruser.Label
	sbuild.flds, sbuild.scan = &tuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwitterUserSelect configured with the given aggregations.
func (tuq *TwitterUserQuery) Aggregate(fns ...AggregateFunc) *TwitterUserSelect {
	return tuq.Select().Aggregate(fns...)
}

func (tuq *TwitterUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tuq); err != nil {
				return err
			}
		}
	}
	for _, f := range tuq.ctx.Fields {
		if !twitteruser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tuq.path != nil {
		prev, err := tuq.path(ctx)
		if err != nil {
			return err
		}
		tuq.sql = prev
	}
	return nil
}

func (tuq *TwitterUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwitterUser, error) {
	var (
		nodes       = []*TwitterUser{}
		_spec       = tuq.querySpec()
		loadedTypes = [1]bool{
			tuq.withReplies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwitterUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwitterUser{config: tuq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tuq.withReplies; query != nil {
		if err := tuq.loadReplies(ctx, query, nodes,
			func(n *TwitterUser) { n.Edges.Replies = []*Tweet{} },
			func(n *TwitterUser, e *Tweet) { n.Edges.Replies = append(n.Edges.Replies, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tuq *TwitterUserQuery) loadReplies(ctx context.Context, query *TweetQuery, nodes []*TwitterUser, init func(*TwitterUser), assign func(*TwitterUser, *Tweet)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TwitterUser)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tweet.FieldReplyTwitterUserID)
	}
	query.Where(predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(twitteruser.RepliesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReplyTwitterUserID
		if fk == nil {
			return fmt.Errorf(`foreign-key "reply_twitter_user_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "reply_twitter_user_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tuq *TwitterUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tuq.querySpec()
	_spec.Node.Columns = tuq.ctx.Fields
	if len(tuq.ctx.Fields) > 0 {
		_spec.Unique = tuq.ctx.Unique != nil && *tuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tuq.driver, _spec)
}

func (tuq *TwitterUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twitteruser.Table, twitteruser.Columns, sqlgraph.NewFieldSpec(twitteruser.FieldID, field.TypeInt))
	_spec.From = tuq.sql
	if unique := tuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tuq.path != nil {
		_spec.Unique = true
	}
	if fields := tuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twitteruser.FieldID)
		for i := range fields {
			if fields[i] != twitteruser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tuq *TwitterUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tuq.driver.Dialect())
	t1 := builder.Table(twitteruser.Table)
	columns := tuq.ctx.Fields
	if len(columns) == 0 {
		columns = twitteruser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tuq.sql != nil {
		selector = tuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tuq.ctx.Unique != nil && *tuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tuq.predicates {
		p(selector)
	}
	for _, p := range tuq.order {
		p(selector)
	}
	if offset := tuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TwitterUserGroupBy is the group-by builder for TwitterUser entities.
type TwitterUserGroupBy struct {
	selector
	build *TwitterUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tugb *TwitterUserGroupBy) Aggregate(fns ...AggregateFunc) *TwitterUserGroupBy {
	tugb.fns = append(tugb.fns, fns...)
	return tugb
}

// Scan applies the selector query and scans the result into the given value.
func (tugb *TwitterUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tugb.build.ctx, ent.OpQueryGroupBy)
	if err := tugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterUserQuery, *TwitterUserGroupBy](ctx, tugb.build, tugb, tugb.build.inters, v)
}

func (tugb *TwitterUserGroupBy) sqlScan(ctx context.Context, root *TwitterUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tugb.fns))
	for _, fn := range tugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tugb.flds)+len(tugb.fns))
		for _, f := range *tugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwitterUserSelect is the builder for selecting fields of TwitterUser entities.
type TwitterUserSelect struct {
	*TwitterUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tus *TwitterUserSelect) Aggregate(fns ...AggregateFunc) *TwitterUserSelect {
	tus.fns = append(tus.fns, fns...)
	return tus
}

// Scan applies the selector query and scans the result into the given value.
func (tus *TwitterUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tus.ctx, ent.OpQuerySelect)
	if err := tus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterUserQuery, *TwitterUserSelect](ctx, tus.TwitterUserQuery, tus, tus.inters, v)
}

func (tus *TwitterUserSelect) sqlScan(ctx context.Context, root *TwitterUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tus.fns))
	for _, fn := range tus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
