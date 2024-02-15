// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/nixxxon/entdemo/ent/todo"
	"github.com/nixxxon/entdemo/ent/todohack"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   todo.Table,
			Columns: todo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: todo.FieldID,
			},
		},
		Type: "Todo",
		Fields: map[string]*sqlgraph.FieldSpec{
			todo.FieldOtherID: {Type: field.TypeUUID, Column: todo.FieldOtherID},
			todo.FieldName:    {Type: field.TypeString, Column: todo.FieldName},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   todohack.Table,
			Columns: todohack.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: todohack.FieldID,
			},
		},
		Type: "TodoHack",
		Fields: map[string]*sqlgraph.FieldSpec{
			todohack.FieldHistoryTime: {Type: field.TypeTime, Column: todohack.FieldHistoryTime},
			todohack.FieldRef:         {Type: field.TypeUUID, Column: todohack.FieldRef},
			todohack.FieldOperation:   {Type: field.TypeEnum, Column: todohack.FieldOperation},
			todohack.FieldOtherID:     {Type: field.TypeUUID, Column: todohack.FieldOtherID},
			todohack.FieldName:        {Type: field.TypeString, Column: todohack.FieldName},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (tq *TodoQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TodoQuery builder.
func (tq *TodoQuery) Filter() *TodoFilter {
	return &TodoFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TodoMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TodoMutation builder.
func (m *TodoMutation) Filter() *TodoFilter {
	return &TodoFilter{config: m.config, predicateAdder: m}
}

// TodoFilter provides a generic filtering capability at runtime for TodoQuery.
type TodoFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TodoFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TodoFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(todo.FieldID))
}

// WhereOtherID applies the entql [16]byte predicate on the other_id field.
func (f *TodoFilter) WhereOtherID(p entql.ValueP) {
	f.Where(p.Field(todo.FieldOtherID))
}

// WhereName applies the entql string predicate on the name field.
func (f *TodoFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(todo.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (thq *TodoHackQuery) addPredicate(pred func(s *sql.Selector)) {
	thq.predicates = append(thq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TodoHackQuery builder.
func (thq *TodoHackQuery) Filter() *TodoHackFilter {
	return &TodoHackFilter{config: thq.config, predicateAdder: thq}
}

// addPredicate implements the predicateAdder interface.
func (m *TodoHackMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TodoHackMutation builder.
func (m *TodoHackMutation) Filter() *TodoHackFilter {
	return &TodoHackFilter{config: m.config, predicateAdder: m}
}

// TodoHackFilter provides a generic filtering capability at runtime for TodoHackQuery.
type TodoHackFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TodoHackFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TodoHackFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(todohack.FieldID))
}

// WhereHistoryTime applies the entql time.Time predicate on the history_time field.
func (f *TodoHackFilter) WhereHistoryTime(p entql.TimeP) {
	f.Where(p.Field(todohack.FieldHistoryTime))
}

// WhereRef applies the entql [16]byte predicate on the ref field.
func (f *TodoHackFilter) WhereRef(p entql.ValueP) {
	f.Where(p.Field(todohack.FieldRef))
}

// WhereOperation applies the entql string predicate on the operation field.
func (f *TodoHackFilter) WhereOperation(p entql.StringP) {
	f.Where(p.Field(todohack.FieldOperation))
}

// WhereOtherID applies the entql [16]byte predicate on the other_id field.
func (f *TodoHackFilter) WhereOtherID(p entql.ValueP) {
	f.Where(p.Field(todohack.FieldOtherID))
}

// WhereName applies the entql string predicate on the name field.
func (f *TodoHackFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(todohack.FieldName))
}