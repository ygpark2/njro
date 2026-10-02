package schema

import (
	"context"
	"strings"
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(
				entproto.Field(1),
			),
		field.String("username").
			NotEmpty().
			Unique().
			Immutable().
			Annotations(
				entproto.Field(2),
			),
		field.String("first_name").
			Default("").
			Annotations(
				entproto.Field(3),
			),
		field.String("last_name").
			Default("").
			Annotations(
				entproto.Field(4),
			),
		field.String("email").
			NotEmpty().
			Unique().
			Immutable().
			Annotations(
				entproto.Field(5),
			),
		field.String("password").
			NotEmpty().
			Sensitive().
			Optional().
			Annotations(
				entproto.Field(6),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entproto.Field(7),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entproto.Field(8),
			),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				type passwordSetter interface {
					Password() (string, bool)
					SetPassword(string)
				}
				if ps, ok := m.(passwordSetter); ok {
					if pwd, ok := ps.Password(); ok && pwd != "" && !strings.HasPrefix(pwd, "$2a$") {
						hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
						if err != nil {
							return nil, err
						}
						ps.SetPassword(string(hashed))
					}
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
