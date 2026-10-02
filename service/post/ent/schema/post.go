package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
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
		field.String("title").
			NotEmpty().
			Annotations(
				entproto.Field(3),
			),
		field.String("slug").
			Default("").
			Annotations(
				entproto.Field(4),
			),
		field.String("category").
			Default("general").
			Annotations(
				entproto.Field(5),
			),
		field.Text("content").
			Default("").
			Annotations(
				entproto.Field(6),
			),
		field.String("email").
			Default("").
			Annotations(
				entproto.Field(7),
			),
		field.String("writer").
			NotEmpty().
			Immutable().
			Annotations(
				entproto.Field(8),
			),
		field.String("password").
			Optional().
			Default("").
			Annotations(
				entproto.Field(9),
			),
		field.JSON("tags", []string{}).
			Optional().
			Annotations(
				entproto.Field(10),
			),
		field.Uint32("read_counts").
			Default(0).
			Annotations(
				entproto.Field(11),
			),
		field.Uint32("comment_counts").
			Default(0).
			Annotations(
				entproto.Field(12),
			),
		field.Uint32("up_vote_counts").
			Default(0).
			Annotations(
				entproto.Field(13),
			),
		field.Uint32("down_vote_counts").
			Default(0).
			Annotations(
				entproto.Field(14),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entproto.Field(15),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entproto.Field(16),
			),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

// Annotations of the Post.
func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
