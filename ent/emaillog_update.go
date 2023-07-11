// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-message-center/ent/emaillog"
	"github.com/suyuan32/simple-admin-message-center/ent/predicate"
)

// EmailLogUpdate is the builder for updating EmailLog entities.
type EmailLogUpdate struct {
	config
	hooks    []Hook
	mutation *EmailLogMutation
}

// Where appends a list predicates to the EmailLogUpdate builder.
func (elu *EmailLogUpdate) Where(ps ...predicate.EmailLog) *EmailLogUpdate {
	elu.mutation.Where(ps...)
	return elu
}

// SetUpdatedAt sets the "updated_at" field.
func (elu *EmailLogUpdate) SetUpdatedAt(t time.Time) *EmailLogUpdate {
	elu.mutation.SetUpdatedAt(t)
	return elu
}

// SetTarget sets the "target" field.
func (elu *EmailLogUpdate) SetTarget(s string) *EmailLogUpdate {
	elu.mutation.SetTarget(s)
	return elu
}

// SetSubject sets the "subject" field.
func (elu *EmailLogUpdate) SetSubject(s string) *EmailLogUpdate {
	elu.mutation.SetSubject(s)
	return elu
}

// SetContent sets the "content" field.
func (elu *EmailLogUpdate) SetContent(s string) *EmailLogUpdate {
	elu.mutation.SetContent(s)
	return elu
}

// SetSendStatus sets the "send_status" field.
func (elu *EmailLogUpdate) SetSendStatus(u uint8) *EmailLogUpdate {
	elu.mutation.ResetSendStatus()
	elu.mutation.SetSendStatus(u)
	return elu
}

// AddSendStatus adds u to the "send_status" field.
func (elu *EmailLogUpdate) AddSendStatus(u int8) *EmailLogUpdate {
	elu.mutation.AddSendStatus(u)
	return elu
}

// SetProvider sets the "provider" field.
func (elu *EmailLogUpdate) SetProvider(s string) *EmailLogUpdate {
	elu.mutation.SetProvider(s)
	return elu
}

