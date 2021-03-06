// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/crspeller/mattermost-server/model"
import store "github.com/crspeller/mattermost-server/store"

// TeamStore is an autogenerated mock type for the TeamStore type
type TeamStore struct {
	mock.Mock
}

// AnalyticsGetTeamCountForScheme provides a mock function with given fields: schemeId
func (_m *TeamStore) AnalyticsGetTeamCountForScheme(schemeId string) store.StoreChannel {
	ret := _m.Called(schemeId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(schemeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// AnalyticsTeamCount provides a mock function with given fields:
func (_m *TeamStore) AnalyticsTeamCount() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ClearAllCustomRoleAssignments provides a mock function with given fields:
func (_m *TeamStore) ClearAllCustomRoleAssignments() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *TeamStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetActiveMemberCount provides a mock function with given fields: teamId
func (_m *TeamStore) GetActiveMemberCount(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TeamStore) GetAll() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllForExportAfter provides a mock function with given fields: limit, afterId
func (_m *TeamStore) GetAllForExportAfter(limit int, afterId string) store.StoreChannel {
	ret := _m.Called(limit, afterId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, string) store.StoreChannel); ok {
		r0 = rf(limit, afterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPage provides a mock function with given fields: offset, limit
func (_m *TeamStore) GetAllPage(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPrivateTeamListing provides a mock function with given fields:
func (_m *TeamStore) GetAllPrivateTeamListing() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPrivateTeamPageListing provides a mock function with given fields: offset, limit
func (_m *TeamStore) GetAllPrivateTeamPageListing(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllTeamListing provides a mock function with given fields:
func (_m *TeamStore) GetAllTeamListing() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllTeamPageListing provides a mock function with given fields: offset, limit
func (_m *TeamStore) GetAllTeamPageListing(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByInviteId provides a mock function with given fields: inviteId
func (_m *TeamStore) GetByInviteId(inviteId string) store.StoreChannel {
	ret := _m.Called(inviteId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(inviteId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByName provides a mock function with given fields: name
func (_m *TeamStore) GetByName(name string) store.StoreChannel {
	ret := _m.Called(name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannelUnreadsForAllTeams provides a mock function with given fields: excludeTeamId, userId
func (_m *TeamStore) GetChannelUnreadsForAllTeams(excludeTeamId string, userId string) store.StoreChannel {
	ret := _m.Called(excludeTeamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(excludeTeamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannelUnreadsForTeam provides a mock function with given fields: teamId, userId
func (_m *TeamStore) GetChannelUnreadsForTeam(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMember provides a mock function with given fields: teamId, userId
func (_m *TeamStore) GetMember(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMembers provides a mock function with given fields: teamId, offset, limit
func (_m *TeamStore) GetMembers(teamId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(teamId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(teamId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMembersByIds provides a mock function with given fields: teamId, userIds
func (_m *TeamStore) GetMembersByIds(teamId string, userIds []string) store.StoreChannel {
	ret := _m.Called(teamId, userIds)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, []string) store.StoreChannel); ok {
		r0 = rf(teamId, userIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamMembersForExport provides a mock function with given fields: userId
func (_m *TeamStore) GetTeamMembersForExport(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamsByScheme provides a mock function with given fields: schemeId, offset, limit
func (_m *TeamStore) GetTeamsByScheme(schemeId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(schemeId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(schemeId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamsByUserId provides a mock function with given fields: userId
func (_m *TeamStore) GetTeamsByUserId(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamsForUser provides a mock function with given fields: userId
func (_m *TeamStore) GetTeamsForUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTeamsForUserWithPagination provides a mock function with given fields: userId, page, perPage
func (_m *TeamStore) GetTeamsForUserWithPagination(userId string, page int, perPage int) store.StoreChannel {
	ret := _m.Called(userId, page, perPage)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTotalMemberCount provides a mock function with given fields: teamId
func (_m *TeamStore) GetTotalMemberCount(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// MigrateTeamMembers provides a mock function with given fields: fromTeamId, fromUserId
func (_m *TeamStore) MigrateTeamMembers(fromTeamId string, fromUserId string) store.StoreChannel {
	ret := _m.Called(fromTeamId, fromUserId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(fromTeamId, fromUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDelete provides a mock function with given fields: teamId
func (_m *TeamStore) PermanentDelete(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveAllMembersByTeam provides a mock function with given fields: teamId
func (_m *TeamStore) RemoveAllMembersByTeam(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveAllMembersByUser provides a mock function with given fields: userId
func (_m *TeamStore) RemoveAllMembersByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveMember provides a mock function with given fields: teamId, userId
func (_m *TeamStore) RemoveMember(teamId string, userId string) store.StoreChannel {
	ret := _m.Called(teamId, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(teamId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ResetAllTeamSchemes provides a mock function with given fields:
func (_m *TeamStore) ResetAllTeamSchemes() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: team
func (_m *TeamStore) Save(team *model.Team) store.StoreChannel {
	ret := _m.Called(team)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Team) store.StoreChannel); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveMember provides a mock function with given fields: member, maxUsersPerTeam
func (_m *TeamStore) SaveMember(member *model.TeamMember, maxUsersPerTeam int) store.StoreChannel {
	ret := _m.Called(member, maxUsersPerTeam)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.TeamMember, int) store.StoreChannel); ok {
		r0 = rf(member, maxUsersPerTeam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchAll provides a mock function with given fields: term
func (_m *TeamStore) SearchAll(term string) store.StoreChannel {
	ret := _m.Called(term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchByName provides a mock function with given fields: name
func (_m *TeamStore) SearchByName(name string) store.StoreChannel {
	ret := _m.Called(name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchOpen provides a mock function with given fields: term
func (_m *TeamStore) SearchOpen(term string) store.StoreChannel {
	ret := _m.Called(term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SearchPrivate provides a mock function with given fields: term
func (_m *TeamStore) SearchPrivate(term string) store.StoreChannel {
	ret := _m.Called(term)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(term)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: team
func (_m *TeamStore) Update(team *model.Team) store.StoreChannel {
	ret := _m.Called(team)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Team) store.StoreChannel); ok {
		r0 = rf(team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateDisplayName provides a mock function with given fields: name, teamId
func (_m *TeamStore) UpdateDisplayName(name string, teamId string) store.StoreChannel {
	ret := _m.Called(name, teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(name, teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateLastTeamIconUpdate provides a mock function with given fields: teamId, curTime
func (_m *TeamStore) UpdateLastTeamIconUpdate(teamId string, curTime int64) store.StoreChannel {
	ret := _m.Called(teamId, curTime)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(teamId, curTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateMember provides a mock function with given fields: member
func (_m *TeamStore) UpdateMember(member *model.TeamMember) store.StoreChannel {
	ret := _m.Called(member)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.TeamMember) store.StoreChannel); ok {
		r0 = rf(member)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
