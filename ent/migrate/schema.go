// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "user_tenant", Type: field.TypeInt, Nullable: true},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizations_users_tenant",
				Columns:    []*schema.Column{OrganizationsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// OrganizationUsersColumns holds the columns for the "organization_users" table.
	OrganizationUsersColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OrganizationUsersTable holds the schema information for the "organization_users" table.
	OrganizationUsersTable = &schema.Table{
		Name:       "organization_users",
		Columns:    OrganizationUsersColumns,
		PrimaryKey: []*schema.Column{OrganizationUsersColumns[0], OrganizationUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_users_organization_id",
				Columns:    []*schema.Column{OrganizationUsersColumns[0]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_users_user_id",
				Columns:    []*schema.Column{OrganizationUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationsTable,
		UsersTable,
		OrganizationUsersTable,
	}
)

func init() {
	OrganizationsTable.ForeignKeys[0].RefTable = UsersTable
	OrganizationUsersTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationUsersTable.ForeignKeys[1].RefTable = UsersTable
}
