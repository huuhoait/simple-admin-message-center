// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-message-center/ent/smsprovider"
)

// SmsProvider is the model entity for the SmsProvider schema.
type SmsProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The SMS provider name | 短信服务的提供商
	Name string `json:"name,omitempty"`
	// The secret ID | 密钥 ID
	SecretID string `json:"secret_id,omitempty"`
	// The secret key | 密钥 Key
	SecretKey string `json:"secret_key,omitempty"`
	// The service region | 服务器所在地区
	Region string `json:"region,omitempty"`
	// Is it the default provider | 是否为默认提供商
	IsDefault    bool `json:"is_default,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SmsProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case smsprovider.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case smsprovider.FieldID:
			values[i] = new(sql.NullInt64)
		case smsprovider.FieldName, smsprovider.FieldSecretID, smsprovider.FieldSecretKey, smsprovider.FieldRegion:
			values[i] = new(sql.NullString)
		case smsprovider.FieldCreatedAt, smsprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SmsProvider fields.
func (sp *SmsProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case smsprovider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = uint64(value.Int64)
		case smsprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case smsprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		case smsprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case smsprovider.FieldSecretID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_id", values[i])
			} else if value.Valid {
				sp.SecretID = value.String
			}
		case smsprovider.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				sp.SecretKey = value.String
			}
		case smsprovider.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				sp.Region = value.String
			}
		case smsprovider.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				sp.IsDefault = value.Bool
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SmsProvider.
// This includes values selected through modifiers, order, etc.
func (sp *SmsProvider) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// Update returns a builder for updating this SmsProvider.
// Note that you need to call SmsProvider.Unwrap() before calling this method if this SmsProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *SmsProvider) Update() *SmsProviderUpdateOne {
	return NewSmsProviderClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the SmsProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *SmsProvider) Unwrap() *SmsProvider {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SmsProvider is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *SmsProvider) String() string {
	var builder strings.Builder
	builder.WriteString("SmsProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", ")
	builder.WriteString("secret_id=")
	builder.WriteString(sp.SecretID)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(sp.SecretKey)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(sp.Region)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsDefault))
	builder.WriteByte(')')
	return builder.String()
}

// SmsProviders is a parsable slice of SmsProvider.
type SmsProviders []*SmsProvider