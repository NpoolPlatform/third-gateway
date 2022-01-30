// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appsmstemplate"
	"github.com/google/uuid"
)

// AppSMSTemplate is the model entity for the AppSMSTemplate schema.
type AppSMSTemplate struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// LangID holds the value of the "lang_id" field.
	LangID uuid.UUID `json:"lang_id,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppSMSTemplate) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appsmstemplate.FieldCreateAt, appsmstemplate.FieldUpdateAt:
			values[i] = new(sql.NullInt64)
		case appsmstemplate.FieldUsedFor, appsmstemplate.FieldSubject, appsmstemplate.FieldMessage:
			values[i] = new(sql.NullString)
		case appsmstemplate.FieldID, appsmstemplate.FieldAppID, appsmstemplate.FieldLangID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppSMSTemplate", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppSMSTemplate fields.
func (ast *AppSMSTemplate) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appsmstemplate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ast.ID = *value
			}
		case appsmstemplate.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ast.AppID = *value
			}
		case appsmstemplate.FieldLangID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field lang_id", values[i])
			} else if value != nil {
				ast.LangID = *value
			}
		case appsmstemplate.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				ast.UsedFor = value.String
			}
		case appsmstemplate.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				ast.Subject = value.String
			}
		case appsmstemplate.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				ast.Message = value.String
			}
		case appsmstemplate.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ast.CreateAt = uint32(value.Int64)
			}
		case appsmstemplate.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ast.UpdateAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppSMSTemplate.
// Note that you need to call AppSMSTemplate.Unwrap() before calling this method if this AppSMSTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (ast *AppSMSTemplate) Update() *AppSMSTemplateUpdateOne {
	return (&AppSMSTemplateClient{config: ast.config}).UpdateOne(ast)
}

// Unwrap unwraps the AppSMSTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ast *AppSMSTemplate) Unwrap() *AppSMSTemplate {
	tx, ok := ast.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppSMSTemplate is not a transactional entity")
	}
	ast.config.driver = tx.drv
	return ast
}

// String implements the fmt.Stringer.
func (ast *AppSMSTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("AppSMSTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v", ast.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ast.AppID))
	builder.WriteString(", lang_id=")
	builder.WriteString(fmt.Sprintf("%v", ast.LangID))
	builder.WriteString(", used_for=")
	builder.WriteString(ast.UsedFor)
	builder.WriteString(", subject=")
	builder.WriteString(ast.Subject)
	builder.WriteString(", message=")
	builder.WriteString(ast.Message)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ast.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ast.UpdateAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppSMSTemplates is a parsable slice of AppSMSTemplate.
type AppSMSTemplates []*AppSMSTemplate

func (ast AppSMSTemplates) config(cfg config) {
	for _i := range ast {
		ast[_i].config = cfg
	}
}
