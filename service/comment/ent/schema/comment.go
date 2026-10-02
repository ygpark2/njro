package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
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
		field.String("content_id").
			Default("").
			Annotations(
				entproto.Field(4),
			),
		field.String("userid").
			NotEmpty().
			Immutable().
			Annotations(
				entproto.Field(5),
			),
		field.String("username").
			Default("").
			Annotations(
				entproto.Field(6),
			),
		field.String("nickname").
			Default("").
			Annotations(
				entproto.Field(7),
			),
		field.String("email").
			Default("").
			Annotations(
				entproto.Field(8),
			),
		field.String("password").
			Optional().
			Default("").
			Annotations(
				entproto.Field(9),
			),
		field.Text("content").
			Default("").
			Annotations(
				entproto.Field(10),
			),
		field.String("url").
			Optional().
			Default("").
			Annotations(
				entproto.Field(11),
			),
		field.Bool("use_html").
			Default(false).
			Annotations(
				entproto.Field(12),
			),
		field.Bool("use_secret").
			Default(false).
			Annotations(
				entproto.Field(13),
			),
		field.Uint32("up_vote_count").
			Default(0).
			Annotations(
				entproto.Field(14),
			),
		field.Uint32("down_vote_count").
			Default(0).
			Annotations(
				entproto.Field(15),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entproto.Field(16),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entproto.Field(17),
			),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

// Annotations of the Comment.
func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
