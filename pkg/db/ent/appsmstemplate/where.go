// Code generated by entc, DO NOT EDIT.

package appsmstemplate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// LangID applies equality check predicate on the "lang_id" field. It's identical to LangIDEQ.
func LangID(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// UsedFor applies equality check predicate on the "used_for" field. It's identical to UsedForEQ.
func UsedFor(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubject), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// LangIDEQ applies the EQ predicate on the "lang_id" field.
func LangIDEQ(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLangID), v))
	})
}

// LangIDNEQ applies the NEQ predicate on the "lang_id" field.
func LangIDNEQ(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLangID), v))
	})
}

// LangIDIn applies the In predicate on the "lang_id" field.
func LangIDIn(vs ...uuid.UUID) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLangID), v...))
	})
}

// LangIDNotIn applies the NotIn predicate on the "lang_id" field.
func LangIDNotIn(vs ...uuid.UUID) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLangID), v...))
	})
}

// LangIDGT applies the GT predicate on the "lang_id" field.
func LangIDGT(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLangID), v))
	})
}

// LangIDGTE applies the GTE predicate on the "lang_id" field.
func LangIDGTE(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLangID), v))
	})
}

// LangIDLT applies the LT predicate on the "lang_id" field.
func LangIDLT(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLangID), v))
	})
}

// LangIDLTE applies the LTE predicate on the "lang_id" field.
func LangIDLTE(v uuid.UUID) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLangID), v))
	})
}

// UsedForEQ applies the EQ predicate on the "used_for" field.
func UsedForEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsedFor), v))
	})
}

// UsedForNEQ applies the NEQ predicate on the "used_for" field.
func UsedForNEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsedFor), v))
	})
}

// UsedForIn applies the In predicate on the "used_for" field.
func UsedForIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsedFor), v...))
	})
}

// UsedForNotIn applies the NotIn predicate on the "used_for" field.
func UsedForNotIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsedFor), v...))
	})
}

// UsedForGT applies the GT predicate on the "used_for" field.
func UsedForGT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsedFor), v))
	})
}

// UsedForGTE applies the GTE predicate on the "used_for" field.
func UsedForGTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsedFor), v))
	})
}

// UsedForLT applies the LT predicate on the "used_for" field.
func UsedForLT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsedFor), v))
	})
}

// UsedForLTE applies the LTE predicate on the "used_for" field.
func UsedForLTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsedFor), v))
	})
}

// UsedForContains applies the Contains predicate on the "used_for" field.
func UsedForContains(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsedFor), v))
	})
}

// UsedForHasPrefix applies the HasPrefix predicate on the "used_for" field.
func UsedForHasPrefix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsedFor), v))
	})
}

// UsedForHasSuffix applies the HasSuffix predicate on the "used_for" field.
func UsedForHasSuffix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsedFor), v))
	})
}

// UsedForEqualFold applies the EqualFold predicate on the "used_for" field.
func UsedForEqualFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsedFor), v))
	})
}

// UsedForContainsFold applies the ContainsFold predicate on the "used_for" field.
func UsedForContainsFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsedFor), v))
	})
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubject), v))
	})
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubject), v))
	})
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubject), v...))
	})
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubject), v...))
	})
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubject), v))
	})
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubject), v))
	})
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubject), v))
	})
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubject), v))
	})
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubject), v))
	})
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubject), v))
	})
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubject), v))
	})
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubject), v))
	})
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubject), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.AppSMSTemplate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppSMSTemplate) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppSMSTemplate) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppSMSTemplate) predicate.AppSMSTemplate {
	return predicate.AppSMSTemplate(func(s *sql.Selector) {
		p(s.Not())
	})
}
