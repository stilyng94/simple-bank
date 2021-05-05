package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type IDMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New)}
}
