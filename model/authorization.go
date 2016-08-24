// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type Permission struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

var P_INVITE_USER *Permission
var P_ADD_USER_TO_TEAM *Permission
var P_USE_SLASH_COMMANDS *Permission
var P_MANAGE_SLASH_COMMANDS *Permission
var P_MANAGE_OTHERS_SLASH_COMMANDS *Permission
var P_CREATE_PUBLIC_CHANNEL *Permission
var P_CREATE_PRIVATE_CHANNEL *Permission
var P_MANAGE_PUBLIC_CHANNEL_MEMBERS *Permission
var P_MANAGE_PRIVATE_CHANNEL_MEMBERS *Permission
var P_ASSIGN_SYSTEM_ADMIN_ROLE *Permission
var P_MANAGE_ROLES *Permission
var P_CREATE_DIRECT_CHANNEL *Permission
var P_MANAGE_PUBLIC_CHANNEL_PROPERTIES *Permission
var P_MANAGE_PRIVATE_CHANNEL_PROPERTIES *Permission
var P_LIST_TEAM_CHANNELS *Permission
var P_JOIN_PUBLIC_CHANNELS *Permission
var P_DELETE_PUBLIC_CHANNEL *Permission
var P_DELETE_PRIVATE_CHANNEL *Permission
var P_EDIT_OTHER_USERS *Permission
var P_READ_CHANNEL *Permission
var P_PERMANENT_DELETE_USER *Permission
var P_UPLOAD_FILE *Permission
var P_GET_PUBLIC_LINK *Permission
var P_MANAGE_WEBHOOKS *Permission
var P_MANAGE_OTHERS_WEBHOOKS *Permission
var P_MANAGE_OAUTH *Permission
var P_MANAGE_SYSTEM_WIDE_OAUTH *Permission
var P_CREATE_POST *Permission
var P_EDIT_POST *Permission
var P_EDIT_OTHERS_POSTS *Permission
var P_REMOVE_USER_FROM_TEAM *Permission
var P_MANAGE_TEAM *Permission
var P_IMPORT_TEAM *Permission

// General permission that encompases all system admin functions
// in the future this could be broken up to allow access to some
// admin functions but not others
var P_MANAGE_SYSTEM *Permission

var R_SYSTEM_USER *Role
var R_SYSTEM_ADMIN *Role

var R_TEAM_USER *Role
var R_TEAM_ADMIN *Role

var R_CHANNEL_USER *Role
var R_CHANNEL_ADMIN *Role
var R_CHANNEL_GUEST *Role

var BuiltInRoles map[string]*Role

