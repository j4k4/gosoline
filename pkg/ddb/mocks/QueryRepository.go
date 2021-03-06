// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import djoemo "github.com/adjoeio/djoemo"
import mock "github.com/stretchr/testify/mock"

// QueryRepository is an autogenerated mock type for the QueryRepository type
type QueryRepository struct {
	mock.Mock
}

// GetItem provides a mock function with given fields: key, item
func (_m *QueryRepository) GetItem(key djoemo.KeyInterface, item interface{}) (bool, error) {
	ret := _m.Called(key, item)

	var r0 bool
	if rf, ok := ret.Get(0).(func(djoemo.KeyInterface, interface{}) bool); ok {
		r0 = rf(key, item)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(djoemo.KeyInterface, interface{}) error); ok {
		r1 = rf(key, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItemWithContext provides a mock function with given fields: ctx, key, item
func (_m *QueryRepository) GetItemWithContext(ctx context.Context, key djoemo.KeyInterface, item interface{}) (bool, error) {
	ret := _m.Called(ctx, key, item)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, djoemo.KeyInterface, interface{}) bool); ok {
		r0 = rf(ctx, key, item)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, djoemo.KeyInterface, interface{}) error); ok {
		r1 = rf(ctx, key, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItems provides a mock function with given fields: key, items
func (_m *QueryRepository) GetItems(key djoemo.KeyInterface, items interface{}) (bool, error) {
	ret := _m.Called(key, items)

	var r0 bool
	if rf, ok := ret.Get(0).(func(djoemo.KeyInterface, interface{}) bool); ok {
		r0 = rf(key, items)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(djoemo.KeyInterface, interface{}) error); ok {
		r1 = rf(key, items)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItemsWithContext provides a mock function with given fields: ctx, key, items
func (_m *QueryRepository) GetItemsWithContext(ctx context.Context, key djoemo.KeyInterface, items interface{}) (bool, error) {
	ret := _m.Called(ctx, key, items)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, djoemo.KeyInterface, interface{}) bool); ok {
		r0 = rf(ctx, key, items)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, djoemo.KeyInterface, interface{}) error); ok {
		r1 = rf(ctx, key, items)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: query, item
func (_m *QueryRepository) Query(query djoemo.QueryInterface, item interface{}) error {
	ret := _m.Called(query, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(djoemo.QueryInterface, interface{}) error); ok {
		r0 = rf(query, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryWithContext provides a mock function with given fields: ctx, query, item
func (_m *QueryRepository) QueryWithContext(ctx context.Context, query djoemo.QueryInterface, item interface{}) error {
	ret := _m.Called(ctx, query, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, djoemo.QueryInterface, interface{}) error); ok {
		r0 = rf(ctx, query, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
