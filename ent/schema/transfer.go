package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Transfer holds the schema definition for the Transfer entity.
type Transfer struct {
	ent.Schema
}

// Fields of the Transfer.
func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.UUID("fromAccountId", uuid.UUID{}).Immutable(),
		field.UUID("toAccountId", uuid.UUID{}).Immutable(),
	}
}
func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Transfer.
func (Transfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fromAccount", Account.Type).Ref("outbounds").Unique().Required().Field("fromAccountId"),
		edge.From("toAccount", Account.Type).Ref("inbounds").Unique().Required().Field("toAccountId"),
	}
}
