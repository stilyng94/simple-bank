// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Entry is the model entity for the Entry schema.
type Entry struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// AccountId holds the value of the "accountId" field.
	AccountId uuid.UUID `json:"accountId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntryQuery when eager-loading is set.
	Edges EntryEdges `json:"edges"`
}

// EntryEdges holds the relations/edges for other nodes in the graph.
type EntryEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntryEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// The edge account was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Entry) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case entry.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case entry.FieldCreateTime, entry.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case entry.FieldID, entry.FieldAccountId:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Entry", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Entry fields.
func (e *Entry) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entry.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case entry.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		case entry.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = value.Time
			}
		case entry.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				e.Amount = value.Float64
			}
		case entry.FieldAccountId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field accountId", values[i])
			} else if value != nil {
				e.AccountId = *value
			}
		}
	}
	return nil
}

// QueryAccount queries the "account" edge of the Entry entity.
func (e *Entry) QueryAccount() *AccountQuery {
	return (&EntryClient{config: e.config}).QueryAccount(e)
}

// Update returns a builder for updating this Entry.
// Note that you need to call Entry.Unwrap() before calling this method if this Entry
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Entry) Update() *EntryUpdateOne {
	return (&EntryClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Entry entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Entry) Unwrap() *Entry {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Entry is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Entry) String() string {
	var builder strings.Builder
	builder.WriteString("Entry(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", e.Amount))
	builder.WriteString(", accountId=")
	builder.WriteString(fmt.Sprintf("%v", e.AccountId))
	builder.WriteByte(')')
	return builder.String()
}

// Entries is a parsable slice of Entry.
type Entries []*Entry

func (e Entries) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
