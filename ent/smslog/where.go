// Code generated by ent, DO NOT EDIT.

package smslog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/huuhoait/simple-admin-message-center/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldPhoneNumber, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldContent, v))
}

// SendStatus applies equality check predicate on the "send_status" field. It's identical to SendStatusEQ.
func SendStatus(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldSendStatus, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldProvider, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContainsFold(FieldContent, v))
}

// SendStatusEQ applies the EQ predicate on the "send_status" field.
func SendStatusEQ(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldSendStatus, v))
}

// SendStatusNEQ applies the NEQ predicate on the "send_status" field.
func SendStatusNEQ(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldSendStatus, v))
}

// SendStatusIn applies the In predicate on the "send_status" field.
func SendStatusIn(vs ...uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldSendStatus, vs...))
}

// SendStatusNotIn applies the NotIn predicate on the "send_status" field.
func SendStatusNotIn(vs ...uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldSendStatus, vs...))
}

// SendStatusGT applies the GT predicate on the "send_status" field.
func SendStatusGT(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldSendStatus, v))
}

// SendStatusGTE applies the GTE predicate on the "send_status" field.
func SendStatusGTE(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldSendStatus, v))
}

// SendStatusLT applies the LT predicate on the "send_status" field.
func SendStatusLT(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldSendStatus, v))
}

// SendStatusLTE applies the LTE predicate on the "send_status" field.
func SendStatusLTE(v uint8) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldSendStatus, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.SmsLog {
	return predicate.SmsLog(sql.FieldContainsFold(FieldProvider, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SmsLog) predicate.SmsLog {
	return predicate.SmsLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SmsLog) predicate.SmsLog {
	return predicate.SmsLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SmsLog) predicate.SmsLog {
	return predicate.SmsLog(sql.NotPredicates(p))
}
