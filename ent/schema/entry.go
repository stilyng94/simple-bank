package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Entry holds the schema definition for the Entry entity.
type Entry struct {
	ent.Schema
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("amount"),
		// field.UUID("account_id", uuid.UUID{}).Immutable().Unique(),
	}

}

func (Entry) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return []ent.Edge{edge.From("account", Account.Type).Ref("entries").Unique().Required()}
}
