package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.String("username").MinLen(1).Unique(),
		field.Bytes("password_hash").MinLen(1).Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connections", Connection.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{StructTag: `json:"-"`},
	}
}
