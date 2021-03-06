// Code generated by entc, DO NOT EDIT.

package ent

import (
	"simple-bank/ent/account"
	"simple-bank/ent/entry"
	"simple-bank/ent/schema"
	"simple-bank/ent/transfer"
	"simple-bank/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountMixinFields1 := accountMixin[1].Fields()
	_ = accountMixinFields1
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreateTime is the schema descriptor for create_time field.
	accountDescCreateTime := accountMixinFields1[0].Descriptor()
	// account.DefaultCreateTime holds the default value on creation for the create_time field.
	account.DefaultCreateTime = accountDescCreateTime.Default.(func() time.Time)
	// accountDescUpdateTime is the schema descriptor for update_time field.
	accountDescUpdateTime := accountMixinFields1[1].Descriptor()
	// account.DefaultUpdateTime holds the default value on creation for the update_time field.
	account.DefaultUpdateTime = accountDescUpdateTime.Default.(func() time.Time)
	// account.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	account.UpdateDefaultUpdateTime = accountDescUpdateTime.UpdateDefault.(func() time.Time)
	// accountDescOwner is the schema descriptor for owner field.
	accountDescOwner := accountFields[0].Descriptor()
	// account.OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	account.OwnerValidator = func() func(string) error {
		validators := accountDescOwner.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(user string) error {
			for _, fn := range fns {
				if err := fn(user); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// accountDescBalance is the schema descriptor for balance field.
	accountDescBalance := accountFields[1].Descriptor()
	// account.DefaultBalance holds the default value on creation for the balance field.
	account.DefaultBalance = accountDescBalance.Default.(float64)
	// accountDescCurrency is the schema descriptor for currency field.
	accountDescCurrency := accountFields[2].Descriptor()
	// account.CurrencyValidator is a validator for the "currency" field. It is called by the builders before save.
	account.CurrencyValidator = accountDescCurrency.Validators[0].(func(string) error)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountMixinFields0[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	entryMixin := schema.Entry{}.Mixin()
	entryMixinFields0 := entryMixin[0].Fields()
	_ = entryMixinFields0
	entryMixinFields1 := entryMixin[1].Fields()
	_ = entryMixinFields1
	entryFields := schema.Entry{}.Fields()
	_ = entryFields
	// entryDescCreateTime is the schema descriptor for create_time field.
	entryDescCreateTime := entryMixinFields1[0].Descriptor()
	// entry.DefaultCreateTime holds the default value on creation for the create_time field.
	entry.DefaultCreateTime = entryDescCreateTime.Default.(func() time.Time)
	// entryDescUpdateTime is the schema descriptor for update_time field.
	entryDescUpdateTime := entryMixinFields1[1].Descriptor()
	// entry.DefaultUpdateTime holds the default value on creation for the update_time field.
	entry.DefaultUpdateTime = entryDescUpdateTime.Default.(func() time.Time)
	// entry.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	entry.UpdateDefaultUpdateTime = entryDescUpdateTime.UpdateDefault.(func() time.Time)
	// entryDescID is the schema descriptor for id field.
	entryDescID := entryMixinFields0[0].Descriptor()
	// entry.DefaultID holds the default value on creation for the id field.
	entry.DefaultID = entryDescID.Default.(func() uuid.UUID)
	transferMixin := schema.Transfer{}.Mixin()
	transferMixinFields0 := transferMixin[0].Fields()
	_ = transferMixinFields0
	transferMixinFields1 := transferMixin[1].Fields()
	_ = transferMixinFields1
	transferFields := schema.Transfer{}.Fields()
	_ = transferFields
	// transferDescCreateTime is the schema descriptor for create_time field.
	transferDescCreateTime := transferMixinFields1[0].Descriptor()
	// transfer.DefaultCreateTime holds the default value on creation for the create_time field.
	transfer.DefaultCreateTime = transferDescCreateTime.Default.(func() time.Time)
	// transferDescUpdateTime is the schema descriptor for update_time field.
	transferDescUpdateTime := transferMixinFields1[1].Descriptor()
	// transfer.DefaultUpdateTime holds the default value on creation for the update_time field.
	transfer.DefaultUpdateTime = transferDescUpdateTime.Default.(func() time.Time)
	// transfer.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	transfer.UpdateDefaultUpdateTime = transferDescUpdateTime.UpdateDefault.(func() time.Time)
	// transferDescID is the schema descriptor for id field.
	transferDescID := transferMixinFields0[0].Descriptor()
	// transfer.DefaultID holds the default value on creation for the id field.
	transfer.DefaultID = transferDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[2].Descriptor()
	// user.FullNameValidator is a validator for the "full_name" field. It is called by the builders before save.
	user.FullNameValidator = func() func(string) error {
		validators := userDescFullName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(full_name string) error {
			for _, fn := range fns {
				if err := fn(full_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPasswordChangedAt is the schema descriptor for password_changed_at field.
	userDescPasswordChangedAt := userFields[4].Descriptor()
	// user.DefaultPasswordChangedAt holds the default value on creation for the password_changed_at field.
	user.DefaultPasswordChangedAt = userDescPasswordChangedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = func() func(string) error {
		validators := userDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
