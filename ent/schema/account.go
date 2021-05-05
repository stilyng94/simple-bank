package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").MaxLen(25).MinLen(2).NotEmpty(),
		field.Float("balance").Default(0.0),
		field.String("currency").NotEmpty().Immutable(),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{IDMixin{}, mixin.Time{}}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entries", Entry.Type).StorageKey(edge.Column("account_id")),
		edge.To("outbounds", Transfer.Type).StorageKey(edge.Column("from_account_id")),
		edge.To("inbounds", Transfer.Type).StorageKey(edge.Column("to_account_id")),
		edge.From("user", User.Type).Ref("accounts").Required().Unique().Field("owner"),
	}
}

func (Account) Indexes() []ent.Index {

	return []ent.Index{

		index.Fields("owner", "currency").
			Unique().StorageKey("owner_currency_key"),
	}

}
