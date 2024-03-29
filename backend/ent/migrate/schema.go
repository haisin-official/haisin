// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ServicesColumns holds the columns for the "services" table.
	ServicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "service", Type: field.TypeEnum, Enums: []string{"twitter", "youtube", "fanbox"}},
		{Name: "url", Type: field.TypeString, Size: 2083},
		{Name: "user_uuid", Type: field.TypeUUID},
	}
	// ServicesTable holds the schema information for the "services" table.
	ServicesTable = &schema.Table{
		Name:       "services",
		Columns:    ServicesColumns,
		PrimaryKey: []*schema.Column{ServicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "services_users_uuid",
				Columns:    []*schema.Column{ServicesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 30},
		{Name: "ga", Type: field.TypeString, Nullable: true, Size: 100},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ServicesTable,
		UsersTable,
	}
)

func init() {
	ServicesTable.ForeignKeys[0].RefTable = UsersTable
}
