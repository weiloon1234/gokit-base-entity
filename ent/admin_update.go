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
	"github.com/weiloon1234/gokit-base-entity/ent/predicate"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where appends a list predicates to the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUsername sets the "username" field.
func (au *AdminUpdate) SetUsername(s string) *AdminUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AdminUpdate) SetNillableUsername(s *string) *AdminUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// ClearUsername clears the value of the "username" field.
func (au *AdminUpdate) ClearUsername() *AdminUpdate {
	au.mutation.ClearUsername()
	return au
}

// SetName sets the "name" field.
func (au *AdminUpdate) SetName(s string) *AdminUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AdminUpdate) SetNillableName(s *string) *AdminUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *AdminUpdate) ClearName() *AdminUpdate {
	au.mutation.ClearName()
	return au
}

// SetEmail sets the "email" field.
func (au *AdminUpdate) SetEmail(s string) *AdminUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (au *AdminUpdate) SetNillableEmail(s *string) *AdminUpdate {
	if s != nil {
		au.SetEmail(*s)
	}
	return au
}

// ClearEmail clears the value of the "email" field.
func (au *AdminUpdate) ClearEmail() *AdminUpdate {
	au.mutation.ClearEmail()
	return au
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (au *AdminUpdate) SetEmailVerifiedAt(t time.Time) *AdminUpdate {
	au.mutation.SetEmailVerifiedAt(t)
	return au
}

// SetNillableEmailVerifiedAt sets the "email_verified_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableEmailVerifiedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetEmailVerifiedAt(*t)
	}
	return au
}

// ClearEmailVerifiedAt clears the value of the "email_verified_at" field.
func (au *AdminUpdate) ClearEmailVerifiedAt() *AdminUpdate {
	au.mutation.ClearEmailVerifiedAt()
	return au
}

// SetPassword sets the "password" field.
func (au *AdminUpdate) SetPassword(s string) *AdminUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AdminUpdate) SetNillablePassword(s *string) *AdminUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetPassword2 sets the "password2" field.
func (au *AdminUpdate) SetPassword2(s string) *AdminUpdate {
	au.mutation.SetPassword2(s)
	return au
}

// SetNillablePassword2 sets the "password2" field if the given value is not nil.
func (au *AdminUpdate) SetNillablePassword2(s *string) *AdminUpdate {
	if s != nil {
		au.SetPassword2(*s)
	}
	return au
}

// SetLang sets the "lang" field.
func (au *AdminUpdate) SetLang(s string) *AdminUpdate {
	au.mutation.SetLang(s)
	return au
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (au *AdminUpdate) SetNillableLang(s *string) *AdminUpdate {
	if s != nil {
		au.SetLang(*s)
	}
	return au
}

// SetAvatar sets the "avatar" field.
func (au *AdminUpdate) SetAvatar(s string) *AdminUpdate {
	au.mutation.SetAvatar(s)
	return au
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (au *AdminUpdate) SetNillableAvatar(s *string) *AdminUpdate {
	if s != nil {
		au.SetAvatar(*s)
	}
	return au
}

// ClearAvatar clears the value of the "avatar" field.
func (au *AdminUpdate) ClearAvatar() *AdminUpdate {
	au.mutation.ClearAvatar()
	return au
}

// SetType sets the "type" field.
func (au *AdminUpdate) SetType(u uint8) *AdminUpdate {
	au.mutation.ResetType()
	au.mutation.SetType(u)
	return au
}

// SetNillableType sets the "type" field if the given value is not nil.
func (au *AdminUpdate) SetNillableType(u *uint8) *AdminUpdate {
	if u != nil {
		au.SetType(*u)
	}
	return au
}

// AddType adds u to the "type" field.
func (au *AdminUpdate) AddType(u int8) *AdminUpdate {
	au.mutation.AddType(u)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AdminUpdate) SetCreatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableCreatedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AdminUpdate) SetUpdatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AdminUpdate) SetDeletedAt(t time.Time) *AdminUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableDeletedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AdminUpdate) ClearDeletedAt() *AdminUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AdminUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Admin.password": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password2(); ok {
		if err := admin.Password2Validator(v); err != nil {
			return &ValidationError{Name: "password2", err: fmt.Errorf(`ent: validator failed for field "Admin.password2": %w`, err)}
		}
	}
	return nil
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if au.mutation.UsernameCleared() {
		_spec.ClearField(admin.FieldUsername, field.TypeString)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(admin.FieldName, field.TypeString, value)
	}
	if au.mutation.NameCleared() {
		_spec.ClearField(admin.FieldName, field.TypeString)
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(admin.FieldEmail, field.TypeString, value)
	}
	if au.mutation.EmailCleared() {
		_spec.ClearField(admin.FieldEmail, field.TypeString)
	}
	if value, ok := au.mutation.EmailVerifiedAt(); ok {
		_spec.SetField(admin.FieldEmailVerifiedAt, field.TypeTime, value)
	}
	if au.mutation.EmailVerifiedAtCleared() {
		_spec.ClearField(admin.FieldEmailVerifiedAt, field.TypeTime)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.Password2(); ok {
		_spec.SetField(admin.FieldPassword2, field.TypeString, value)
	}
	if value, ok := au.mutation.Lang(); ok {
		_spec.SetField(admin.FieldLang, field.TypeString, value)
	}
	if value, ok := au.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
	}
	if au.mutation.AvatarCleared() {
		_spec.ClearField(admin.FieldAvatar, field.TypeString)
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.SetField(admin.FieldType, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedType(); ok {
		_spec.AddField(admin.FieldType, field.TypeUint8, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(admin.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdminMutation
}

// SetUsername sets the "username" field.
func (auo *AdminUpdateOne) SetUsername(s string) *AdminUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableUsername(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// ClearUsername clears the value of the "username" field.
func (auo *AdminUpdateOne) ClearUsername() *AdminUpdateOne {
	auo.mutation.ClearUsername()
	return auo
}

// SetName sets the "name" field.
func (auo *AdminUpdateOne) SetName(s string) *AdminUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableName(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *AdminUpdateOne) ClearName() *AdminUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetEmail sets the "email" field.
func (auo *AdminUpdateOne) SetEmail(s string) *AdminUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableEmail(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetEmail(*s)
	}
	return auo
}

// ClearEmail clears the value of the "email" field.
func (auo *AdminUpdateOne) ClearEmail() *AdminUpdateOne {
	auo.mutation.ClearEmail()
	return auo
}

// SetEmailVerifiedAt sets the "email_verified_at" field.
func (auo *AdminUpdateOne) SetEmailVerifiedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetEmailVerifiedAt(t)
	return auo
}

// SetNillableEmailVerifiedAt sets the "email_verified_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableEmailVerifiedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetEmailVerifiedAt(*t)
	}
	return auo
}

// ClearEmailVerifiedAt clears the value of the "email_verified_at" field.
func (auo *AdminUpdateOne) ClearEmailVerifiedAt() *AdminUpdateOne {
	auo.mutation.ClearEmailVerifiedAt()
	return auo
}

// SetPassword sets the "password" field.
func (auo *AdminUpdateOne) SetPassword(s string) *AdminUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillablePassword(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetPassword2 sets the "password2" field.
func (auo *AdminUpdateOne) SetPassword2(s string) *AdminUpdateOne {
	auo.mutation.SetPassword2(s)
	return auo
}

// SetNillablePassword2 sets the "password2" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillablePassword2(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetPassword2(*s)
	}
	return auo
}

