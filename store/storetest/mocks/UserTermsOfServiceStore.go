// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/crspeller/mattermost-server/model"
import store "github.com/crspeller/mattermost-server/store"

// UserTermsOfServiceStore is an autogenerated mock type for the UserTermsOfServiceStore type
type UserTermsOfServiceStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userId, termsOfServiceId
func (_m *UserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) store.StoreChannel {
	ret := _m.Called(userId, termsOfServiceId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, termsOfServiceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByUser provides a mock function with given fields: userId
func (_m *UserTermsOfServiceStore) GetByUser(userId string) store.StoreChannel {
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

// Save provides a mock function with given fields: userTermsOfService
func (_m *UserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) store.StoreChannel {
	ret := _m.Called(userTermsOfService)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.UserTermsOfService) store.StoreChannel); ok {
		r0 = rf(userTermsOfService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
