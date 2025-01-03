// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldEmailVerifiedAt holds the string denoting the email_verified_at field in the database.
	FieldEmailVerifiedAt = "email_verified_at"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPassword2 holds the string denoting the password2 field in the database.
	FieldPassword2 = "password2"
	// FieldCountryID holds the string denoting the country_id field in the database.
	FieldCountryID = "country_id"
	// FieldContactCountryID holds the string denoting the contact_country_id field in the database.
	FieldContactCountryID = "contact_country_id"
	// FieldContactNumber holds the string denoting the contact_number field in the database.
	FieldContactNumber = "contact_number"
	// FieldFullContactNumber holds the string denoting the full_contact_number field in the database.
	FieldFullContactNumber = "full_contact_number"
	// FieldIntroducerUserID holds the string denoting the introducer_user_id field in the database.
	FieldIntroducerUserID = "introducer_user_id"
	// FieldLang holds the string denoting the lang field in the database.
	FieldLang = "lang"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldCredit1 holds the string denoting the credit_1 field in the database.
	FieldCredit1 = "credit_1"
	// FieldCredit2 holds the string denoting the credit_2 field in the database.
	FieldCredit2 = "credit_2"
	// FieldCredit3 holds the string denoting the credit_3 field in the database.
	FieldCredit3 = "credit_3"
	// FieldCredit4 holds the string denoting the credit_4 field in the database.
	FieldCredit4 = "credit_4"
	// FieldCredit5 holds the string denoting the credit_5 field in the database.
	FieldCredit5 = "credit_5"
	// FieldBankID holds the string denoting the bank_id field in the database.
	FieldBankID = "bank_id"
	// FieldBankAccountName holds the string denoting the bank_account_name field in the database.
	FieldBankAccountName = "bank_account_name"
	// FieldBankAccountNumber holds the string denoting the bank_account_number field in the database.
	FieldBankAccountNumber = "bank_account_number"
	// FieldNationalID holds the string denoting the national_id field in the database.
	FieldNationalID = "national_id"
	// FieldFirstLogin holds the string denoting the first_login field in the database.
	FieldFirstLogin = "first_login"
	// FieldBanUntil holds the string denoting the ban_until field in the database.
	FieldBanUntil = "ban_until"
	// FieldNewLoginAt holds the string denoting the new_login_at field in the database.
	FieldNewLoginAt = "new_login_at"
	// FieldLastLoginAt holds the string denoting the last_login_at field in the database.
	FieldLastLoginAt = "last_login_at"
	// FieldUnilevel holds the string denoting the unilevel field in the database.
	FieldUnilevel = "unilevel"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeCountry holds the string denoting the country edge name in mutations.
	EdgeCountry = "country"
	// EdgeContactCountry holds the string denoting the contact_country edge name in mutations.
	EdgeContactCountry = "contact_country"
	// EdgeIntroducer holds the string denoting the introducer edge name in mutations.
	EdgeIntroducer = "introducer"
	// EdgeBank holds the string denoting the bank edge name in mutations.
	EdgeBank = "bank"
	// EdgeIntroducedUsers holds the string denoting the introduced_users edge name in mutations.
	EdgeIntroducedUsers = "introduced_users"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CountryTable is the table that holds the country relation/edge.
	CountryTable = "users"
	// CountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountryInverseTable = "countries"
	// CountryColumn is the table column denoting the country relation/edge.
	CountryColumn = "country_id"
	// ContactCountryTable is the table that holds the contact_country relation/edge.
	ContactCountryTable = "users"
	// ContactCountryInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	ContactCountryInverseTable = "countries"
	// ContactCountryColumn is the table column denoting the contact_country relation/edge.
	ContactCountryColumn = "contact_country_id"
	// IntroducerTable is the table that holds the introducer relation/edge.
	IntroducerTable = "users"
	// IntroducerColumn is the table column denoting the introducer relation/edge.
	IntroducerColumn = "introducer_user_id"
	// BankTable is the table that holds the bank relation/edge.
	BankTable = "users"
	// BankInverseTable is the table name for the Bank entity.
	// It exists in this package in order to avoid circular dependency with the "bank" package.
	BankInverseTable = "banks"
	// BankColumn is the table column denoting the bank relation/edge.
	BankColumn = "bank_id"
	// IntroducedUsersTable is the table that holds the introduced_users relation/edge.
	IntroducedUsersTable = "users"
	// IntroducedUsersColumn is the table column denoting the introduced_users relation/edge.
	IntroducedUsersColumn = "introducer_user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldName,
	FieldEmail,
	FieldEmailVerifiedAt,
	FieldPassword,
	FieldPassword2,
	FieldCountryID,
	FieldContactCountryID,
	FieldContactNumber,
	FieldFullContactNumber,
	FieldIntroducerUserID,
	FieldLang,
	FieldAvatar,
	FieldCredit1,
	FieldCredit2,
	FieldCredit3,
	FieldCredit4,
	FieldCredit5,
	FieldBankID,
	FieldBankAccountName,
	FieldBankAccountNumber,
	FieldNationalID,
	FieldFirstLogin,
	FieldBanUntil,
	FieldNewLoginAt,
	FieldLastLoginAt,
	FieldUnilevel,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// Password2Validator is a validator for the "password2" field. It is called by the builders before save.
	Password2Validator func(string) error
	// DefaultLang holds the default value on creation for the "lang" field.
	DefaultLang string
	// DefaultCredit1 holds the default value on creation for the "credit_1" field.
	DefaultCredit1 float64
	// DefaultCredit2 holds the default value on creation for the "credit_2" field.
	DefaultCredit2 float64
	// DefaultCredit3 holds the default value on creation for the "credit_3" field.
	DefaultCredit3 float64
	// DefaultCredit4 holds the default value on creation for the "credit_4" field.
	DefaultCredit4 float64
	// DefaultCredit5 holds the default value on creation for the "credit_5" field.
	DefaultCredit5 float64
	// DefaultFirstLogin holds the default value on creation for the "first_login" field.
	DefaultFirstLogin bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByEmailVerifiedAt orders the results by the email_verified_at field.
func ByEmailVerifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailVerifiedAt, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByPassword2 orders the results by the password2 field.
func ByPassword2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword2, opts...).ToFunc()
}

// ByCountryID orders the results by the country_id field.
func ByCountryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryID, opts...).ToFunc()
}

// ByContactCountryID orders the results by the contact_country_id field.
func ByContactCountryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactCountryID, opts...).ToFunc()
}

// ByContactNumber orders the results by the contact_number field.
func ByContactNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactNumber, opts...).ToFunc()
}

// ByFullContactNumber orders the results by the full_contact_number field.
func ByFullContactNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullContactNumber, opts...).ToFunc()
}

// ByIntroducerUserID orders the results by the introducer_user_id field.
func ByIntroducerUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntroducerUserID, opts...).ToFunc()
}

// ByLang orders the results by the lang field.
func ByLang(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLang, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByCredit1 orders the results by the credit_1 field.
func ByCredit1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredit1, opts...).ToFunc()
}

// ByCredit2 orders the results by the credit_2 field.
func ByCredit2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredit2, opts...).ToFunc()
}

// ByCredit3 orders the results by the credit_3 field.
func ByCredit3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredit3, opts...).ToFunc()
}

// ByCredit4 orders the results by the credit_4 field.
func ByCredit4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredit4, opts...).ToFunc()
}

// ByCredit5 orders the results by the credit_5 field.
func ByCredit5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredit5, opts...).ToFunc()
}

// ByBankID orders the results by the bank_id field.
func ByBankID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankID, opts...).ToFunc()
}

// ByBankAccountName orders the results by the bank_account_name field.
func ByBankAccountName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankAccountName, opts...).ToFunc()
}

// ByBankAccountNumber orders the results by the bank_account_number field.
func ByBankAccountNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankAccountNumber, opts...).ToFunc()
}

// ByNationalID orders the results by the national_id field.
func ByNationalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNationalID, opts...).ToFunc()
}

// ByFirstLogin orders the results by the first_login field.
func ByFirstLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstLogin, opts...).ToFunc()
}

// ByBanUntil orders the results by the ban_until field.
func ByBanUntil(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBanUntil, opts...).ToFunc()
}

// ByNewLoginAt orders the results by the new_login_at field.
func ByNewLoginAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNewLoginAt, opts...).ToFunc()
}

// ByLastLoginAt orders the results by the last_login_at field.
func ByLastLoginAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginAt, opts...).ToFunc()
}

// ByUnilevel orders the results by the unilevel field.
func ByUnilevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnilevel, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByCountryField orders the results by country field.
func ByCountryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountryStep(), sql.OrderByField(field, opts...))
	}
}

// ByContactCountryField orders the results by contact_country field.
func ByContactCountryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContactCountryStep(), sql.OrderByField(field, opts...))
	}
}

// ByIntroducerField orders the results by introducer field.
func ByIntroducerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIntroducerStep(), sql.OrderByField(field, opts...))
	}
}

// ByBankField orders the results by bank field.
func ByBankField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBankStep(), sql.OrderByField(field, opts...))
	}
}

// ByIntroducedUsersCount orders the results by introduced_users count.
func ByIntroducedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIntroducedUsersStep(), opts...)
	}
}

// ByIntroducedUsers orders the results by introduced_users terms.
func ByIntroducedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIntroducedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCountryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CountryTable, CountryColumn),
	)
}
func newContactCountryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContactCountryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ContactCountryTable, ContactCountryColumn),
	)
}
func newIntroducerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, IntroducerTable, IntroducerColumn),
	)
}
func newBankStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BankInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BankTable, BankColumn),
	)
}
func newIntroducedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IntroducedUsersTable, IntroducedUsersColumn),
	)
}
