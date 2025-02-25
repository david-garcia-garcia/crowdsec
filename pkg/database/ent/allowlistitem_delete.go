// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/allowlistitem"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
)

// AllowListItemDelete is the builder for deleting a AllowListItem entity.
type AllowListItemDelete struct {
	config
	hooks    []Hook
	mutation *AllowListItemMutation
}

// Where appends a list predicates to the AllowListItemDelete builder.
func (alid *AllowListItemDelete) Where(ps ...predicate.AllowListItem) *AllowListItemDelete {
	alid.mutation.Where(ps...)
	return alid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (alid *AllowListItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, alid.sqlExec, alid.mutation, alid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (alid *AllowListItemDelete) ExecX(ctx context.Context) int {
	n, err := alid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (alid *AllowListItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(allowlistitem.Table, sqlgraph.NewFieldSpec(allowlistitem.FieldID, field.TypeInt))
	if ps := alid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, alid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	alid.mutation.done = true
	return affected, err
}

// AllowListItemDeleteOne is the builder for deleting a single AllowListItem entity.
type AllowListItemDeleteOne struct {
	alid *AllowListItemDelete
}

// Where appends a list predicates to the AllowListItemDelete builder.
func (alido *AllowListItemDeleteOne) Where(ps ...predicate.AllowListItem) *AllowListItemDeleteOne {
	alido.alid.mutation.Where(ps...)
	return alido
}

// Exec executes the deletion query.
func (alido *AllowListItemDeleteOne) Exec(ctx context.Context) error {
	n, err := alido.alid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{allowlistitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (alido *AllowListItemDeleteOne) ExecX(ctx context.Context) {
	if err := alido.Exec(ctx); err != nil {
		panic(err)
	}
}
