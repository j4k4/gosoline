// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// AllKeys provides a mock function with given fields:
func (_m *Config) AllKeys() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// AugmentString provides a mock function with given fields: str
func (_m *Config) AugmentString(str string) string {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Bind provides a mock function with given fields: obj
func (_m *Config) Bind(obj interface{}) {
	_m.Called(obj)
}

// Get provides a mock function with given fields: _a0
func (_m *Config) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBool provides a mock function with given fields: key
func (_m *Config) GetBool(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetDuration provides a mock function with given fields: _a0
func (_m *Config) GetDuration(_a0 string) time.Duration {
	ret := _m.Called(_a0)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(string) time.Duration); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// GetFloat64 provides a mock function with given fields: _a0
func (_m *Config) GetFloat64(_a0 string) float64 {
	ret := _m.Called(_a0)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetInt provides a mock function with given fields: _a0
func (_m *Config) GetInt(_a0 string) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetString provides a mock function with given fields: _a0
func (_m *Config) GetString(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStringMapString provides a mock function with given fields: key
func (_m *Config) GetStringMapString(key string) map[string]string {
	ret := _m.Called(key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetStringSlice provides a mock function with given fields: key
func (_m *Config) GetStringSlice(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// IsSet provides a mock function with given fields: _a0
func (_m *Config) IsSet(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Unmarshal provides a mock function with given fields: key, val
func (_m *Config) Unmarshal(key string, val interface{}) {
	_m.Called(key, val)
}
