// Code generated by entc, DO NOT EDIT.

package appemailtemplate

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appemailtemplate type in the database.
	Label = "app_email_template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldLangID holds the string denoting the lang_id field in the database.
	FieldLangID = "lang_id"
	// FieldDefaultToUsername holds the string denoting the default_to_username field in the database.
	FieldDefaultToUsername = "default_to_username"
	// FieldUsedFor holds the string denoting the used_for field in the database.
	FieldUsedFor = "used_for"
	// FieldSender holds the string denoting the sender field in the database.
	FieldSender = "sender"
	// FieldReplyTos holds the string denoting the reply_tos field in the database.
	FieldReplyTos = "reply_tos"
	// FieldCcTos holds the string denoting the cc_tos field in the database.
	FieldCcTos = "cc_tos"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the appemailtemplate in the database.
	Table = "app_email_templates"
)

// Columns holds all SQL columns for appemailtemplate fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldLangID,
	FieldDefaultToUsername,
	FieldUsedFor,
	FieldSender,
	FieldReplyTos,
	FieldCcTos,
	FieldSubject,
	FieldBody,
	FieldCreateAt,
	FieldUpdateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// BodyValidator is a validator for the "body" field. It is called by the builders before save.
	BodyValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
