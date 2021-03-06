// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"simple-bank/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntryUpdate is the builder for updating Entry entities.
type EntryUpdate struct {
	config
	hooks    []Hook
	mutation *EntryMutation
}

// Where adds a new predicate for the EntryUpdate builder.
func (eu *EntryUpdate) Where(ps ...predicate.Entry) *EntryUpdate {
	eu.mutation.predicates = append(eu.mutation.predicates, ps...)
	return eu
}

// SetAmount sets the "amount" field.
func (eu *EntryUpdate) SetAmount(f float64) *EntryUpdate {
	eu.mutation.ResetAmount()
	eu.mutation.SetAmount(f)
	return eu
}

// AddAmount adds f to the "amount" field.
func (eu *EntryUpdate) AddAmount(f float64) *EntryUpdate {
	eu.mutation.AddAmount(f)
	return eu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (eu *EntryUpdate) SetAccountID(id uuid.UUID) *EntryUpdate {
	eu.mutation.SetAccountID(id)
	return eu
}

// SetAccount sets the "account" edge to the Account entity.
func (eu *EntryUpdate) SetAccount(a *Account) *EntryUpdate {
	return eu.SetAccountID(a.ID)
}

// Mutation returns the EntryMutation object of the builder.
func (eu *EntryUpdate) Mutation() *EntryMutation {
	return eu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (eu *EntryUpdate) ClearAccount() *EntryUpdate {
	eu.mutation.ClearAccount()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eu.defaults()
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntryUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntryUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntryUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EntryUpdate) defaults() {
	if _, ok := eu.mutation.UpdateTime(); !ok {
		v := entry.UpdateDefaultUpdateTime()
		eu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EntryUpdate) check() error {
	if _, ok := eu.mutation.AccountID(); eu.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	return nil
}

func (eu *EntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entry.Table,
			Columns: entry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entry.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entry.FieldUpdateTime,
		})
	}
	if value, ok := eu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entry.FieldAmount,
		})
	}
	if value, ok := eu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entry.FieldAmount,
		})
	}
	if eu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.AccountTable,
			Columns: []string{entry.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.AccountTable,
			Columns: []string{entry.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entry.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EntryUpdateOne is the builder for updating a single Entry entity.
type EntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntryMutation
}

// SetAmount sets the "amount" field.
func (euo *EntryUpdateOne) SetAmount(f float64) *EntryUpdateOne {
	euo.mutation.ResetAmount()
	euo.mutation.SetAmount(f)
	return euo
}

// AddAmount adds f to the "amount" field.
func (euo *EntryUpdateOne) AddAmount(f float64) *EntryUpdateOne {
	euo.mutation.AddAmount(f)
	return euo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (euo *EntryUpdateOne) SetAccountID(id uuid.UUID) *EntryUpdateOne {
	euo.mutation.SetAccountID(id)
	return euo
}

// SetAccount sets the "account" edge to the Account entity.
func (euo *EntryUpdateOne) SetAccount(a *Account) *EntryUpdateOne {
	return euo.SetAccountID(a.ID)
}

// Mutation returns the EntryMutation object of the builder.
func (euo *EntryUpdateOne) Mutation() *EntryMutation {
	return euo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (euo *EntryUpdateOne) ClearAccount() *EntryUpdateOne {
	euo.mutation.ClearAccount()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntryUpdateOne) Select(field string, fields ...string) *EntryUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entry entity.
func (euo *EntryUpdateOne) Save(ctx context.Context) (*Entry, error) {
	var (
		err  error
		node *Entry
	)
	euo.defaults()
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntryUpdateOne) SaveX(ctx context.Context) *Entry {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntryUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntryUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EntryUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdateTime(); !ok {
		v := entry.UpdateDefaultUpdateTime()
		euo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EntryUpdateOne) check() error {
	if _, ok := euo.mutation.AccountID(); euo.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	return nil
}

func (euo *EntryUpdateOne) sqlSave(ctx context.Context) (_node *Entry, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entry.Table,
			Columns: entry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entry.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Entry.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entry.FieldID)
		for _, f := range fields {
			if !entry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entry.FieldUpdateTime,
		})
	}
	if value, ok := euo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entry.FieldAmount,
		})
	}
	if value, ok := euo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entry.FieldAmount,
		})
	}
	if euo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.AccountTable,
			Columns: []string{entry.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entry.AccountTable,
			Columns: []string{entry.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Entry{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entry.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
