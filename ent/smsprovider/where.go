// Code generated by ent, DO NOT EDIT.

package smsprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/huuhoait/simple-admin-message-center/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldName, v))
}

// SecretID applies equality check predicate on the "secret_id" field. It's identical to SecretIDEQ.
func SecretID(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldSecretID, v))
}

// SecretKey applies equality check predicate on the "secret_key" field. It's identical to SecretKeyEQ.
func SecretKey(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldSecretKey, v))
}

// Region applies equality check predicate on the "region" field. It's identical to RegionEQ.
func Region(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldRegion, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldIsDefault, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContainsFold(FieldName, v))
}

// SecretIDEQ applies the EQ predicate on the "secret_id" field.
func SecretIDEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldSecretID, v))
}

// SecretIDNEQ applies the NEQ predicate on the "secret_id" field.
func SecretIDNEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldSecretID, v))
}

// SecretIDIn applies the In predicate on the "secret_id" field.
func SecretIDIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldSecretID, vs...))
}

// SecretIDNotIn applies the NotIn predicate on the "secret_id" field.
func SecretIDNotIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldSecretID, vs...))
}

// SecretIDGT applies the GT predicate on the "secret_id" field.
func SecretIDGT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldSecretID, v))
}

// SecretIDGTE applies the GTE predicate on the "secret_id" field.
func SecretIDGTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldSecretID, v))
}

// SecretIDLT applies the LT predicate on the "secret_id" field.
func SecretIDLT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldSecretID, v))
}

// SecretIDLTE applies the LTE predicate on the "secret_id" field.
func SecretIDLTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldSecretID, v))
}

// SecretIDContains applies the Contains predicate on the "secret_id" field.
func SecretIDContains(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContains(FieldSecretID, v))
}

// SecretIDHasPrefix applies the HasPrefix predicate on the "secret_id" field.
func SecretIDHasPrefix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasPrefix(FieldSecretID, v))
}

// SecretIDHasSuffix applies the HasSuffix predicate on the "secret_id" field.
func SecretIDHasSuffix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasSuffix(FieldSecretID, v))
}

// SecretIDEqualFold applies the EqualFold predicate on the "secret_id" field.
func SecretIDEqualFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEqualFold(FieldSecretID, v))
}

// SecretIDContainsFold applies the ContainsFold predicate on the "secret_id" field.
func SecretIDContainsFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContainsFold(FieldSecretID, v))
}

// SecretKeyEQ applies the EQ predicate on the "secret_key" field.
func SecretKeyEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldSecretKey, v))
}

// SecretKeyNEQ applies the NEQ predicate on the "secret_key" field.
func SecretKeyNEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldSecretKey, v))
}

// SecretKeyIn applies the In predicate on the "secret_key" field.
func SecretKeyIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldSecretKey, vs...))
}

// SecretKeyNotIn applies the NotIn predicate on the "secret_key" field.
func SecretKeyNotIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldSecretKey, vs...))
}

// SecretKeyGT applies the GT predicate on the "secret_key" field.
func SecretKeyGT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldSecretKey, v))
}

// SecretKeyGTE applies the GTE predicate on the "secret_key" field.
func SecretKeyGTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldSecretKey, v))
}

// SecretKeyLT applies the LT predicate on the "secret_key" field.
func SecretKeyLT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldSecretKey, v))
}

// SecretKeyLTE applies the LTE predicate on the "secret_key" field.
func SecretKeyLTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldSecretKey, v))
}

// SecretKeyContains applies the Contains predicate on the "secret_key" field.
func SecretKeyContains(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContains(FieldSecretKey, v))
}

// SecretKeyHasPrefix applies the HasPrefix predicate on the "secret_key" field.
func SecretKeyHasPrefix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasPrefix(FieldSecretKey, v))
}

// SecretKeyHasSuffix applies the HasSuffix predicate on the "secret_key" field.
func SecretKeyHasSuffix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasSuffix(FieldSecretKey, v))
}

// SecretKeyEqualFold applies the EqualFold predicate on the "secret_key" field.
func SecretKeyEqualFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEqualFold(FieldSecretKey, v))
}

// SecretKeyContainsFold applies the ContainsFold predicate on the "secret_key" field.
func SecretKeyContainsFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContainsFold(FieldSecretKey, v))
}

// RegionEQ applies the EQ predicate on the "region" field.
func RegionEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldRegion, v))
}

// RegionNEQ applies the NEQ predicate on the "region" field.
func RegionNEQ(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldRegion, v))
}

// RegionIn applies the In predicate on the "region" field.
func RegionIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldIn(FieldRegion, vs...))
}

// RegionNotIn applies the NotIn predicate on the "region" field.
func RegionNotIn(vs ...string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNotIn(FieldRegion, vs...))
}

// RegionGT applies the GT predicate on the "region" field.
func RegionGT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGT(FieldRegion, v))
}

// RegionGTE applies the GTE predicate on the "region" field.
func RegionGTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldGTE(FieldRegion, v))
}

// RegionLT applies the LT predicate on the "region" field.
func RegionLT(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLT(FieldRegion, v))
}

// RegionLTE applies the LTE predicate on the "region" field.
func RegionLTE(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldLTE(FieldRegion, v))
}

// RegionContains applies the Contains predicate on the "region" field.
func RegionContains(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContains(FieldRegion, v))
}

// RegionHasPrefix applies the HasPrefix predicate on the "region" field.
func RegionHasPrefix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasPrefix(FieldRegion, v))
}

// RegionHasSuffix applies the HasSuffix predicate on the "region" field.
func RegionHasSuffix(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldHasSuffix(FieldRegion, v))
}

// RegionEqualFold applies the EqualFold predicate on the "region" field.
func RegionEqualFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEqualFold(FieldRegion, v))
}

// RegionContainsFold applies the ContainsFold predicate on the "region" field.
func RegionContainsFold(v string) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldContainsFold(FieldRegion, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.SmsProvider {
	return predicate.SmsProvider(sql.FieldNEQ(FieldIsDefault, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SmsProvider) predicate.SmsProvider {
	return predicate.SmsProvider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SmsProvider) predicate.SmsProvider {
	return predicate.SmsProvider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SmsProvider) predicate.SmsProvider {
	return predicate.SmsProvider(sql.NotPredicates(p))
}
