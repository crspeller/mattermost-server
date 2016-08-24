// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	"testing"

	"github.com/mattermost/platform/model"
)

/*func TestCheckIfRolesGrantPermission(t *testing.T) {
	th := Setup().InitBasic()
}*/

func TestCheckIfRolesGrantPermission(t *testing.T) {
	Setup()

	cases := []struct {
		roles        []string
		permissionId string
		shouldGrant  bool
	}{
		{[]string{model.R_SYSTEM_ADMIN.Id}, model.R_SYSTEM_ADMIN.Permissions[0], true},
		{[]string{model.R_SYSTEM_ADMIN.Id}, "non-existant-permission", false},
		{[]string{model.R_CHANNEL_USER.Id}, model.R_CHANNEL_USER.Permissions[0], true},
		{[]string{model.R_CHANNEL_USER.Id}, model.P_MANAGE_SYSTEM.Id, false},
		{[]string{model.R_SYSTEM_ADMIN.Id, model.R_CHANNEL_USER.Id}, model.P_MANAGE_SYSTEM.Id, true},
		{[]string{model.R_CHANNEL_USER.Id, model.R_SYSTEM_ADMIN.Id}, model.P_MANAGE_SYSTEM.Id, true},
		{[]string{model.R_TEAM_USER.Id, model.R_TEAM_ADMIN.Id}, model.P_MANAGE_SLASH_COMMANDS.Id, true},
		{[]string{model.R_TEAM_ADMIN.Id, model.R_TEAM_USER.Id}, model.P_MANAGE_SLASH_COMMANDS.Id, true},
	}

	for testnum, testcase := range cases {
		if CheckIfRolesGrantPermission(testcase.roles, testcase.permissionId) != testcase.shouldGrant {
			t.Fatal("Failed test case ", testnum)
		}
	}

}
