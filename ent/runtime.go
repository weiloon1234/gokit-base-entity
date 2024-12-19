// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/weiloon1234/gokit-base-entity/ent/admin"
	"github.com/weiloon1234/gokit-base-entity/ent/bank"
	"github.com/weiloon1234/gokit-base-entity/ent/country"
	"github.com/weiloon1234/gokit-base-entity/ent/countrylocation"
	"github.com/weiloon1234/gokit-base-entity/ent/schema"
	"github.com/weiloon1234/gokit-base-entity/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescPassword is the schema descriptor for password field.
	adminDescPassword := adminFields[5].Descriptor()
	// admin.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	admin.PasswordValidator = adminDescPassword.Validators[0].(func(string) error)
	// adminDescPassword2 is the schema descriptor for password2 field.
	adminDescPassword2 := adminFields[6].Descriptor()
	// admin.Password2Validator is a validator for the "password2" field. It is called by the builders before save.
	admin.Password2Validator = adminDescPassword2.Validators[0].(func(string) error)
	// adminDescLang is the schema descriptor for lang field.
	adminDescLang := adminFields[7].Descriptor()
	// admin.DefaultLang holds the default value on creation for the lang field.
	admin.DefaultLang = adminDescLang.Default.(string)
	// adminDescType is the schema descriptor for type field.
	adminDescType := adminFields[9].Descriptor()
	// admin.DefaultType holds the default value on creation for the type field.
	admin.DefaultType = adminDescType.Default.(uint8)
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminFields[10].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminFields[11].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	bankFields := schema.Bank{}.Fields()
	_ = bankFields
	// bankDescNameEn is the schema descriptor for name_en field.
	bankDescNameEn := bankFields[1].Descriptor()
	// bank.NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	bank.NameEnValidator = bankDescNameEn.Validators[0].(func(string) error)
	// bankDescNameZh is the schema descriptor for name_zh field.
	bankDescNameZh := bankFields[2].Descriptor()
	// bank.NameZhValidator is a validator for the "name_zh" field. It is called by the builders before save.
	bank.NameZhValidator = bankDescNameZh.Validators[0].(func(string) error)
	// bankDescCreatedAt is the schema descriptor for created_at field.
	bankDescCreatedAt := bankFields[4].Descriptor()
	// bank.DefaultCreatedAt holds the default value on creation for the created_at field.
	bank.DefaultCreatedAt = bankDescCreatedAt.Default.(func() time.Time)
	// bankDescUpdatedAt is the schema descriptor for updated_at field.
	bankDescUpdatedAt := bankFields[5].Descriptor()
	// bank.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bank.DefaultUpdatedAt = bankDescUpdatedAt.Default.(func() time.Time)
	// bank.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bank.UpdateDefaultUpdatedAt = bankDescUpdatedAt.UpdateDefault.(func() time.Time)
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescIso2 is the schema descriptor for iso2 field.
	countryDescIso2 := countryFields[1].Descriptor()
	// country.Iso2Validator is a validator for the "iso2" field. It is called by the builders before save.
	country.Iso2Validator = countryDescIso2.Validators[0].(func(string) error)
	// countryDescIso3 is the schema descriptor for iso3 field.
	countryDescIso3 := countryFields[2].Descriptor()
	// country.Iso3Validator is a validator for the "iso3" field. It is called by the builders before save.
	country.Iso3Validator = countryDescIso3.Validators[0].(func(string) error)
	// countryDescName is the schema descriptor for name field.
	countryDescName := countryFields[3].Descriptor()
	// country.NameValidator is a validator for the "name" field. It is called by the builders before save.
	country.NameValidator = countryDescName.Validators[0].(func(string) error)
	// countryDescConversionRate is the schema descriptor for conversion_rate field.
	countryDescConversionRate := countryFields[11].Descriptor()
	// country.DefaultConversionRate holds the default value on creation for the conversion_rate field.
	country.DefaultConversionRate = countryDescConversionRate.Default.(float64)
	// countryDescStatus is the schema descriptor for status field.
	countryDescStatus := countryFields[12].Descriptor()
	// country.DefaultStatus holds the default value on creation for the status field.
	country.DefaultStatus = countryDescStatus.Default.(uint8)
	// countryDescCreatedAt is the schema descriptor for created_at field.
	countryDescCreatedAt := countryFields[13].Descriptor()
	// country.DefaultCreatedAt holds the default value on creation for the created_at field.
	country.DefaultCreatedAt = countryDescCreatedAt.Default.(func() time.Time)
	// countryDescUpdatedAt is the schema descriptor for updated_at field.
	countryDescUpdatedAt := countryFields[14].Descriptor()
	// country.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	country.DefaultUpdatedAt = countryDescUpdatedAt.Default.(func() time.Time)
	// country.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	country.UpdateDefaultUpdatedAt = countryDescUpdatedAt.UpdateDefault.(func() time.Time)
	countrylocationFields := schema.CountryLocation{}.Fields()
	_ = countrylocationFields
	// countrylocationDescSorting is the schema descriptor for sorting field.
	countrylocationDescSorting := countrylocationFields[3].Descriptor()
	// countrylocation.DefaultSorting holds the default value on creation for the sorting field.
	countrylocation.DefaultSorting = countrylocationDescSorting.Default.(uint64)
	// countrylocationDescNameEn is the schema descriptor for name_en field.
	countrylocationDescNameEn := countrylocationFields[4].Descriptor()
	// countrylocation.NameEnValidator is a validator for the "name_en" field. It is called by the builders before save.
	countrylocation.NameEnValidator = countrylocationDescNameEn.Validators[0].(func(string) error)
	// countrylocationDescNameZh is the schema descriptor for name_zh field.
	countrylocationDescNameZh := countrylocationFields[5].Descriptor()
	// countrylocation.NameZhValidator is a validator for the "name_zh" field. It is called by the builders before save.
	countrylocation.NameZhValidator = countrylocationDescNameZh.Validators[0].(func(string) error)
	// countrylocationDescCreatedAt is the schema descriptor for created_at field.
	countrylocationDescCreatedAt := countrylocationFields[6].Descriptor()
	// countrylocation.DefaultCreatedAt holds the default value on creation for the created_at field.
	countrylocation.DefaultCreatedAt = countrylocationDescCreatedAt.Default.(func() time.Time)
	// countrylocationDescUpdatedAt is the schema descriptor for updated_at field.
	countrylocationDescUpdatedAt := countrylocationFields[7].Descriptor()
	// countrylocation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	countrylocation.DefaultUpdatedAt = countrylocationDescUpdatedAt.Default.(time.Time)
	// countrylocation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	countrylocation.UpdateDefaultUpdatedAt = countrylocationDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[5].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescPassword2 is the schema descriptor for password2 field.
	userDescPassword2 := userFields[6].Descriptor()
	// user.Password2Validator is a validator for the "password2" field. It is called by the builders before save.
	user.Password2Validator = userDescPassword2.Validators[0].(func(string) error)
	// userDescLang is the schema descriptor for lang field.
	userDescLang := userFields[12].Descriptor()
	// user.DefaultLang holds the default value on creation for the lang field.
	user.DefaultLang = userDescLang.Default.(string)
	// userDescCredit1 is the schema descriptor for credit_1 field.
	userDescCredit1 := userFields[14].Descriptor()
	// user.DefaultCredit1 holds the default value on creation for the credit_1 field.
	user.DefaultCredit1 = userDescCredit1.Default.(float64)
	// userDescCredit2 is the schema descriptor for credit_2 field.
	userDescCredit2 := userFields[15].Descriptor()
	// user.DefaultCredit2 holds the default value on creation for the credit_2 field.
	user.DefaultCredit2 = userDescCredit2.Default.(float64)
	// userDescCredit3 is the schema descriptor for credit_3 field.
	userDescCredit3 := userFields[16].Descriptor()
	// user.DefaultCredit3 holds the default value on creation for the credit_3 field.
	user.DefaultCredit3 = userDescCredit3.Default.(float64)
	// userDescCredit4 is the schema descriptor for credit_4 field.
	userDescCredit4 := userFields[17].Descriptor()
	// user.DefaultCredit4 holds the default value on creation for the credit_4 field.
	user.DefaultCredit4 = userDescCredit4.Default.(float64)
	// userDescCredit5 is the schema descriptor for credit_5 field.
	userDescCredit5 := userFields[18].Descriptor()
	// user.DefaultCredit5 holds the default value on creation for the credit_5 field.
	user.DefaultCredit5 = userDescCredit5.Default.(float64)
	// userDescFirstLogin is the schema descriptor for first_login field.
	userDescFirstLogin := userFields[23].Descriptor()
	// user.DefaultFirstLogin holds the default value on creation for the first_login field.
	user.DefaultFirstLogin = userDescFirstLogin.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[28].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[29].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
