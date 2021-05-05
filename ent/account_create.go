// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"simple-bank/ent/transfer"
	"simple-bank/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ac *AccountCreate) SetCreateTime(t time.Time) *AccountCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AccountCreate) SetUpdateTime(t time.Time) *AccountCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdateTime(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetOwner sets the "owner" field.
func (ac *AccountCreate) SetOwner(s string) *AccountCreate {
	ac.mutation.SetOwner(s)
	return ac
}

// SetBalance sets the "balance" field.
func (ac *AccountCreate) SetBalance(f float64) *AccountCreate {
	ac.mutation.SetBalance(f)
	return ac
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (ac *AccountCreate) SetNillableBalance(f *float64) *AccountCreate {
	if f != nil {
		ac.SetBalance(*f)
	}
	return ac
}

// SetCurrency sets the "currency" field.
func (ac *AccountCreate) SetCurrency(s string) *AccountCreate {
	ac.mutation.SetCurrency(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(u uuid.UUID) *AccountCreate {
	ac.mutation.SetID(u)
	return ac
}

// AddEntryIDs adds the "entries" edge to the Entry entity by IDs.
func (ac *AccountCreate) AddEntryIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddEntryIDs(ids...)
	return ac
}

// AddEntries adds the "entries" edges to the Entry entity.
func (ac *AccountCreate) AddEntries(e ...*Entry) *AccountCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ac.AddEntryIDs(ids...)
}

// AddOutboundIDs adds the "outbounds" edge to the Transfer entity by IDs.
func (ac *AccountCreate) AddOutboundIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddOutboundIDs(ids...)
	return ac
}

// AddOutbounds adds the "outbounds" edges to the Transfer entity.
func (ac *AccountCreate) AddOutbounds(t ...*Transfer) *AccountCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddOutboundIDs(ids...)
}

// AddInboundIDs adds the "inbounds" edge to the Transfer entity by IDs.
func (ac *AccountCreate) AddInboundIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddInboundIDs(ids...)
	return ac
}

// AddInbounds adds the "inbounds" edges to the Transfer entity.
func (ac *AccountCreate) AddInbounds(t ...*Transfer) *AccountCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddInboundIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AccountCreate) SetUserID(id string) *AccountCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AccountCreate) SetUser(u *User) *AccountCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := account.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := account.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
	if _, ok := ac.mutation.Balance(); !ok {
		v := account.DefaultBalance
		ac.mutation.SetBalance(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := account.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := ac.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("ent: missing required field \"owner\"")}
	}
	if v, ok := ac.mutation.Owner(); ok {
		if err := account.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if _, ok := ac.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New("ent: missing required field \"balance\"")}
	}
	if _, ok := ac.mutation.Currency(); !ok {
		return &ValidationError{Name: "currency", err: errors.New("ent: missing required field \"currency\"")}
	}
	if v, ok := ac.mutation.Currency(); ok {
		if err := account.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf("ent: validator failed for field \"currency\": %w", err)}
		}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: account.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := ac.mutation.Currency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldCurrency,
		})
		_node.Currency = value
	}
	if nodes := ac.mutation.EntriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.EntriesTable,
			Columns: []string{account.EntriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.OutboundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.OutboundsTable,
			Columns: []string{account.OutboundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.InboundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.InboundsTable,
			Columns: []string{account.InboundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transfer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Owner = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
