// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "workspace_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString},
		{Name: "created_by", Type: field.TypeString},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
	}
	// ChannelMessagesColumns holds the columns for the "channel_messages" table.
	ChannelMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message_ids", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ChannelMessagesTable holds the schema information for the "channel_messages" table.
	ChannelMessagesTable = &schema.Table{
		Name:       "channel_messages",
		Columns:    ChannelMessagesColumns,
		PrimaryKey: []*schema.Column{ChannelMessagesColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "data", Type: field.TypeString},
		{Name: "attachment", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "workspace_id", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// UserChannelsColumns holds the columns for the "user_channels" table.
	UserChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "channel_id", Type: field.TypeInt64},
		{Name: "role_type", Type: field.TypeString},
	}
	// UserChannelsTable holds the schema information for the "user_channels" table.
	UserChannelsTable = &schema.Table{
		Name:       "user_channels",
		Columns:    UserChannelsColumns,
		PrimaryKey: []*schema.Column{UserChannelsColumns[0]},
	}
	// UserWorkspacesColumns holds the columns for the "user_workspaces" table.
	UserWorkspacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "workspace_id", Type: field.TypeInt64},
		{Name: "role_type", Type: field.TypeString},
	}
	// UserWorkspacesTable holds the schema information for the "user_workspaces" table.
	UserWorkspacesTable = &schema.Table{
		Name:       "user_workspaces",
		Columns:    UserWorkspacesColumns,
		PrimaryKey: []*schema.Column{UserWorkspacesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "phone_number", Type: field.TypeString},
		{Name: "user_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "version_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WorkspacesColumns holds the columns for the "workspaces" table.
	WorkspacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// WorkspacesTable holds the schema information for the "workspaces" table.
	WorkspacesTable = &schema.Table{
		Name:       "workspaces",
		Columns:    WorkspacesColumns,
		PrimaryKey: []*schema.Column{WorkspacesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		ChannelMessagesTable,
		MessagesTable,
		UserChannelsTable,
		UserWorkspacesTable,
		UsersTable,
		WorkspacesTable,
	}
)

func init() {
}
