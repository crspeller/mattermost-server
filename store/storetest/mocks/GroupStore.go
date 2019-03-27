// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/crspeller/mattermost-server/model"
import store "github.com/crspeller/mattermost-server/store"

// GroupStore is an autogenerated mock type for the GroupStore type
type GroupStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: group
func (_m *GroupStore) Create(group *model.Group) store.StoreChannel {
	ret := _m.Called(group)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Group) store.StoreChannel); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// CreateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	ret := _m.Called(groupSyncable)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) store.StoreChannel); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// CreateOrRestoreMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) CreateOrRestoreMember(groupID string, userID string) store.StoreChannel {
	ret := _m.Called(groupID, userID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: groupID
func (_m *GroupStore) Delete(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) store.StoreChannel); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) DeleteMember(groupID string, userID string) store.StoreChannel {
	ret := _m.Called(groupID, userID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: groupID
func (_m *GroupStore) Get(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllBySource provides a mock function with given fields: groupSource
func (_m *GroupStore) GetAllBySource(groupSource model.GroupSource) store.StoreChannel {
	ret := _m.Called(groupSource)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(model.GroupSource) store.StoreChannel); ok {
		r0 = rf(groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllGroupSyncablesByGroupId provides a mock function with given fields: groupID, syncableType
func (_m *GroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) store.StoreChannel {
	ret := _m.Called(groupID, syncableType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType) store.StoreChannel); ok {
		r0 = rf(groupID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByRemoteID provides a mock function with given fields: remoteID, groupSource
func (_m *GroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) store.StoreChannel {
	ret := _m.Called(remoteID, groupSource)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, model.GroupSource) store.StoreChannel); ok {
		r0 = rf(remoteID, groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) store.StoreChannel); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMemberCount provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberCount(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMemberUsers provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberUsers(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMemberUsersPage provides a mock function with given fields: groupID, offset, limit
func (_m *GroupStore) GetMemberUsersPage(groupID string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(groupID, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(groupID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PendingAutoAddChannelMembers provides a mock function with given fields: minGroupMembersCreateAt
func (_m *GroupStore) PendingAutoAddChannelMembers(minGroupMembersCreateAt int64) store.StoreChannel {
	ret := _m.Called(minGroupMembersCreateAt)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64) store.StoreChannel); ok {
		r0 = rf(minGroupMembersCreateAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PendingAutoAddTeamMembers provides a mock function with given fields: minGroupMembersCreateAt
func (_m *GroupStore) PendingAutoAddTeamMembers(minGroupMembersCreateAt int64) store.StoreChannel {
	ret := _m.Called(minGroupMembersCreateAt)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64) store.StoreChannel); ok {
		r0 = rf(minGroupMembersCreateAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: group
func (_m *GroupStore) Update(group *model.Group) store.StoreChannel {
	ret := _m.Called(group)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Group) store.StoreChannel); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	ret := _m.Called(groupSyncable)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) store.StoreChannel); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