func InitalizePermissions() {
	P_INVITE_USER = &Permission{
		"invite_user",
		"authentication.permissions.team_invite_user.name",
		"authentication.permissions.team_invite_user.description",
	}
	P_ADD_USER_TO_TEAM = &Permission{
		"add_user_to_team",
		"authentication.permissions.add_user_to_team.name",
		"authentication.permissions.add_user_to_team.description",
	}
	P_USE_SLASH_COMMANDS = &Permission{
		"use_slash_commands",
		"authentication.permissions.team_use_slash_commands.name",
		"authentication.permissions.team_use_slash_commands.description",
	}
	P_MANAGE_SLASH_COMMANDS = &Permission{
		"manage_slash_commands",
		"authentication.permissions.manage_slash_commands.name",
		"authentication.permissions.manage_slash_commands.description",
	}
	P_MANAGE_OTHERS_SLASH_COMMANDS = &Permission{
		"manage_others_slash_commands",
		"authentication.permissions.manage_others_slash_commands.name",
		"authentication.permissions.manage_others_slash_commands.description",
	}
	P_CREATE_PUBLIC_CHANNEL = &Permission{
		"create_public_channel",
		"authentication.permissions.create_public_channel.name",
		"authentication.permissions.create_public_channel.description",
	}
	P_CREATE_PRIVATE_CHANNEL = &Permission{
		"create_private_channel",
		"authentication.permissions.create_private_channel.name",
		"authentication.permissions.create_private_channel.description",
	}
	P_MANAGE_PUBLIC_CHANNEL_MEMBERS = &Permission{
		"manage_public_channel_members",
		"authentication.permissions.manage_public_channel_members.name",
		"authentication.permissions.manage_public_channel_members.description",
	}
	P_MANAGE_PRIVATE_CHANNEL_MEMBERS = &Permission{
		"manage_private_channel_members",
		"authentication.permissions.manage_private_channel_members.name",
		"authentication.permissions.manage_private_channel_members.description",
	}
	P_ASSIGN_SYSTEM_ADMIN_ROLE = &Permission{
		"assign_system_admin_role",
		"authentication.permissions.assign_system_admin_role.name",
		"authentication.permissions.assign_system_admin_role.description",
	}
	P_MANAGE_ROLES = &Permission{
		"manage_roles",
		"authentication.permissions.manage_roles.name",
		"authentication.permissions.manage_roles.description",
	}
	P_MANAGE_SYSTEM = &Permission{
		"manage_system",
		"authentication.permissions.manage_system.name",
		"authentication.permissions.manage_system.description",
	}
	P_CREATE_DIRECT_CHANNEL = &Permission{
		"create_direct_channel",
		"authentication.permissions.create_direct_channel.name",
		"authentication.permissions.create_direct_channel.description",
	}
	P_MANAGE_PUBLIC_CHANNEL_PROPERTIES = &Permission{
		"manage__publicchannel_properties",
		"authentication.permissions.manage_public_channel_properties.name",
		"authentication.permissions.manage_public_channel_properties.description",
	}
	P_MANAGE_PRIVATE_CHANNEL_PROPERTIES = &Permission{
		"manage_private_channel_properties",
		"authentication.permissions.manage_private_channel_properties.name",
		"authentication.permissions.manage_private_channel_properties.description",
	}
	P_LIST_TEAM_CHANNELS = &Permission{
		"list_team_channels",
		"authentication.permissions.list_team_channels.name",
		"authentication.permissions.list_team_channels.description",
	}
	P_JOIN_PUBLIC_CHANNELS = &Permission{
		"join_public_channels",
		"authentication.permissions.join_public_channels.name",
		"authentication.permissions.join_public_channels.description",
	}
	P_DELETE_PUBLIC_CHANNEL = &Permission{
		"delete_public_channel",
		"authentication.permissions.delete_public_channel.name",
		"authentication.permissions.delete_public_channel.description",
	}
	P_DELETE_PRIVATE_CHANNEL = &Permission{
		"delete_private_channel",
		"authentication.permissions.delete_private_channel.name",
		"authentication.permissions.delete_private_channel.description",
	}
	P_EDIT_OTHER_USERS = &Permission{
		"edit_other_users",
		"authentication.permissions.edit_other_users.name",
		"authentication.permissions.edit_other_users.description",
	}
	P_READ_CHANNEL = &Permission{
		"read_channel",
		"authentication.permissions.read_channel.name",
		"authentication.permissions.read_channel.description",
	}
	P_PERMANENT_DELETE_USER = &Permission{
		"permanent_delete_user",
		"authentication.permissions.permanent_delete_user.name",
		"authentication.permissions.permanent_delete_user.description",
	}
	P_UPLOAD_FILE = &Permission{
		"upload_file",
		"authentication.permissions.upload_file.name",
		"authentication.permissions.upload_file.description",
	}
	P_GET_PUBLIC_LINK = &Permission{
		"get_public_link",
		"authentication.permissions.get_public_link.name",
		"authentication.permissions.get_public_link.description",
	}
	P_MANAGE_WEBHOOKS = &Permission{
		"manage_webhooks",
		"authentication.permissions.manage_webhooks.name",
		"authentication.permissions.manage_webhooks.description",
	}
	P_MANAGE_OTHERS_WEBHOOKS = &Permission{
		"manage_others_webhooks",
		"authentication.permissions.manage_others_webhooks.name",
		"authentication.permissions.manage_others_webhooks.description",
	}
	P_MANAGE_OAUTH = &Permission{
		"manage_oauth",
		"authentication.permissions.manage_oauth.name",
		"authentication.permissions.manage_oauth.description",
	}
	P_MANAGE_SYSTEM_WIDE_OAUTH = &Permission{
		"manage_sytem_wide_oauth",
		"authentication.permissions.manage_sytem_wide_oauth.name",
		"authentication.permissions.manage_sytem_wide_oauth.description",
	}
	P_CREATE_POST = &Permission{
		"create_post",
		"authentication.permissions.create_post.name",
		"authentication.permissions.create_post.description",
	}
	P_EDIT_POST = &Permission{
		"edit_post",
		"authentication.permissions.edit_post.name",
		"authentication.permissions.edit_post.description",
	}
	P_EDIT_OTHERS_POSTS = &Permission{
		"edit_others_posts",
		"authentication.permissions.edit_others_posts.name",
		"authentication.permissions.edit_others_posts.description",
	}
	P_REMOVE_USER_FROM_TEAM = &Permission{
		"remove_user_from_team",
		"authentication.permissions.remove_user_from_team.name",
		"authentication.permissions.remove_user_from_team.description",
	}
	P_MANAGE_TEAM = &Permission{
		"manage_team",
		"authentication.permissions.manage_team.name",
		"authentication.permissions.manage_team.description",
	}
	P_IMPORT_TEAM = &Permission{
		"import_team",
		"authentication.permissions.import_team.name",
		"authentication.permissions.import_team.description",
	}
}

