package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/contrib/entproto"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		// entproto.Service(), // also generate a gRPC service definition
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
		Annotations(
			entproto.Field(1),
		),
		field.String("name").
		Annotations(
			entproto.Field(2),
		),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin", User.Type).
			Unique().
			StorageKey(edge.Column("admin_id")).
			Annotations(entproto.Field(3)),
	}
}
