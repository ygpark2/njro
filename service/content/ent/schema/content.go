package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(
				entproto.Field(1),
			),
		field.String("board_id").
			NotEmpty().
			Immutable().
			Annotations(
				entproto.Field(2),
			),
		field.String("post_id").
			NotEmpty().
			Immutable().
			Annotations(
				entproto.Field(3),
			),
		field.String("comment_id").
			Default("").
			Annotations(
				entproto.Field(4),
			),
		field.Text("content").
			NotEmpty().
			Annotations(
				entproto.Field(5),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entproto.Field(6),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entproto.Field(7),
			),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return nil
}

// Annotations of the Content.
func (Content) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
