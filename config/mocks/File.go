// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *File) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *File) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *File) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAt provides a mock function with given fields: p, off
func (_m *File) ReadAt(p []byte, off int64) (int, error) {
	ret := _m.Called(p, off)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte, int64) int); ok {
		r0 = rf(p, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, int64) error); ok {
		r1 = rf(p, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readdir provides a mock function with given fields: count
func (_m *File) Readdir(count int) ([]os.FileInfo, error) {
	ret := _m.Called(count)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(int) []os.FileInfo); ok {
		r0 = rf(count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readdirnames provides a mock function with given fields: n
func (_m *File) Readdirnames(n int) ([]string, error) {
	ret := _m.Called(n)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int) []string); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Seek provides a mock function with given fields: offset, whence
func (_m *File) Seek(offset int64, whence int) (int64, error) {
	ret := _m.Called(offset, whence)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int) int64); ok {
		r0 = rf(offset, whence)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(offset, whence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields:
func (_m *File) Stat() (os.FileInfo, error) {
	ret := _m.Called()

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func() os.FileInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sync provides a mock function with given fields:
func (_m *File) Sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Truncate provides a mock function with given fields: size
func (_m *File) Truncate(size int64) error {
	ret := _m.Called(size)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *File) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteAt provides a mock function with given fields: p, off
func (_m *File) WriteAt(p []byte, off int64) (int, error) {
	ret := _m.Called(p, off)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte, int64) int); ok {
		r0 = rf(p, off)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, int64) error); ok {
		r1 = rf(p, off)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteString provides a mock function with given fields: s
func (_m *File) WriteString(s string) (int, error) {
	ret := _m.Called(s)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