// SetLang sets the "lang" field.
func (auo *AdminUpdateOne) SetLang(s string) *AdminUpdateOne {
	auo.mutation.SetLang(s)
	return auo
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableLang(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetLang(*s)
	}
	return auo
}

// SetAvatar sets the "avatar" field.
func (auo *AdminUpdateOne) SetAvatar(s string) *AdminUpdateOne {
	auo.mutation.SetAvatar(s)
	return auo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableAvatar(s *string) *AdminUpdateOne {
	if s != nil {
		auo.SetAvatar(*s)
	}
	return auo
}

// ClearAvatar clears the value of the "avatar" field.
func (auo *AdminUpdateOne) ClearAvatar() *AdminUpdateOne {
	auo.mutation.ClearAvatar()
	return auo
}

// SetType sets the "type" field.
func (auo *AdminUpdateOne) SetType(u uint8) *AdminUpdateOne {
	auo.mutation.ResetType()
	auo.mutation.SetType(u)
	return auo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableType(u *uint8) *AdminUpdateOne {
	if u != nil {
		auo.SetType(*u)
	}
	return auo
}

// AddType adds u to the "type" field.
func (auo *AdminUpdateOne) AddType(u int8) *AdminUpdateOne {
	auo.mutation.AddType(u)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AdminUpdateOne) SetCreatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AdminUpdateOne) SetUpdatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AdminUpdateOne) SetDeletedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableDeletedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AdminUpdateOne) ClearDeletedAt() *AdminUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// Where appends a list predicates to the AdminUpdate builder.
func (auo *AdminUpdateOne) Where(ps ...predicate.Admin) *AdminUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdminUpdateOne) Select(field string, fields ...string) *AdminUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AdminUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Admin.password": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password2(); ok {
		if err := admin.Password2Validator(v); err != nil {
			return &ValidationError{Name: "password2", err: fmt.Errorf(`ent: validator failed for field "Admin.password2": %w`, err)}
		}
	}
	return nil
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(admin.Table, admin.Columns, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Admin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, admin.FieldID)
		for _, f := range fields {
			if !admin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != admin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(admin.FieldUsername, field.TypeString, value)
	}
	if auo.mutation.UsernameCleared() {
		_spec.ClearField(admin.FieldUsername, field.TypeString)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(admin.FieldName, field.TypeString, value)
	}
	if auo.mutation.NameCleared() {
		_spec.ClearField(admin.FieldName, field.TypeString)
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(admin.FieldEmail, field.TypeString, value)
	}
	if auo.mutation.EmailCleared() {
		_spec.ClearField(admin.FieldEmail, field.TypeString)
	}
	if value, ok := auo.mutation.EmailVerifiedAt(); ok {
		_spec.SetField(admin.FieldEmailVerifiedAt, field.TypeTime, value)
	}
	if auo.mutation.EmailVerifiedAtCleared() {
		_spec.ClearField(admin.FieldEmailVerifiedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password2(); ok {
		_spec.SetField(admin.FieldPassword2, field.TypeString, value)
	}
	if value, ok := auo.mutation.Lang(); ok {
		_spec.SetField(admin.FieldLang, field.TypeString, value)
	}
	if value, ok := auo.mutation.Avatar(); ok {
		_spec.SetField(admin.FieldAvatar, field.TypeString, value)
	}
	if auo.mutation.AvatarCleared() {
		_spec.ClearField(admin.FieldAvatar, field.TypeString)
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.SetField(admin.FieldType, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedType(); ok {
		_spec.AddField(admin.FieldType, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(admin.FieldDeletedAt, field.TypeTime)
	}
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
