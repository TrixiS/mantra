package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Connection holds the schema definition for the Connection entity.
type Connection struct {
	ent.Schema
}

// Fields of the Connection.
func (Connection) Fields() []ent.Field {
	return []ent.Field{
		field.Int("local_id"),
		field.String("name").MinLen(1),
		field.String("host").MinLen(1),
		field.Uint("port"),
		field.String("user").MinLen(1),
		// TODO: we need to encrypt connection password with user password ok cool (search for some algo)
		field.String("password").MinLen(1), // this is bad, but we need it
		field.String("args"),
	}
}

// Edges of the Connection.
func (Connection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("connections").
			Required().
			Unique().
			Annotations(edge.Annotation{
				StructTag: `json:"-"`,
			}),
	}
}
