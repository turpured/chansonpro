// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	os "os"
	time "time"

	afero "github.com/spf13/afero"
	mock "github.com/stretchr/testify/mock"
)

// Fs is an autogenerated mock type for the Fs type
type Fs struct {
	mock.Mock
}

// Chmod provides a mock function with given fields: name, mode
func (_m *Fs) Chmod(name string, mode os.FileMode) error {
	ret := _m.Called(name, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, os.FileMode) error); ok {
		r0 = rf(name, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Chtimes provides a mock function with given fields: name, atime, mtime
func (_m *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	ret := _m.Called(name, atime, mtime)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) error); ok {
		r0 = rf(name, atime, mtime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: name
func (_m *Fs) Create(name string) (afero.File, error) {
	ret := _m.Called(name)

	var r0 afero.File
	if rf, ok := ret.Get(0).(func(string) afero.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(afero.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mkdir provides a mock function with given fields: name, perm
func (_m *Fs) Mkdir(name string, perm os.FileMode) error {
	ret := _m.Called(name, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, os.FileMode) error); ok {
		r0 = rf(name, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MkdirAll provides a mock function with given fields: path, perm
func (_m *Fs) MkdirAll(path string, perm os.FileMode) error {
	ret := _m.Called(path, perm)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, os.FileMode) error); ok {
		r0 = rf(path, perm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Fs) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Open provides a mock function with given fields: name
func (_m *Fs) Open(name string) (afero.File, error) {
	ret := _m.Called(name)

	var r0 afero.File
	if rf, ok := ret.Get(0).(func(string) afero.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(afero.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenFile provides a mock function with given fields: name, flag, perm
func (_m *Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	ret := _m.Called(name, flag, perm)

	var r0 afero.File
	if rf, ok := ret.Get(0).(func(string, int, os.FileMode) afero.File); ok {
		r0 = rf(name, flag, perm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(afero.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, os.FileMode) error); ok {
		r1 = rf(name, flag, perm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: name
func (_m *Fs) Remove(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: path
func (_m *Fs) RemoveAll(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rename provides a mock function with given fields: oldname, newname
func (_m *Fs) Rename(oldname string, newname string) error {
	ret := _m.Called(oldname, newname)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(oldname, newname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stat provides a mock function with given fields: name
func (_m *Fs) Stat(name string) (os.FileInfo, error) {
	ret := _m.Called(name)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(string) os.FileInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
