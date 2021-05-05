package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().MaxLen(25).MinLen(2).NotEmpty().Match(regexp.MustCompile("[a-zA-Z_]+$")).StorageKey("username").StructTag(`json:"username,omitempty"`),
		field.String("password").NotEmpty().Sensitive().MaxLen(256),
		field.String("full_name").MaxLen(75).MinLen(2).NotEmpty(),
		field.String("email").Unique().MaxLen(256).MinLen(10).NotEmpty(),
		field.Time("password_changed_at").Default(time.Now),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type).StorageKey(edge.Column("owner")),
	}
}
