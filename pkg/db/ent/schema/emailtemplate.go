package schema

import "entgo.io/ent"

// EmailTemplate holds the schema definition for the EmailTemplate entity.
type EmailTemplate struct {
	ent.Schema
}

// Fields of the EmailTemplate.
func (EmailTemplate) Fields() []ent.Field {
	return nil
}

// Edges of the EmailTemplate.
func (EmailTemplate) Edges() []ent.Edge {
	return nil
}
