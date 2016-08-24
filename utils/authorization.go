// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import "github.com/mattermost/platform/model"

func SetDefaultRolesBasedOnConfig() {
	// Reset the roles to default to make this logic easier
	model.InitalizeRoles()

	switch *Cfg.TeamSettings.RestrictPublicChannelManagement {
	case model.PERMISSIONS_ALL:
		model.R_CHANNEL_USER.Permissions = append(
			model.R_CHANNEL_USER.Permissions,
			model.P_MANAGE_PUBLIC_CHANNEL_PROPERTIES.Id,
		)
		model.R_TEAM_USER.Permissions = append(
			model.R_TEAM_USER.Permissions,
			model.P_DELETE_PUBLIC_CHANNEL.Id,
			model.P_CREATE_PUBLIC_CHANNEL.Id,
		)
		break
	case model.PERMISSIONS_TEAM_ADMIN:
		model.R_TEAM_ADMIN.Permissions = append(
			model.R_TEAM_ADMIN.Permissions,
			model.P_MANAGE_PUBLIC_CHANNEL_PROPERTIES.Id,
			model.P_DELETE_PUBLIC_CHANNEL.Id,
			model.P_CREATE_PUBLIC_CHANNEL.Id,
		)
		break
	}

	switch *Cfg.TeamSettings.RestrictPrivateChannelManagement {
	case model.PERMISSIONS_ALL:
		model.R_CHANNEL_USER.Permissions = append(
			model.R_CHANNEL_USER.Permissions,
			model.P_MANAGE_PRIVATE_CHANNEL_PROPERTIES.Id,
		)
		model.R_TEAM_USER.Permissions = append(
			model.R_TEAM_USER.Permissions,
			model.P_DELETE_PRIVATE_CHANNEL.Id,
			model.P_CREATE_PRIVATE_CHANNEL.Id,
		)
		break
	case model.PERMISSIONS_TEAM_ADMIN:
		model.R_TEAM_ADMIN.Permissions = append(
			model.R_TEAM_ADMIN.Permissions,
			model.P_MANAGE_PRIVATE_CHANNEL_PROPERTIES.Id,
			model.P_DELETE_PRIVATE_CHANNEL.Id,
			model.P_CREATE_PRIVATE_CHANNEL.Id,
		)
		break
	}

	if !*Cfg.ServiceSettings.EnableOnlyAdminIntegrations {
		model.R_TEAM_USER.Permissions = append(
			model.R_TEAM_USER.Permissions,
			model.P_MANAGE_WEBHOOKS.Id,
			model.P_MANAGE_SLASH_COMMANDS.Id,
		)
		model.R_SYSTEM_USER.Permissions = append(
			model.R_SYSTEM_USER.Permissions,
			model.P_MANAGE_OAUTH.Id,
		)
	}

	// If team admins are given permission
	if *Cfg.TeamSettings.RestrictTeamInvite == model.PERMISSIONS_TEAM_ADMIN {
		model.R_TEAM_ADMIN.Permissions = append(
			model.R_TEAM_ADMIN.Permissions,
			model.P_INVITE_USER.Id,
		)
		// If it's not restricted to system admin or team admin, then give all users permission
	} else if *Cfg.TeamSettings.RestrictTeamInvite != model.PERMISSIONS_SYSTEM_ADMIN {
		model.R_SYSTEM_USER.Permissions = append(
			model.R_SYSTEM_USER.Permissions,
			model.P_INVITE_USER.Id,
		)
	}
}
