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
	"github.com/weiloon1234/gokit-base-entity/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// Primary key for the user
	ID uint64 `json:"id,omitempty"`
	// Username of the user
	Username string `json:"username,omitempty"`
	// Full name of the user
	Name string `json:"name,omitempty"`
	// Email address of the user
	Email string `json:"email,omitempty"`
	// Timestamp when email was verified
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	// Password for the user
	Password string `json:"password,omitempty"`
	// Second password for the user
	Password2 string `json:"password2,omitempty"`
	// Country ID of the user
	CountryID *uint64 `json:"country_id,omitempty"`
	// Contact country ID of the user
	ContactCountryID *uint64 `json:"contact_country_id,omitempty"`
	// Contact number of the user
	ContactNumber *string `json:"contact_number,omitempty"`
	// Full contact number of the user
	FullContactNumber *string `json:"full_contact_number,omitempty"`
	// ID of the introducer user
	IntroducerUserID *uint64 `json:"introducer_user_id,omitempty"`
	// Preferred language of the user
	Lang string `json:"lang,omitempty"`
	// Avatar of the user
	Avatar *string `json:"avatar,omitempty"`
	// Credit 1 balance
	Credit1 float64 `json:"credit_1,omitempty"`
	// Credit 2 balance
	Credit2 float64 `json:"credit_2,omitempty"`
	// Credit 3 balance
	Credit3 float64 `json:"credit_3,omitempty"`
	// Credit 4 balance
	Credit4 float64 `json:"credit_4,omitempty"`
	// Credit 5 balance
	Credit5 float64 `json:"credit_5,omitempty"`
	// Bank ID of the user
	BankID *uint64 `json:"bank_id,omitempty"`
	// Bank account name of the user
	BankAccountName *string `json:"bank_account_name,omitempty"`
	// Bank account number of the user
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// National ID of the user
	NationalID *string `json:"national_id,omitempty"`
	// Whether this is the user's first login
	FirstLogin bool `json:"first_login,omitempty"`
	// Timestamp until the user is banned
	BanUntil *time.Time `json:"ban_until,omitempty"`
	// Timestamp of the user's latest login
	NewLoginAt *time.Time `json:"new_login_at,omitempty"`
	// Timestamp of the user's last login
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	// Unilevel of the user
	Unilevel *uint64 `json:"unilevel,omitempty"`
	// Record creation timestamp
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Record update timestamp
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Record deleted timestamp
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Country of the user
	Country *Country `json:"country,omitempty"`
	// Contact country of the user
	ContactCountry *Country `json:"contact_country,omitempty"`
	// Introducer user
	Introducer *User `json:"introducer,omitempty"`
	// Bank associated with the user
	Bank *Bank `json:"bank,omitempty"`
	// Users introduced by this user
	IntroducedUsers []*User `json:"introduced_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CountryOrErr() (*Country, error) {
	if e.Country != nil {
		return e.Country, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: country.Label}
	}
	return nil, &NotLoadedError{edge: "country"}
}

// ContactCountryOrErr returns the ContactCountry value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ContactCountryOrErr() (*Country, error) {
	if e.ContactCountry != nil {
		return e.ContactCountry, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: country.Label}
	}
	return nil, &NotLoadedError{edge: "contact_country"}
}

// IntroducerOrErr returns the Introducer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) IntroducerOrErr() (*User, error) {
	if e.Introducer != nil {
		return e.Introducer, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "introducer"}
}

// BankOrErr returns the Bank value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) BankOrErr() (*Bank, error) {
	if e.Bank != nil {
		return e.Bank, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: bank.Label}
	}
	return nil, &NotLoadedError{edge: "bank"}
}

// IntroducedUsersOrErr returns the IntroducedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IntroducedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.IntroducedUsers, nil
	}
	return nil, &NotLoadedError{edge: "introduced_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldFirstLogin:
			values[i] = new(sql.NullBool)
		case user.FieldCredit1, user.FieldCredit2, user.FieldCredit3, user.FieldCredit4, user.FieldCredit5:
			values[i] = new(sql.NullFloat64)
		case user.FieldID, user.FieldCountryID, user.FieldContactCountryID, user.FieldIntroducerUserID, user.FieldBankID, user.FieldUnilevel:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldName, user.FieldEmail, user.FieldPassword, user.FieldPassword2, user.FieldContactNumber, user.FieldFullContactNumber, user.FieldLang, user.FieldAvatar, user.FieldBankAccountName, user.FieldBankAccountNumber, user.FieldNationalID:
			values[i] = new(sql.NullString)
		case user.FieldEmailVerifiedAt, user.FieldBanUntil, user.FieldNewLoginAt, user.FieldLastLoginAt, user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint64(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldEmailVerifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified_at", values[i])
			} else if value.Valid {
				u.EmailVerifiedAt = new(time.Time)
				*u.EmailVerifiedAt = value.Time
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldPassword2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password2", values[i])
			} else if value.Valid {
				u.Password2 = value.String
			}
		case user.FieldCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				u.CountryID = new(uint64)
				*u.CountryID = uint64(value.Int64)
			}
		case user.FieldContactCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contact_country_id", values[i])
			} else if value.Valid {
				u.ContactCountryID = new(uint64)
				*u.ContactCountryID = uint64(value.Int64)
			}
		case user.FieldContactNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_number", values[i])
			} else if value.Valid {
				u.ContactNumber = new(string)
				*u.ContactNumber = value.String
			}
		case user.FieldFullContactNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_contact_number", values[i])
			} else if value.Valid {
				u.FullContactNumber = new(string)
				*u.FullContactNumber = value.String
			}
		case user.FieldIntroducerUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field introducer_user_id", values[i])
			} else if value.Valid {
				u.IntroducerUserID = new(uint64)
				*u.IntroducerUserID = uint64(value.Int64)
			}
		case user.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				u.Lang = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = new(string)
				*u.Avatar = value.String
			}
		case user.FieldCredit1:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_1", values[i])
			} else if value.Valid {
				u.Credit1 = value.Float64
			}
		case user.FieldCredit2:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_2", values[i])
			} else if value.Valid {
				u.Credit2 = value.Float64
			}
		case user.FieldCredit3:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_3", values[i])
			} else if value.Valid {
				u.Credit3 = value.Float64
			}
		case user.FieldCredit4:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_4", values[i])
			} else if value.Valid {
				u.Credit4 = value.Float64
			}
		case user.FieldCredit5:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_5", values[i])
			} else if value.Valid {
				u.Credit5 = value.Float64
			}
		case user.FieldBankID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bank_id", values[i])
			} else if value.Valid {
				u.BankID = new(uint64)
				*u.BankID = uint64(value.Int64)
			}
		case user.FieldBankAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account_name", values[i])
			} else if value.Valid {
				u.BankAccountName = new(string)
				*u.BankAccountName = value.String
			}
		case user.FieldBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account_number", values[i])
			} else if value.Valid {
				u.BankAccountNumber = new(string)
				*u.BankAccountNumber = value.String
			}
		case user.FieldNationalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field national_id", values[i])
			} else if value.Valid {
				u.NationalID = new(string)
				*u.NationalID = value.String
			}
		case user.FieldFirstLogin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field first_login", values[i])
			} else if value.Valid {
				u.FirstLogin = value.Bool
			}
		case user.FieldBanUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ban_until", values[i])
			} else if value.Valid {
				u.BanUntil = new(time.Time)
				*u.BanUntil = value.Time
			}
		case user.FieldNewLoginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field new_login_at", values[i])
			} else if value.Valid {
				u.NewLoginAt = new(time.Time)
				*u.NewLoginAt = value.Time
			}
		case user.FieldLastLoginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_at", values[i])
			} else if value.Valid {
				u.LastLoginAt = new(time.Time)
				*u.LastLoginAt = value.Time
			}
		case user.FieldUnilevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unilevel", values[i])
			} else if value.Valid {
				u.Unilevel = new(uint64)
				*u.Unilevel = uint64(value.Int64)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryCountry queries the "country" edge of the User entity.
func (u *User) QueryCountry() *CountryQuery {
	return NewUserClient(u.config).QueryCountry(u)
}

// QueryContactCountry queries the "contact_country" edge of the User entity.
func (u *User) QueryContactCountry() *CountryQuery {
	return NewUserClient(u.config).QueryContactCountry(u)
}

// QueryIntroducer queries the "introducer" edge of the User entity.
func (u *User) QueryIntroducer() *UserQuery {
	return NewUserClient(u.config).QueryIntroducer(u)
}

// QueryBank queries the "bank" edge of the User entity.
func (u *User) QueryBank() *BankQuery {
	return NewUserClient(u.config).QueryBank(u)
}

// QueryIntroducedUsers queries the "introduced_users" edge of the User entity.
func (u *User) QueryIntroducedUsers() *UserQuery {
	return NewUserClient(u.config).QueryIntroducedUsers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	if v := u.EmailVerifiedAt; v != nil {
		builder.WriteString("email_verified_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("password2=")
	builder.WriteString(u.Password2)
	builder.WriteString(", ")
	if v := u.CountryID; v != nil {
		builder.WriteString("country_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.ContactCountryID; v != nil {
		builder.WriteString("contact_country_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.ContactNumber; v != nil {
		builder.WriteString("contact_number=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.FullContactNumber; v != nil {
		builder.WriteString("full_contact_number=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.IntroducerUserID; v != nil {
		builder.WriteString("introducer_user_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(u.Lang)
	builder.WriteString(", ")
	if v := u.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("credit_1=")
	builder.WriteString(fmt.Sprintf("%v", u.Credit1))
	builder.WriteString(", ")
	builder.WriteString("credit_2=")
	builder.WriteString(fmt.Sprintf("%v", u.Credit2))
	builder.WriteString(", ")
	builder.WriteString("credit_3=")
	builder.WriteString(fmt.Sprintf("%v", u.Credit3))
	builder.WriteString(", ")
	builder.WriteString("credit_4=")
	builder.WriteString(fmt.Sprintf("%v", u.Credit4))
	builder.WriteString(", ")
	builder.WriteString("credit_5=")
	builder.WriteString(fmt.Sprintf("%v", u.Credit5))
	builder.WriteString(", ")
	if v := u.BankID; v != nil {
		builder.WriteString("bank_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.BankAccountName; v != nil {
		builder.WriteString("bank_account_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.BankAccountNumber; v != nil {
		builder.WriteString("bank_account_number=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.NationalID; v != nil {
		builder.WriteString("national_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("first_login=")
	builder.WriteString(fmt.Sprintf("%v", u.FirstLogin))
	builder.WriteString(", ")
	if v := u.BanUntil; v != nil {
		builder.WriteString("ban_until=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.NewLoginAt; v != nil {
		builder.WriteString("new_login_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.LastLoginAt; v != nil {
		builder.WriteString("last_login_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.Unilevel; v != nil {
		builder.WriteString("unilevel=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
