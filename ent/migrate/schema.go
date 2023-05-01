// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "password", Type: field.TypeString},
		{Name: "card", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "cp", Type: field.TypeInt, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true, Default: "https://card.syui.ai"},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_card", Type: field.TypeInt},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_users_card",
				Columns:    []*schema.Column{CardsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "group_name",
				Unique:  true,
				Columns: []*schema.Column{GroupsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "did", Type: field.TypeString, Nullable: true},
		{Name: "delete", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "next", Type: field.TypeString, Nullable: true, Default: "20230503"},
		{Name: "group_users", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_users",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		GroupsTable,
		UsersTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
}
