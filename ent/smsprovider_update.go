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
	"github.com/huuhoait/simple-admin-message-center/ent/predicate"
	"github.com/huuhoait/simple-admin-message-center/ent/smsprovider"
)

// SmsProviderUpdate is the builder for updating SmsProvider entities.
type SmsProviderUpdate struct {
	config
	hooks    []Hook
	mutation *SmsProviderMutation
}

// Where appends a list predicates to the SmsProviderUpdate builder.
func (spu *SmsProviderUpdate) Where(ps ...predicate.SmsProvider) *SmsProviderUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetUpdatedAt sets the "updated_at" field.
func (spu *SmsProviderUpdate) SetUpdatedAt(t time.Time) *SmsProviderUpdate {
	spu.mutation.SetUpdatedAt(t)
	return spu
}

// SetName sets the "name" field.
func (spu *SmsProviderUpdate) SetName(s string) *SmsProviderUpdate {
	spu.mutation.SetName(s)
	return spu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (spu *SmsProviderUpdate) SetNillableName(s *string) *SmsProviderUpdate {
	if s != nil {
		spu.SetName(*s)
	}
	return spu
}

// SetSecretID sets the "secret_id" field.
func (spu *SmsProviderUpdate) SetSecretID(s string) *SmsProviderUpdate {
	spu.mutation.SetSecretID(s)
	return spu
}

// SetNillableSecretID sets the "secret_id" field if the given value is not nil.
func (spu *SmsProviderUpdate) SetNillableSecretID(s *string) *SmsProviderUpdate {
	if s != nil {
		spu.SetSecretID(*s)
	}
	return spu
}

// SetSecretKey sets the "secret_key" field.
func (spu *SmsProviderUpdate) SetSecretKey(s string) *SmsProviderUpdate {
	spu.mutation.SetSecretKey(s)
	return spu
}

// SetNillableSecretKey sets the "secret_key" field if the given value is not nil.
func (spu *SmsProviderUpdate) SetNillableSecretKey(s *string) *SmsProviderUpdate {
	if s != nil {
		spu.SetSecretKey(*s)
	}
	return spu
}

// SetRegion sets the "region" field.
func (spu *SmsProviderUpdate) SetRegion(s string) *SmsProviderUpdate {
	spu.mutation.SetRegion(s)
	return spu
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (spu *SmsProviderUpdate) SetNillableRegion(s *string) *SmsProviderUpdate {
	if s != nil {
		spu.SetRegion(*s)
	}
	return spu
}

// SetIsDefault sets the "is_default" field.
func (spu *SmsProviderUpdate) SetIsDefault(b bool) *SmsProviderUpdate {
	spu.mutation.SetIsDefault(b)
	return spu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (spu *SmsProviderUpdate) SetNillableIsDefault(b *bool) *SmsProviderUpdate {
	if b != nil {
		spu.SetIsDefault(*b)
	}
	return spu
}

// Mutation returns the SmsProviderMutation object of the builder.
func (spu *SmsProviderUpdate) Mutation() *SmsProviderMutation {
	return spu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *SmsProviderUpdate) Save(ctx context.Context) (int, error) {
	spu.defaults()
	return withHooks(ctx, spu.sqlSave, spu.mutation, spu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spu *SmsProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *SmsProviderUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *SmsProviderUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spu *SmsProviderUpdate) defaults() {
	if _, ok := spu.mutation.UpdatedAt(); !ok {
		v := smsprovider.UpdateDefaultUpdatedAt()
		spu.mutation.SetUpdatedAt(v)
	}
}

func (spu *SmsProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(smsprovider.Table, smsprovider.Columns, sqlgraph.NewFieldSpec(smsprovider.FieldID, field.TypeUint64))
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.UpdatedAt(); ok {
		_spec.SetField(smsprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := spu.mutation.Name(); ok {
		_spec.SetField(smsprovider.FieldName, field.TypeString, value)
	}
	if value, ok := spu.mutation.SecretID(); ok {
		_spec.SetField(smsprovider.FieldSecretID, field.TypeString, value)
	}
	if value, ok := spu.mutation.SecretKey(); ok {
		_spec.SetField(smsprovider.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := spu.mutation.Region(); ok {
		_spec.SetField(smsprovider.FieldRegion, field.TypeString, value)
	}
	if value, ok := spu.mutation.IsDefault(); ok {
		_spec.SetField(smsprovider.FieldIsDefault, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smsprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	spu.mutation.done = true
	return n, nil
}

// SmsProviderUpdateOne is the builder for updating a single SmsProvider entity.
type SmsProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SmsProviderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (spuo *SmsProviderUpdateOne) SetUpdatedAt(t time.Time) *SmsProviderUpdateOne {
	spuo.mutation.SetUpdatedAt(t)
	return spuo
}

// SetName sets the "name" field.
func (spuo *SmsProviderUpdateOne) SetName(s string) *SmsProviderUpdateOne {
	spuo.mutation.SetName(s)
	return spuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (spuo *SmsProviderUpdateOne) SetNillableName(s *string) *SmsProviderUpdateOne {
	if s != nil {
		spuo.SetName(*s)
	}
	return spuo
}

// SetSecretID sets the "secret_id" field.
func (spuo *SmsProviderUpdateOne) SetSecretID(s string) *SmsProviderUpdateOne {
	spuo.mutation.SetSecretID(s)
	return spuo
}

// SetNillableSecretID sets the "secret_id" field if the given value is not nil.
func (spuo *SmsProviderUpdateOne) SetNillableSecretID(s *string) *SmsProviderUpdateOne {
	if s != nil {
		spuo.SetSecretID(*s)
	}
	return spuo
}

// SetSecretKey sets the "secret_key" field.
func (spuo *SmsProviderUpdateOne) SetSecretKey(s string) *SmsProviderUpdateOne {
	spuo.mutation.SetSecretKey(s)
	return spuo
}

// SetNillableSecretKey sets the "secret_key" field if the given value is not nil.
func (spuo *SmsProviderUpdateOne) SetNillableSecretKey(s *string) *SmsProviderUpdateOne {
	if s != nil {
		spuo.SetSecretKey(*s)
	}
	return spuo
}

// SetRegion sets the "region" field.
func (spuo *SmsProviderUpdateOne) SetRegion(s string) *SmsProviderUpdateOne {
	spuo.mutation.SetRegion(s)
	return spuo
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (spuo *SmsProviderUpdateOne) SetNillableRegion(s *string) *SmsProviderUpdateOne {
	if s != nil {
		spuo.SetRegion(*s)
	}
	return spuo
}

// SetIsDefault sets the "is_default" field.
func (spuo *SmsProviderUpdateOne) SetIsDefault(b bool) *SmsProviderUpdateOne {
	spuo.mutation.SetIsDefault(b)
	return spuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (spuo *SmsProviderUpdateOne) SetNillableIsDefault(b *bool) *SmsProviderUpdateOne {
	if b != nil {
		spuo.SetIsDefault(*b)
	}
	return spuo
}

// Mutation returns the SmsProviderMutation object of the builder.
func (spuo *SmsProviderUpdateOne) Mutation() *SmsProviderMutation {
	return spuo.mutation
}

// Where appends a list predicates to the SmsProviderUpdate builder.
func (spuo *SmsProviderUpdateOne) Where(ps ...predicate.SmsProvider) *SmsProviderUpdateOne {
	spuo.mutation.Where(ps...)
	return spuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *SmsProviderUpdateOne) Select(field string, fields ...string) *SmsProviderUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated SmsProvider entity.
func (spuo *SmsProviderUpdateOne) Save(ctx context.Context) (*SmsProvider, error) {
	spuo.defaults()
	return withHooks(ctx, spuo.sqlSave, spuo.mutation, spuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *SmsProviderUpdateOne) SaveX(ctx context.Context) *SmsProvider {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *SmsProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *SmsProviderUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spuo *SmsProviderUpdateOne) defaults() {
	if _, ok := spuo.mutation.UpdatedAt(); !ok {
		v := smsprovider.UpdateDefaultUpdatedAt()
		spuo.mutation.SetUpdatedAt(v)
	}
}

func (spuo *SmsProviderUpdateOne) sqlSave(ctx context.Context) (_node *SmsProvider, err error) {
	_spec := sqlgraph.NewUpdateSpec(smsprovider.Table, smsprovider.Columns, sqlgraph.NewFieldSpec(smsprovider.FieldID, field.TypeUint64))
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SmsProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smsprovider.FieldID)
		for _, f := range fields {
			if !smsprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != smsprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spuo.mutation.UpdatedAt(); ok {
		_spec.SetField(smsprovider.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := spuo.mutation.Name(); ok {
		_spec.SetField(smsprovider.FieldName, field.TypeString, value)
	}
	if value, ok := spuo.mutation.SecretID(); ok {
		_spec.SetField(smsprovider.FieldSecretID, field.TypeString, value)
	}
	if value, ok := spuo.mutation.SecretKey(); ok {
		_spec.SetField(smsprovider.FieldSecretKey, field.TypeString, value)
	}
	if value, ok := spuo.mutation.Region(); ok {
		_spec.SetField(smsprovider.FieldRegion, field.TypeString, value)
	}
	if value, ok := spuo.mutation.IsDefault(); ok {
		_spec.SetField(smsprovider.FieldIsDefault, field.TypeBool, value)
	}
	_node = &SmsProvider{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smsprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	spuo.mutation.done = true
	return _node, nil
}
