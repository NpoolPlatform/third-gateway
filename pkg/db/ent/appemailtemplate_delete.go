// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appemailtemplate"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/predicate"
)

// AppEmailTemplateDelete is the builder for deleting a AppEmailTemplate entity.
type AppEmailTemplateDelete struct {
	config
	hooks    []Hook
	mutation *AppEmailTemplateMutation
}

// Where appends a list predicates to the AppEmailTemplateDelete builder.
func (aetd *AppEmailTemplateDelete) Where(ps ...predicate.AppEmailTemplate) *AppEmailTemplateDelete {
	aetd.mutation.Where(ps...)
	return aetd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aetd *AppEmailTemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aetd.hooks) == 0 {
		affected, err = aetd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppEmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aetd.mutation = mutation
			affected, err = aetd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aetd.hooks) - 1; i >= 0; i-- {
			if aetd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aetd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aetd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aetd *AppEmailTemplateDelete) ExecX(ctx context.Context) int {
	n, err := aetd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aetd *AppEmailTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appemailtemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: appemailtemplate.FieldID,
			},
		},
	}
	if ps := aetd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aetd.driver, _spec)
}

// AppEmailTemplateDeleteOne is the builder for deleting a single AppEmailTemplate entity.
type AppEmailTemplateDeleteOne struct {
	aetd *AppEmailTemplateDelete
}

// Exec executes the deletion query.
func (aetdo *AppEmailTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := aetdo.aetd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appemailtemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aetdo *AppEmailTemplateDeleteOne) ExecX(ctx context.Context) {
	aetdo.aetd.ExecX(ctx)
}