// Mutation returns the EmailLogMutation object of the builder.
func (elu *EmailLogUpdate) Mutation() *EmailLogMutation {
	return elu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (elu *EmailLogUpdate) Save(ctx context.Context) (int, error) {
	elu.defaults()
	return withHooks(ctx, elu.sqlSave, elu.mutation, elu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (elu *EmailLogUpdate) SaveX(ctx context.Context) int {
	affected, err := elu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (elu *EmailLogUpdate) Exec(ctx context.Context) error {
	_, err := elu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elu *EmailLogUpdate) ExecX(ctx context.Context) {
	if err := elu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (elu *EmailLogUpdate) defaults() {
	if _, ok := elu.mutation.UpdatedAt(); !ok {
		v := emaillog.UpdateDefaultUpdatedAt()
		elu.mutation.SetUpdatedAt(v)
	}
}

func (elu *EmailLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(emaillog.Table, emaillog.Columns, sqlgraph.NewFieldSpec(emaillog.FieldID, field.TypeUUID))
	if ps := elu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := elu.mutation.UpdatedAt(); ok {
		_spec.SetField(emaillog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := elu.mutation.Target(); ok {
		_spec.SetField(emaillog.FieldTarget, field.TypeString, value)
	}
	if value, ok := elu.mutation.Subject(); ok {
		_spec.SetField(emaillog.FieldSubject, field.TypeString, value)
	}
	if value, ok := elu.mutation.Content(); ok {
		_spec.SetField(emaillog.FieldContent, field.TypeString, value)
	}
	if value, ok := elu.mutation.SendStatus(); ok {
		_spec.SetField(emaillog.FieldSendStatus, field.TypeUint8, value)
	}
	if value, ok := elu.mutation.AddedSendStatus(); ok {
		_spec.AddField(emaillog.FieldSendStatus, field.TypeUint8, value)
	}
	if value, ok := elu.mutation.Provider(); ok {
		_spec.SetField(emaillog.FieldProvider, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, elu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emaillog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	elu.mutation.done = true
	return n, nil
}

// EmailLogUpdateOne is the builder for updating a single EmailLog entity.
type EmailLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailLogMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (eluo *EmailLogUpdateOne) SetUpdatedAt(t time.Time) *EmailLogUpdateOne {
	eluo.mutation.SetUpdatedAt(t)
	return eluo
}

// SetTarget sets the "target" field.
func (eluo *EmailLogUpdateOne) SetTarget(s string) *EmailLogUpdateOne {
	eluo.mutation.SetTarget(s)
	return eluo
}

// SetSubject sets the "subject" field.
func (eluo *EmailLogUpdateOne) SetSubject(s string) *EmailLogUpdateOne {
	eluo.mutation.SetSubject(s)
	return eluo
}

// SetContent sets the "content" field.
func (eluo *EmailLogUpdateOne) SetContent(s string) *EmailLogUpdateOne {
	eluo.mutation.SetContent(s)
	return eluo
}

// SetSendStatus sets the "send_status" field.
func (eluo *EmailLogUpdateOne) SetSendStatus(u uint8) *EmailLogUpdateOne {
	eluo.mutation.ResetSendStatus()
	eluo.mutation.SetSendStatus(u)
	return eluo
}

// AddSendStatus adds u to the "send_status" field.
func (eluo *EmailLogUpdateOne) AddSendStatus(u int8) *EmailLogUpdateOne {
	eluo.mutation.AddSendStatus(u)
	return eluo
}

// SetProvider sets the "provider" field.
func (eluo *EmailLogUpdateOne) SetProvider(s string) *EmailLogUpdateOne {
	eluo.mutation.SetProvider(s)
	return eluo
}

// Mutation returns the EmailLogMutation object of the builder.
func (eluo *EmailLogUpdateOne) Mutation() *EmailLogMutation {
	return eluo.mutation
}

// Where appends a list predicates to the EmailLogUpdate builder.
func (eluo *EmailLogUpdateOne) Where(ps ...predicate.EmailLog) *EmailLogUpdateOne {
	eluo.mutation.Where(ps...)
	return eluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eluo *EmailLogUpdateOne) Select(field string, fields ...string) *EmailLogUpdateOne {
	eluo.fields = append([]string{field}, fields...)
	return eluo
}

// Save executes the query and returns the updated EmailLog entity.
func (eluo *EmailLogUpdateOne) Save(ctx context.Context) (*EmailLog, error) {
	eluo.defaults()
	return withHooks(ctx, eluo.sqlSave, eluo.mutation, eluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eluo *EmailLogUpdateOne) SaveX(ctx context.Context) *EmailLog {
	node, err := eluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eluo *EmailLogUpdateOne) Exec(ctx context.Context) error {
	_, err := eluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eluo *EmailLogUpdateOne) ExecX(ctx context.Context) {
	if err := eluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eluo *EmailLogUpdateOne) defaults() {
	if _, ok := eluo.mutation.UpdatedAt(); !ok {
		v := emaillog.UpdateDefaultUpdatedAt()
		eluo.mutation.SetUpdatedAt(v)
	}
}

func (eluo *EmailLogUpdateOne) sqlSave(ctx context.Context) (_node *EmailLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(emaillog.Table, emaillog.Columns, sqlgraph.NewFieldSpec(emaillog.FieldID, field.TypeUUID))
	id, ok := eluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emaillog.FieldID)
		for _, f := range fields {
			if !emaillog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emaillog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eluo.mutation.UpdatedAt(); ok {
		_spec.SetField(emaillog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eluo.mutation.Target(); ok {
		_spec.SetField(emaillog.FieldTarget, field.TypeString, value)
	}
	if value, ok := eluo.mutation.Subject(); ok {
		_spec.SetField(emaillog.FieldSubject, field.TypeString, value)
	}
	if value, ok := eluo.mutation.Content(); ok {
		_spec.SetField(emaillog.FieldContent, field.TypeString, value)
	}
	if value, ok := eluo.mutation.SendStatus(); ok {
		_spec.SetField(emaillog.FieldSendStatus, field.TypeUint8, value)
	}
	if value, ok := eluo.mutation.AddedSendStatus(); ok {
		_spec.AddField(emaillog.FieldSendStatus, field.TypeUint8, value)
	}
	if value, ok := eluo.mutation.Provider(); ok {
		_spec.SetField(emaillog.FieldProvider, field.TypeString, value)
	}
	_node = &EmailLog{config: eluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emaillog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eluo.mutation.done = true
	return _node, nil
}
