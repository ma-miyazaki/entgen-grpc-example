package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/contrib/entproto"
	"github.com/oklog/ulid/v2"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // also generate a gRPC service definition
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return ulid.Make().String()
			}).
			Annotations(
				entproto.Field(1),
			),
		field.Int("age").
			Annotations(
				entproto.Field(2),
			),
		field.String("name").
			Annotations(
				entproto.Field(3),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("administered", Category.Type).
			Ref("admin").
			Annotations(entproto.Field(4)),
	}
}
