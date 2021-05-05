// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "currency", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_users_accounts",
				Columns:    []*schema.Column{AccountsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "owner_currency_key",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[5], AccountsColumns[4]},
			},
		},
	}
	// EntriesColumns holds the columns for the "entries" table.
	EntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "account_id", Type: field.TypeUUID, Nullable: true},
	}
	// EntriesTable holds the schema information for the "entries" table.
	EntriesTable = &schema.Table{
		Name:       "entries",
		Columns:    EntriesColumns,
		PrimaryKey: []*schema.Column{EntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "entries_accounts_entries",
				Columns:    []*schema.Column{EntriesColumns[4]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransfersColumns holds the columns for the "transfers" table.
	TransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "from_account_id", Type: field.TypeUUID, Nullable: true},
		{Name: "to_account_id", Type: field.TypeUUID, Nullable: true},
	}
	// TransfersTable holds the schema information for the "transfers" table.
	TransfersTable = &schema.Table{
		Name:       "transfers",
		Columns:    TransfersColumns,
		PrimaryKey: []*schema.Column{TransfersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transfers_accounts_outbounds",
				Columns:    []*schema.Column{TransfersColumns[4]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transfers_accounts_inbounds",
				Columns:    []*schema.Column{TransfersColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "username", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "password", Type: field.TypeString, Size: 256},
		{Name: "full_name", Type: field.TypeString, Size: 75},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 256},
		{Name: "password_changed_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		EntriesTable,
		TransfersTable,
		UsersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = UsersTable
	EntriesTable.ForeignKeys[0].RefTable = AccountsTable
	TransfersTable.ForeignKeys[0].RefTable = AccountsTable
	TransfersTable.ForeignKeys[1].RefTable = AccountsTable
}
