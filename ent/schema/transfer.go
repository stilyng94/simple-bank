package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Transfer holds the schema definition for the Transfer entity.
type Transfer struct {
	ent.Schema
}

// Fields of the Transfer.
func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("amount"),
		// field.UUID("from_account_id", uuid.UUID{}).Immutable().Unique(),
		// field.UUID("to_account_id", uuid.UUID{}).Immutable().Unique(),
	}
}
func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Transfer.
func (Transfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fromAccount", Account.Type).Ref("outbound").Unique().Required(),
		edge.From("toAccount", Account.Type).Ref("inbound").Unique().Required(),
	}
}
