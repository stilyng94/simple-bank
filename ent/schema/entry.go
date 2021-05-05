package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.UUID("accountId", uuid.UUID{}).Immutable(),
	}

}

func (Entry) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return []ent.Edge{edge.From("account", Account.Type).Ref("entries").Unique().Required().Field("accountId")}
}
