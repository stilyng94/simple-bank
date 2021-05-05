// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntryCreate is the builder for creating a Entry entity.
type EntryCreate struct {
	config
	mutation *EntryMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ec *EntryCreate) SetCreateTime(t time.Time) *EntryCreate {
	ec.mutation.SetCreateTime(t)
	return ec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ec *EntryCreate) SetNillableCreateTime(t *time.Time) *EntryCreate {
	if t != nil {
		ec.SetCreateTime(*t)
	}
	return ec
}

// SetUpdateTime sets the "update_time" field.
func (ec *EntryCreate) SetUpdateTime(t time.Time) *EntryCreate {
	ec.mutation.SetUpdateTime(t)
	return ec
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ec *EntryCreate) SetNillableUpdateTime(t *time.Time) *EntryCreate {
	if t != nil {
		ec.SetUpdateTime(*t)
	}
	return ec
}

// SetAmount sets the "amount" field.
func (ec *EntryCreate) SetAmount(f float64) *EntryCreate {
	ec.mutation.SetAmount(f)
	return ec
}

// SetAccountId sets the "accountId" field.
func (ec *EntryCreate) SetAccountId(u uuid.UUID) *EntryCreate {
	ec.mutation.SetAccountId(u)
	return ec
}

// SetID sets the "id" field.
func (ec *EntryCreate) SetID(u uuid.UUID) *EntryCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (ec *EntryCreate) SetAccountID(id uuid.UUID) *EntryCreate {
	ec.mutation.SetAccountID(id)
	return ec
}

// SetAccount sets the "account" edge to the Account entity.
func (ec *EntryCreate) SetAccount(a *Account) *EntryCreate {
	return ec.SetAccountID(a.ID)
}

// Mutation returns the EntryMutation object of the builder.
func (ec *EntryCreate) Mutation() *EntryMutation {
	return ec.mutation
}

// Save creates the Entry in the database.
func (ec *EntryCreate) Save(ctx context.Context) (*Entry, error) {
	var (
		err  error
		node *Entry
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EntryCreate) SaveX(ctx context.Context) *Entry {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ec *EntryCreate) defaults() {
	if _, ok := ec.mutation.CreateTime(); !ok {
		v := entry.DefaultCreateTime()
		ec.mutation.SetCreateTime(v)
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		v := entry.DefaultUpdateTime()
		ec.mutation.SetUpdateTime(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := entry.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EntryCreate) check() error {
	if _, ok := ec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := ec.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if _, ok := ec.mutation.AccountId(); !ok {
		return &ValidationError{Name: "accountId", err: errors.New("ent: missing required field \"accountId\"")}
	}
	if _, ok := ec.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New("ent: missing required edge \"account\"")}
	}
	return nil
}

func (ec *EntryCreate) sqlSave(ctx context.Context) (*Entry, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ec *EntryCreate) createSpec() (*Entry, *sqlgraph.CreateSpec) {
	var (
		_node = &Entry{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entry.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entry.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entry.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ec.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entry.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ec.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: entry.FieldAmount,
		})
		_node.Amount = value
	}
	if nodes := ec.mutation.AccountIDs(); len(nodes) > 0 {
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
		_node.AccountId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EntryCreateBulk is the builder for creating many Entry entities in bulk.
type EntryCreateBulk struct {
	config
	builders []*EntryCreate
}

// Save creates the Entry entities in the database.
func (ecb *EntryCreateBulk) Save(ctx context.Context) ([]*Entry, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Entry, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EntryCreateBulk) SaveX(ctx context.Context) []*Entry {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
