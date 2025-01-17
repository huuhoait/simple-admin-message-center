// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/huuhoait/simple-admin-message-center/ent/smslog"
)

// SmsLogCreate is the builder for creating a SmsLog entity.
type SmsLogCreate struct {
	config
	mutation *SmsLogMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (slc *SmsLogCreate) SetCreatedAt(t time.Time) *SmsLogCreate {
	slc.mutation.SetCreatedAt(t)
	return slc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (slc *SmsLogCreate) SetNillableCreatedAt(t *time.Time) *SmsLogCreate {
	if t != nil {
		slc.SetCreatedAt(*t)
	}
	return slc
}

// SetUpdatedAt sets the "updated_at" field.
func (slc *SmsLogCreate) SetUpdatedAt(t time.Time) *SmsLogCreate {
	slc.mutation.SetUpdatedAt(t)
	return slc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (slc *SmsLogCreate) SetNillableUpdatedAt(t *time.Time) *SmsLogCreate {
	if t != nil {
		slc.SetUpdatedAt(*t)
	}
	return slc
}

// SetPhoneNumber sets the "phone_number" field.
func (slc *SmsLogCreate) SetPhoneNumber(s string) *SmsLogCreate {
	slc.mutation.SetPhoneNumber(s)
	return slc
}

// SetContent sets the "content" field.
func (slc *SmsLogCreate) SetContent(s string) *SmsLogCreate {
	slc.mutation.SetContent(s)
	return slc
}

// SetSendStatus sets the "send_status" field.
func (slc *SmsLogCreate) SetSendStatus(u uint8) *SmsLogCreate {
	slc.mutation.SetSendStatus(u)
	return slc
}

// SetProvider sets the "provider" field.
func (slc *SmsLogCreate) SetProvider(s string) *SmsLogCreate {
	slc.mutation.SetProvider(s)
	return slc
}

// SetID sets the "id" field.
func (slc *SmsLogCreate) SetID(u uuid.UUID) *SmsLogCreate {
	slc.mutation.SetID(u)
	return slc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (slc *SmsLogCreate) SetNillableID(u *uuid.UUID) *SmsLogCreate {
	if u != nil {
		slc.SetID(*u)
	}
	return slc
}

// Mutation returns the SmsLogMutation object of the builder.
func (slc *SmsLogCreate) Mutation() *SmsLogMutation {
	return slc.mutation
}

// Save creates the SmsLog in the database.
func (slc *SmsLogCreate) Save(ctx context.Context) (*SmsLog, error) {
	slc.defaults()
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *SmsLogCreate) SaveX(ctx context.Context) *SmsLog {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *SmsLogCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *SmsLogCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *SmsLogCreate) defaults() {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		v := smslog.DefaultCreatedAt()
		slc.mutation.SetCreatedAt(v)
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		v := smslog.DefaultUpdatedAt()
		slc.mutation.SetUpdatedAt(v)
	}
	if _, ok := slc.mutation.ID(); !ok {
		v := smslog.DefaultID()
		slc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *SmsLogCreate) check() error {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SmsLog.created_at"`)}
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SmsLog.updated_at"`)}
	}
	if _, ok := slc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "SmsLog.phone_number"`)}
	}
	if _, ok := slc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "SmsLog.content"`)}
	}
	if _, ok := slc.mutation.SendStatus(); !ok {
		return &ValidationError{Name: "send_status", err: errors.New(`ent: missing required field "SmsLog.send_status"`)}
	}
	if _, ok := slc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "SmsLog.provider"`)}
	}
	return nil
}

func (slc *SmsLogCreate) sqlSave(ctx context.Context) (*SmsLog, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *SmsLogCreate) createSpec() (*SmsLog, *sqlgraph.CreateSpec) {
	var (
		_node = &SmsLog{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(smslog.Table, sqlgraph.NewFieldSpec(smslog.FieldID, field.TypeUUID))
	)
	if id, ok := slc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := slc.mutation.CreatedAt(); ok {
		_spec.SetField(smslog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := slc.mutation.UpdatedAt(); ok {
		_spec.SetField(smslog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := slc.mutation.PhoneNumber(); ok {
		_spec.SetField(smslog.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := slc.mutation.Content(); ok {
		_spec.SetField(smslog.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := slc.mutation.SendStatus(); ok {
		_spec.SetField(smslog.FieldSendStatus, field.TypeUint8, value)
		_node.SendStatus = value
	}
	if value, ok := slc.mutation.Provider(); ok {
		_spec.SetField(smslog.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	return _node, _spec
}

// SmsLogCreateBulk is the builder for creating many SmsLog entities in bulk.
type SmsLogCreateBulk struct {
	config
	err      error
	builders []*SmsLogCreate
}

// Save creates the SmsLog entities in the database.
func (slcb *SmsLogCreateBulk) Save(ctx context.Context) ([]*SmsLog, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*SmsLog, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SmsLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *SmsLogCreateBulk) SaveX(ctx context.Context) []*SmsLog {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *SmsLogCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *SmsLogCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}
