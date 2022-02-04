// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appemailtemplate"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appsmstemplate"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appemailtemplateFields := schema.AppEmailTemplate{}.Fields()
	_ = appemailtemplateFields
	// appemailtemplateDescBody is the schema descriptor for body field.
	appemailtemplateDescBody := appemailtemplateFields[9].Descriptor()
	// appemailtemplate.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	appemailtemplate.BodyValidator = appemailtemplateDescBody.Validators[0].(func(string) error)
	// appemailtemplateDescCreateAt is the schema descriptor for create_at field.
	appemailtemplateDescCreateAt := appemailtemplateFields[10].Descriptor()
	// appemailtemplate.DefaultCreateAt holds the default value on creation for the create_at field.
	appemailtemplate.DefaultCreateAt = appemailtemplateDescCreateAt.Default.(func() uint32)
	// appemailtemplateDescUpdateAt is the schema descriptor for update_at field.
	appemailtemplateDescUpdateAt := appemailtemplateFields[11].Descriptor()
	// appemailtemplate.DefaultUpdateAt holds the default value on creation for the update_at field.
	appemailtemplate.DefaultUpdateAt = appemailtemplateDescUpdateAt.Default.(func() uint32)
	// appemailtemplate.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appemailtemplate.UpdateDefaultUpdateAt = appemailtemplateDescUpdateAt.UpdateDefault.(func() uint32)
	// appemailtemplateDescID is the schema descriptor for id field.
	appemailtemplateDescID := appemailtemplateFields[0].Descriptor()
	// appemailtemplate.DefaultID holds the default value on creation for the id field.
	appemailtemplate.DefaultID = appemailtemplateDescID.Default.(func() uuid.UUID)
	appsmstemplateFields := schema.AppSMSTemplate{}.Fields()
	_ = appsmstemplateFields
	// appsmstemplateDescCreateAt is the schema descriptor for create_at field.
	appsmstemplateDescCreateAt := appsmstemplateFields[6].Descriptor()
	// appsmstemplate.DefaultCreateAt holds the default value on creation for the create_at field.
	appsmstemplate.DefaultCreateAt = appsmstemplateDescCreateAt.Default.(func() uint32)
	// appsmstemplateDescUpdateAt is the schema descriptor for update_at field.
	appsmstemplateDescUpdateAt := appsmstemplateFields[7].Descriptor()
	// appsmstemplate.DefaultUpdateAt holds the default value on creation for the update_at field.
	appsmstemplate.DefaultUpdateAt = appsmstemplateDescUpdateAt.Default.(func() uint32)
	// appsmstemplate.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appsmstemplate.UpdateDefaultUpdateAt = appsmstemplateDescUpdateAt.UpdateDefault.(func() uint32)
	// appsmstemplateDescID is the schema descriptor for id field.
	appsmstemplateDescID := appsmstemplateFields[0].Descriptor()
	// appsmstemplate.DefaultID holds the default value on creation for the id field.
	appsmstemplate.DefaultID = appsmstemplateDescID.Default.(func() uuid.UUID)
}
