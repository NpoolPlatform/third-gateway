package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppContact holds the schema definition for the AppContact entity.
type AppContact struct {
	ent.Schema
}

// Fields of the AppContact.
func (AppContact) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.String("used_for").MaxLen(32),
		field.String("sender"),
		field.String("account"),
		field.String("account_type"),
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

// Edges of the AppContact.
func (AppContact) Edges() []ent.Edge {
	return nil
}

func (AppContact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "account", "used_for", "account_type").
			Unique(),
	}
}
