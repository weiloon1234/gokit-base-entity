// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/weiloon1234/gokit-base-entity/ent/bank"
	"github.com/weiloon1234/gokit-base-entity/ent/country"
)

// Bank is the model entity for the Bank schema.
type Bank struct {
	config `json:"-"`
	// ID of the ent.
	// Primary key for the bank
	ID uint64 `json:"id,omitempty"`
	// Bank name in English
	NameEn string `json:"name_en,omitempty"`
	// Bank name in Chinese
	NameZh string `json:"name_zh,omitempty"`
	// Country of the bank
	CountryID *uint64 `json:"country_id,omitempty"`
	// Record creation timestamp
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Record update timestamp
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Record deleted timestamp
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BankQuery when eager-loading is set.
	Edges        BankEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BankEdges holds the relations/edges for other nodes in the graph.
type BankEdges struct {
	// Country to which the bank belongs
	Country *Country `json:"country,omitempty"`
	// User bind with this bank
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BankEdges) CountryOrErr() (*Country, error) {
	if e.Country != nil {
		return e.Country, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: country.Label}
	}
	return nil, &NotLoadedError{edge: "country"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e BankEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bank) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bank.FieldID, bank.FieldCountryID:
			values[i] = new(sql.NullInt64)
		case bank.FieldNameEn, bank.FieldNameZh:
			values[i] = new(sql.NullString)
		case bank.FieldCreatedAt, bank.FieldUpdatedAt, bank.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bank fields.
func (b *Bank) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bank.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = uint64(value.Int64)
		case bank.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				b.NameEn = value.String
			}
		case bank.FieldNameZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_zh", values[i])
			} else if value.Valid {
				b.NameZh = value.String
			}
		case bank.FieldCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				b.CountryID = new(uint64)
				*b.CountryID = uint64(value.Int64)
			}
		case bank.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bank.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case bank.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = new(time.Time)
				*b.DeletedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bank.
// This includes values selected through modifiers, order, etc.
func (b *Bank) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryCountry queries the "country" edge of the Bank entity.
func (b *Bank) QueryCountry() *CountryQuery {
	return NewBankClient(b.config).QueryCountry(b)
}

// QueryUsers queries the "users" edge of the Bank entity.
func (b *Bank) QueryUsers() *UserQuery {
	return NewBankClient(b.config).QueryUsers(b)
}

// Update returns a builder for updating this Bank.
// Note that you need to call Bank.Unwrap() before calling this method if this Bank
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bank) Update() *BankUpdateOne {
	return NewBankClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bank entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bank) Unwrap() *Bank {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bank is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bank) String() string {
	var builder strings.Builder
	builder.WriteString("Bank(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("name_en=")
	builder.WriteString(b.NameEn)
	builder.WriteString(", ")
	builder.WriteString("name_zh=")
	builder.WriteString(b.NameZh)
	builder.WriteString(", ")
	if v := b.CountryID; v != nil {
		builder.WriteString("country_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := b.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Banks is a parsable slice of Bank.
type Banks []*Bank