func InitalizeRoles() {
	InitalizePermissions()
	BuiltInRoles = make(map[string]*Role)

	R_CHANNEL_USER = &Role{
		"channel_user",
		"authentication.roles.channel_user.name",
		"authentication.roles.channel_user.description",
		[]string{
			P_READ_CHANNEL.Id,
			P_MANAGE_PUBLIC_CHANNEL_MEMBERS.Id,
			P_MANAGE_PRIVATE_CHANNEL_MEMBERS.Id,
			P_UPLOAD_FILE.Id,
			P_GET_PUBLIC_LINK.Id,
			P_CREATE_POST.Id,
			P_EDIT_POST.Id,
			P_USE_SLASH_COMMANDS.Id,
		},
	}
	BuiltInRoles[R_CHANNEL_USER.Id] = R_CHANNEL_USER
	R_CHANNEL_ADMIN = &Role{
		"channel_admin",
		"authentication.roles.channel_admin.name",
		"authentication.roles.channel_admin.description",
		[]string{},
	}
	BuiltInRoles[R_CHANNEL_ADMIN.Id] = R_CHANNEL_ADMIN
	R_CHANNEL_GUEST = &Role{
		"guest",
		"authentication.roles.global_guest.name",
		"authentication.roles.global_guest.description",
		[]string{},
	}
	BuiltInRoles[R_CHANNEL_GUEST.Id] = R_CHANNEL_GUEST

	R_TEAM_USER = &Role{
		"team_user",
		"authentication.roles.team_user.name",
		"authentication.roles.team_user.description",
		[]string{
			P_LIST_TEAM_CHANNELS.Id,
			P_JOIN_PUBLIC_CHANNELS.Id,
		},
	}
	BuiltInRoles[R_TEAM_USER.Id] = R_TEAM_USER
	R_TEAM_ADMIN = &Role{
		"team_admin",
		"authentication.roles.team_admin.name",
		"authentication.roles.team_admin.description",
		[]string{
			P_EDIT_OTHERS_POSTS.Id,
			P_ADD_USER_TO_TEAM.Id,
			P_REMOVE_USER_FROM_TEAM.Id,
			P_MANAGE_TEAM.Id,
			P_IMPORT_TEAM.Id,
			P_MANAGE_ROLES.Id,
			P_MANAGE_OTHERS_WEBHOOKS.Id,
			P_MANAGE_SLASH_COMMANDS.Id,
			P_MANAGE_OTHERS_SLASH_COMMANDS.Id,
			P_MANAGE_WEBHOOKS.Id,
		},
	}
	BuiltInRoles[R_TEAM_ADMIN.Id] = R_TEAM_ADMIN

	R_SYSTEM_USER = &Role{
		"system_user",
		"authentication.roles.global_user.name",
		"authentication.roles.global_user.description",
		[]string{
			P_CREATE_DIRECT_CHANNEL.Id,
			P_PERMANENT_DELETE_USER.Id,
		},
	}
	BuiltInRoles[R_SYSTEM_USER.Id] = R_SYSTEM_USER
	R_SYSTEM_ADMIN = &Role{
		"system_admin",
		"authentication.roles.global_admin.name",
		"authentication.roles.global_admin.description",
		// System admins can do anything channel and team admins can do
		// plus everything members of teams and channels can do to all teams
		// and channels on the system
		append(
			append(
				append(
					append(
						[]string{
							P_ASSIGN_SYSTEM_ADMIN_ROLE.Id,
							P_MANAGE_SYSTEM.Id,
							P_MANAGE_PUBLIC_CHANNEL_PROPERTIES.Id,
							P_DELETE_PUBLIC_CHANNEL.Id,
							P_CREATE_PUBLIC_CHANNEL.Id,
							P_MANAGE_PRIVATE_CHANNEL_PROPERTIES.Id,
							P_DELETE_PRIVATE_CHANNEL.Id,
							P_CREATE_PRIVATE_CHANNEL.Id,
							P_MANAGE_SYSTEM_WIDE_OAUTH.Id,
							P_MANAGE_OTHERS_WEBHOOKS.Id,
							P_EDIT_OTHER_USERS.Id,
							P_MANAGE_OAUTH.Id,
							P_INVITE_USER.Id,
						},
						R_TEAM_USER.Permissions...,
					),
					R_CHANNEL_USER.Permissions...,
				),
				R_TEAM_ADMIN.Permissions...,
			),
			R_CHANNEL_ADMIN.Permissions...,
		),
	}
	BuiltInRoles[R_SYSTEM_ADMIN.Id] = R_SYSTEM_ADMIN

}

func RoleIdsToString(roles []string) string {
	output := ""
	for _, role := range roles {
		output += role + ", "
	}

	if output == "" {
		return "[<NO ROLES>]"
	}

	return output[:len(output)-1]
}

func init() {
	InitalizeRoles()
}
