package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AppUserEmailTemplate holds the schema definition for the AppUserEmailTemplate entity.
type AppUserEmailTemplate struct {
	ent.Schema
}

// Fields of the AppUserEmailTemplate.
func (AppUserEmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("subject"),
		field.String("body"),
		field.String("sender"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
	}
}

// Edges of the AppUserEmailTemplate.
func (AppUserEmailTemplate) Edges() []ent.Edge {
	return nil
}
