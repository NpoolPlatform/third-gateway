// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppContactsColumns holds the columns for the "app_contacts" table.
	AppContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString, Size: 32},
		{Name: "sender", Type: field.TypeString},
		{Name: "account", Type: field.TypeString},
		{Name: "account_type", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
	}
	// AppContactsTable holds the schema information for the "app_contacts" table.
	AppContactsTable = &schema.Table{
		Name:       "app_contacts",
		Columns:    AppContactsColumns,
		PrimaryKey: []*schema.Column{AppContactsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appcontact_app_id_account_used_for_account_type",
				Unique:  true,
				Columns: []*schema.Column{AppContactsColumns[1], AppContactsColumns[4], AppContactsColumns[2], AppContactsColumns[5]},
			},
		},
	}
	// AppEmailTemplatesColumns holds the columns for the "app_email_templates" table.
	AppEmailTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "default_to_username", Type: field.TypeString},
		{Name: "used_for", Type: field.TypeString},
		{Name: "sender", Type: field.TypeString},
		{Name: "reply_tos", Type: field.TypeJSON},
		{Name: "cc_tos", Type: field.TypeJSON},
		{Name: "subject", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 8192},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
	}
	// AppEmailTemplatesTable holds the schema information for the "app_email_templates" table.
	AppEmailTemplatesTable = &schema.Table{
		Name:       "app_email_templates",
		Columns:    AppEmailTemplatesColumns,
		PrimaryKey: []*schema.Column{AppEmailTemplatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appemailtemplate_app_id_lang_id_used_for",
				Unique:  true,
				Columns: []*schema.Column{AppEmailTemplatesColumns[1], AppEmailTemplatesColumns[2], AppEmailTemplatesColumns[4]},
			},
		},
	}
	// AppSmsTemplatesColumns holds the columns for the "app_sms_templates" table.
	AppSmsTemplatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "lang_id", Type: field.TypeUUID},
		{Name: "used_for", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
	}
	// AppSmsTemplatesTable holds the schema information for the "app_sms_templates" table.
	AppSmsTemplatesTable = &schema.Table{
		Name:       "app_sms_templates",
		Columns:    AppSmsTemplatesColumns,
		PrimaryKey: []*schema.Column{AppSmsTemplatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appsmstemplate_app_id_lang_id_used_for",
				Unique:  true,
				Columns: []*schema.Column{AppSmsTemplatesColumns[1], AppSmsTemplatesColumns[2], AppSmsTemplatesColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppContactsTable,
		AppEmailTemplatesTable,
		AppSmsTemplatesTable,
	}
)

func init() {
}
