// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/zacmm/zacmm-server/model"
	mock "github.com/stretchr/testify/mock"
)

// PreferenceStore is an autogenerated mock type for the PreferenceStore type
type PreferenceStore struct {
	mock.Mock
}

// CleanupFlagsBatch provides a mock function with given fields: limit
func (_m *PreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userId, category, name
func (_m *PreferenceStore) Delete(userId string, category string, name string) error {
	ret := _m.Called(userId, category, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(userId, category, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: userId, category
func (_m *PreferenceStore) DeleteCategory(userId string, category string) error {
	ret := _m.Called(userId, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategoryAndName provides a mock function with given fields: category, name
func (_m *PreferenceStore) DeleteCategoryAndName(category string, name string) error {
	ret := _m.Called(category, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(category, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: userId, category, name
func (_m *PreferenceStore) Get(userId string, category string, name string) (*model.Preference, error) {
	ret := _m.Called(userId, category, name)

	var r0 *model.Preference
	if rf, ok := ret.Get(0).(func(string, string, string) *model.Preference); ok {
		r0 = rf(userId, category, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Preference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userId, category, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: userId
func (_m *PreferenceStore) GetAll(userId string) (model.Preferences, error) {
	ret := _m.Called(userId)

	var r0 model.Preferences
	if rf, ok := ret.Get(0).(func(string) model.Preferences); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Preferences)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategory provides a mock function with given fields: userId, category
func (_m *PreferenceStore) GetCategory(userId string, category string) (model.Preferences, error) {
	ret := _m.Called(userId, category)

	var r0 model.Preferences
	if rf, ok := ret.Get(0).(func(string, string) model.Preferences); ok {
		r0 = rf(userId, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Preferences)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *PreferenceStore) PermanentDeleteByUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: preferences
func (_m *PreferenceStore) Save(preferences *model.Preferences) error {
	ret := _m.Called(preferences)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Preferences) error); ok {
		r0 = rf(preferences)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
