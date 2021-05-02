package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").MinLen(5).MaxLen(50).NotEmpty().Unique(),
		field.Int32("balance").Default(0.0),
		field.String("currency").NotEmpty(),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entries", Entry.Type),
		edge.To("outbound", Transfer.Type),
		edge.To("inbound", Transfer.Type),
	}
}
