// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appcontact"
	"github.com/google/uuid"
)

// AppContact is the model entity for the AppContact schema.
type AppContact struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Account holds the value of the "account" field.
	Account string `json:"account,omitempty"`
	// AccountType holds the value of the "account_type" field.
	AccountType string `json:"account_type,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppContact) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appcontact.FieldCreateAt, appcontact.FieldUpdateAt:
			values[i] = new(sql.NullInt64)
		case appcontact.FieldUsedFor, appcontact.FieldSender, appcontact.FieldAccount, appcontact.FieldAccountType:
			values[i] = new(sql.NullString)
		case appcontact.FieldID, appcontact.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppContact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppContact fields.
func (ac *AppContact) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appcontact.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ac.ID = *value
			}
		case appcontact.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ac.AppID = *value
			}
		case appcontact.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				ac.UsedFor = value.String
			}
		case appcontact.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				ac.Sender = value.String
			}
		case appcontact.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				ac.Account = value.String
			}
		case appcontact.FieldAccountType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_type", values[i])
			} else if value.Valid {
				ac.AccountType = value.String
			}
		case appcontact.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ac.CreateAt = uint32(value.Int64)
			}
		case appcontact.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ac.UpdateAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppContact.
// Note that you need to call AppContact.Unwrap() before calling this method if this AppContact
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppContact) Update() *AppContactUpdateOne {
	return (&AppContactClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AppContact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppContact) Unwrap() *AppContact {
	tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppContact is not a transactional entity")
	}
	ac.config.driver = tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppContact) String() string {
	var builder strings.Builder
	builder.WriteString("AppContact(")
	builder.WriteString(fmt.Sprintf("id=%v", ac.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.AppID))
	builder.WriteString(", used_for=")
	builder.WriteString(ac.UsedFor)
	builder.WriteString(", sender=")
	builder.WriteString(ac.Sender)
	builder.WriteString(", account=")
	builder.WriteString(ac.Account)
	builder.WriteString(", account_type=")
	builder.WriteString(ac.AccountType)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.UpdateAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppContacts is a parsable slice of AppContact.
type AppContacts []*AppContact

func (ac AppContacts) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
