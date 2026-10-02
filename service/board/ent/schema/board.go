package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(
				entproto.Field(1),
			),
		field.String("title").
			NotEmpty().
			Annotations(
				entproto.Field(2),
			),
		field.String("mobile_title").
			NotEmpty().
			Annotations(
				entproto.Field(3),
			),
		field.Uint32("order").
			Default(0).
			Annotations(
				entproto.Field(4),
			),
		field.Bool("search").
			Default(true).
			Annotations(
				entproto.Field(5),
			),
		field.String("description").
			Default("").
			Annotations(
				entproto.Field(6),
			),
		field.JSON("notices", []string{}).
			Optional().
			Annotations(
				entproto.Field(7),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entproto.Field(8),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entproto.Field(9),
			),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return nil
}

// Annotations of the Board.
func (Board) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
