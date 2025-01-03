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
	"github.com/weiloon1234/gokit-base-entity/ent/admin"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsername sets the "username" field.
func (ac *AdminCreate) SetUsername(s string) *AdminCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUsername(s *string) *AdminCreate {
	if s != nil {
		ac.SetUsername(*s)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AdminCreate) SetName(s string) *AdminCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *AdminCreate) SetNillableName(s *string) *AdminCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetEmail sets the "email" field.
func (ac *AdminCreate) SetEmail(s string) *AdminCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ac *AdminCreate) SetNillableEmail(s *string) *AdminCreate {
	if s != nil {
		ac.SetEmail(*s)
	}
	return ac
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (ac *AdminCreate) SetEmailVerifiedAt(t time.Time) *AdminCreate {
	ac.mutation.SetEmailVerifiedAt(t)
	return ac
}

// SetNillableEmailVerifiedAt sets the "email_verified_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableEmailVerifiedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetEmailVerifiedAt(*t)
	}
	return ac
}

// SetPassword sets the "password" field.
func (ac *AdminCreate) SetPassword(s string) *AdminCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetPassword2 sets the "password2" field.
func (ac *AdminCreate) SetPassword2(s string) *AdminCreate {
	ac.mutation.SetPassword2(s)
	return ac
}

// SetLang sets the "lang" field.
func (ac *AdminCreate) SetLang(s string) *AdminCreate {
	ac.mutation.SetLang(s)
	return ac
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (ac *AdminCreate) SetNillableLang(s *string) *AdminCreate {
	if s != nil {
		ac.SetLang(*s)
	}
	return ac
}

// SetAvatar sets the "avatar" field.
func (ac *AdminCreate) SetAvatar(s string) *AdminCreate {
	ac.mutation.SetAvatar(s)
	return ac
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (ac *AdminCreate) SetNillableAvatar(s *string) *AdminCreate {
	if s != nil {
		ac.SetAvatar(*s)
	}
	return ac
}

// SetType sets the "type" field.
func (ac *AdminCreate) SetType(u uint8) *AdminCreate {
	ac.mutation.SetType(u)
	return ac
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ac *AdminCreate) SetNillableType(u *uint8) *AdminCreate {
	if u != nil {
		ac.SetType(*u)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminCreate) SetCreatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminCreate) SetUpdatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AdminCreate) SetDeletedAt(t time.Time) *AdminCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableDeletedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AdminCreate) SetID(u uint64) *AdminCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.Lang(); !ok {
		v := admin.DefaultLang
		ac.mutation.SetLang(v)
	}
	if _, ok := ac.mutation.GetType(); !ok {
		v := admin.DefaultType
		ac.mutation.SetType(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admin.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admin.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Admin.password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Admin.password": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Password2(); !ok {
		return &ValidationError{Name: "password2", err: errors.New(`ent: missing required field "Admin.password2"`)}
	}
	if v, ok := ac.mutation.Password2(); ok {
		if err := admin.Password2Validator(v); err != nil {
			return &ValidationError{Name: "password2", err: fmt.Errorf(`ent: validator failed for field "Admin.password2": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "Admin.lang"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Admin.type"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Admin.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Admin.updated_at"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(admin.Table, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(admin.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(admin.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.EmailVerifiedAt(); ok {
		_spec.SetField(admin.FieldEmailVerifiedAt, field.TypeTime, value)
		_node.EmailVerifiedAt = &value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.Password2(); ok {
		_spec.SetField(admin.FieldPassword2, field.TypeString, value)
		_node.Password2 = value
	}
	if value, ok := ac.mutation.Lang(); ok {
		_spec.SetField(admin.FieldLang, field.TypeString, value)
		_node.Lang = value
	}
	if value, ok := ac.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
		_node.Avatar = &value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(admin.FieldType, field.TypeUint8, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.Create().
//		SetUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (ac *AdminCreate) OnConflict(opts ...sql.ConflictOption) *AdminUpsertOne {
	ac.conflict = opts
	return &AdminUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AdminCreate) OnConflictColumns(columns ...string) *AdminUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertOne{
		create: ac,
	}
}

type (
	// AdminUpsertOne is the builder for "upsert"-ing
	//  one Admin node.
	AdminUpsertOne struct {
		create *AdminCreate
	}

	// AdminUpsert is the "OnConflict" setter.
	AdminUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsername sets the "username" field.
func (u *AdminUpsert) SetUsername(v string) *AdminUpsert {
	u.Set(admin.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUsername() *AdminUpsert {
	u.SetExcluded(admin.FieldUsername)
	return u
}

// ClearUsername clears the value of the "username" field.
func (u *AdminUpsert) ClearUsername() *AdminUpsert {
	u.SetNull(admin.FieldUsername)
	return u
}

// SetName sets the "name" field.
func (u *AdminUpsert) SetName(v string) *AdminUpsert {
	u.Set(admin.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AdminUpsert) UpdateName() *AdminUpsert {
	u.SetExcluded(admin.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *AdminUpsert) ClearName() *AdminUpsert {
	u.SetNull(admin.FieldName)
	return u
}

// SetEmail sets the "email" field.
func (u *AdminUpsert) SetEmail(v string) *AdminUpsert {
	u.Set(admin.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AdminUpsert) UpdateEmail() *AdminUpsert {
	u.SetExcluded(admin.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *AdminUpsert) ClearEmail() *AdminUpsert {
	u.SetNull(admin.FieldEmail)
	return u
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (u *AdminUpsert) SetEmailVerifiedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldEmailVerifiedAt, v)
	return u
}

// UpdateEmailVerifiedAt sets the "email_verified_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateEmailVerifiedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldEmailVerifiedAt)
	return u
}

// ClearEmailVerifiedAt clears the value of the "email_verified_at" field.
func (u *AdminUpsert) ClearEmailVerifiedAt() *AdminUpsert {
	u.SetNull(admin.FieldEmailVerifiedAt)
	return u
}

// SetPassword sets the "password" field.
func (u *AdminUpsert) SetPassword(v string) *AdminUpsert {
	u.Set(admin.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsert) UpdatePassword() *AdminUpsert {
	u.SetExcluded(admin.FieldPassword)
	return u
}

// SetPassword2 sets the "password2" field.
func (u *AdminUpsert) SetPassword2(v string) *AdminUpsert {
	u.Set(admin.FieldPassword2, v)
	return u
}

// UpdatePassword2 sets the "password2" field to the value that was provided on create.
func (u *AdminUpsert) UpdatePassword2() *AdminUpsert {
	u.SetExcluded(admin.FieldPassword2)
	return u
}

// SetLang sets the "lang" field.
func (u *AdminUpsert) SetLang(v string) *AdminUpsert {
	u.Set(admin.FieldLang, v)
	return u
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *AdminUpsert) UpdateLang() *AdminUpsert {
	u.SetExcluded(admin.FieldLang)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *AdminUpsert) SetAvatar(v string) *AdminUpsert {
	u.Set(admin.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AdminUpsert) UpdateAvatar() *AdminUpsert {
	u.SetExcluded(admin.FieldAvatar)
	return u
}

// ClearAvatar clears the value of the "avatar" field.
func (u *AdminUpsert) ClearAvatar() *AdminUpsert {
	u.SetNull(admin.FieldAvatar)
	return u
}

// SetType sets the "type" field.
func (u *AdminUpsert) SetType(v uint8) *AdminUpsert {
	u.Set(admin.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AdminUpsert) UpdateType() *AdminUpsert {
	u.SetExcluded(admin.FieldType)
	return u
}

// AddType adds v to the "type" field.
func (u *AdminUpsert) AddType(v uint8) *AdminUpsert {
	u.Add(admin.FieldType, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsert) SetCreatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateCreatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsert) SetUpdatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUpdatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsert) SetDeletedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateDeletedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AdminUpsert) ClearDeletedAt() *AdminUpsert {
	u.SetNull(admin.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(admin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdminUpsertOne) UpdateNewValues() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(admin.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AdminUpsertOne) Ignore() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertOne) DoNothing() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreate.OnConflict
// documentation for more info.
func (u *AdminUpsertOne) Update(set func(*AdminUpsert)) *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *AdminUpsertOne) SetUsername(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUsername() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *AdminUpsertOne) ClearUsername() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearUsername()
	})
}

// SetName sets the "name" field.
func (u *AdminUpsertOne) SetName(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateName() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AdminUpsertOne) ClearName() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearName()
	})
}

// SetEmail sets the "email" field.
func (u *AdminUpsertOne) SetEmail(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateEmail() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *AdminUpsertOne) ClearEmail() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearEmail()
	})
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (u *AdminUpsertOne) SetEmailVerifiedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetEmailVerifiedAt(v)
	})
}

// UpdateEmailVerifiedAt sets the "email_verified_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateEmailVerifiedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateEmailVerifiedAt()
	})
}

// ClearEmailVerifiedAt clears the value of the "email_verified_at" field.
func (u *AdminUpsertOne) ClearEmailVerifiedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearEmailVerifiedAt()
	})
}

// SetPassword sets the "password" field.
func (u *AdminUpsertOne) SetPassword(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdatePassword() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword()
	})
}

// SetPassword2 sets the "password2" field.
func (u *AdminUpsertOne) SetPassword2(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword2(v)
	})
}

// UpdatePassword2 sets the "password2" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdatePassword2() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword2()
	})
}

// SetLang sets the "lang" field.
func (u *AdminUpsertOne) SetLang(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateLang() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateLang()
	})
}

// SetAvatar sets the "avatar" field.
func (u *AdminUpsertOne) SetAvatar(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateAvatar() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateAvatar()
	})
}

// ClearAvatar clears the value of the "avatar" field.
func (u *AdminUpsertOne) ClearAvatar() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearAvatar()
	})
}

// SetType sets the "type" field.
func (u *AdminUpsertOne) SetType(v uint8) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *AdminUpsertOne) AddType(v uint8) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateType() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsertOne) SetCreatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateCreatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertOne) SetUpdatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUpdatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsertOne) SetDeletedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateDeletedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AdminUpsertOne) ClearDeletedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AdminUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AdminCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AdminUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AdminUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	err      error
	builders []*AdminCreate
	conflict []sql.ConflictOption
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (acb *AdminCreateBulk) OnConflict(opts ...sql.ConflictOption) *AdminUpsertBulk {
	acb.conflict = opts
	return &AdminUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AdminCreateBulk) OnConflictColumns(columns ...string) *AdminUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertBulk{
		create: acb,
	}
}

// AdminUpsertBulk is the builder for "upsert"-ing
// a bulk of Admin nodes.
type AdminUpsertBulk struct {
	create *AdminCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(admin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdminUpsertBulk) UpdateNewValues() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(admin.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AdminUpsertBulk) Ignore() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertBulk) DoNothing() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreateBulk.OnConflict
// documentation for more info.
func (u *AdminUpsertBulk) Update(set func(*AdminUpsert)) *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *AdminUpsertBulk) SetUsername(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUsername() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *AdminUpsertBulk) ClearUsername() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearUsername()
	})
}

// SetName sets the "name" field.
func (u *AdminUpsertBulk) SetName(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateName() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AdminUpsertBulk) ClearName() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearName()
	})
}

// SetEmail sets the "email" field.
func (u *AdminUpsertBulk) SetEmail(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateEmail() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *AdminUpsertBulk) ClearEmail() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearEmail()
	})
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (u *AdminUpsertBulk) SetEmailVerifiedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetEmailVerifiedAt(v)
	})
}

// UpdateEmailVerifiedAt sets the "email_verified_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateEmailVerifiedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateEmailVerifiedAt()
	})
}

// ClearEmailVerifiedAt clears the value of the "email_verified_at" field.
func (u *AdminUpsertBulk) ClearEmailVerifiedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearEmailVerifiedAt()
	})
}

// SetPassword sets the "password" field.
func (u *AdminUpsertBulk) SetPassword(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdatePassword() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword()
	})
}

// SetPassword2 sets the "password2" field.
func (u *AdminUpsertBulk) SetPassword2(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword2(v)
	})
}

// UpdatePassword2 sets the "password2" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdatePassword2() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword2()
	})
}

// SetLang sets the "lang" field.
func (u *AdminUpsertBulk) SetLang(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateLang() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateLang()
	})
}

// SetAvatar sets the "avatar" field.
func (u *AdminUpsertBulk) SetAvatar(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateAvatar() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateAvatar()
	})
}

// ClearAvatar clears the value of the "avatar" field.
func (u *AdminUpsertBulk) ClearAvatar() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearAvatar()
	})
}

// SetType sets the "type" field.
func (u *AdminUpsertBulk) SetType(v uint8) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *AdminUpsertBulk) AddType(v uint8) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateType() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsertBulk) SetCreatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateCreatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertBulk) SetUpdatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUpdatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsertBulk) SetDeletedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateDeletedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AdminUpsertBulk) ClearDeletedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AdminUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AdminCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AdminCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
