package schema

import (
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique().
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				return err
			}),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Organization.Type).
			Required(),
		edge.From("organizations", Organization.Type).
			Ref("users"),
	}
}
