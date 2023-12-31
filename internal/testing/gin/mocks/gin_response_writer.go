// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	bufio "bufio"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	net "net"
)

// ResponseWriter is an autogenerated mock type for the ResponseWriter type
type ResponseWriter struct {
	mock.Mock
}

type ResponseWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *ResponseWriter) EXPECT() *ResponseWriter_Expecter {
	return &ResponseWriter_Expecter{mock: &_m.Mock}
}

// CloseNotify provides a mock function with given fields:
func (_m *ResponseWriter) CloseNotify() <-chan bool {
	ret := _m.Called()

	var r0 <-chan bool
	if rf, ok := ret.Get(0).(func() <-chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan bool)
		}
	}

	return r0
}

// ResponseWriter_CloseNotify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseNotify'
type ResponseWriter_CloseNotify_Call struct {
	*mock.Call
}

// CloseNotify is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) CloseNotify() *ResponseWriter_CloseNotify_Call {
	return &ResponseWriter_CloseNotify_Call{Call: _e.mock.On("CloseNotify")}
}

func (_c *ResponseWriter_CloseNotify_Call) Run(run func()) *ResponseWriter_CloseNotify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_CloseNotify_Call) Return(_a0 <-chan bool) *ResponseWriter_CloseNotify_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_CloseNotify_Call) RunAndReturn(run func() <-chan bool) *ResponseWriter_CloseNotify_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields:
func (_m *ResponseWriter) Flush() {
	_m.Called()
}

// ResponseWriter_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type ResponseWriter_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Flush() *ResponseWriter_Flush_Call {
	return &ResponseWriter_Flush_Call{Call: _e.mock.On("Flush")}
}

func (_c *ResponseWriter_Flush_Call) Run(run func()) *ResponseWriter_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Flush_Call) Return() *ResponseWriter_Flush_Call {
	_c.Call.Return()
	return _c
}

func (_c *ResponseWriter_Flush_Call) RunAndReturn(run func()) *ResponseWriter_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *ResponseWriter) Header() http.Header {
	ret := _m.Called()

	var r0 http.Header
	if rf, ok := ret.Get(0).(func() http.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Header)
		}
	}

	return r0
}

// ResponseWriter_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ResponseWriter_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Header() *ResponseWriter_Header_Call {
	return &ResponseWriter_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *ResponseWriter_Header_Call) Run(run func()) *ResponseWriter_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Header_Call) Return(_a0 http.Header) *ResponseWriter_Header_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_Header_Call) RunAndReturn(run func() http.Header) *ResponseWriter_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Hijack provides a mock function with given fields:
func (_m *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	ret := _m.Called()

	var r0 net.Conn
	var r1 *bufio.ReadWriter
	var r2 error
	if rf, ok := ret.Get(0).(func() (net.Conn, *bufio.ReadWriter, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() net.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	if rf, ok := ret.Get(1).(func() *bufio.ReadWriter); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*bufio.ReadWriter)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResponseWriter_Hijack_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hijack'
type ResponseWriter_Hijack_Call struct {
	*mock.Call
}

// Hijack is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Hijack() *ResponseWriter_Hijack_Call {
	return &ResponseWriter_Hijack_Call{Call: _e.mock.On("Hijack")}
}

func (_c *ResponseWriter_Hijack_Call) Run(run func()) *ResponseWriter_Hijack_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Hijack_Call) Return(_a0 net.Conn, _a1 *bufio.ReadWriter, _a2 error) *ResponseWriter_Hijack_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ResponseWriter_Hijack_Call) RunAndReturn(run func() (net.Conn, *bufio.ReadWriter, error)) *ResponseWriter_Hijack_Call {
	_c.Call.Return(run)
	return _c
}

// Pusher provides a mock function with given fields:
func (_m *ResponseWriter) Pusher() http.Pusher {
	ret := _m.Called()

	var r0 http.Pusher
	if rf, ok := ret.Get(0).(func() http.Pusher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Pusher)
		}
	}

	return r0
}

// ResponseWriter_Pusher_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pusher'
type ResponseWriter_Pusher_Call struct {
	*mock.Call
}

// Pusher is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Pusher() *ResponseWriter_Pusher_Call {
	return &ResponseWriter_Pusher_Call{Call: _e.mock.On("Pusher")}
}

func (_c *ResponseWriter_Pusher_Call) Run(run func()) *ResponseWriter_Pusher_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Pusher_Call) Return(_a0 http.Pusher) *ResponseWriter_Pusher_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_Pusher_Call) RunAndReturn(run func() http.Pusher) *ResponseWriter_Pusher_Call {
	_c.Call.Return(run)
	return _c
}

// Size provides a mock function with given fields:
func (_m *ResponseWriter) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ResponseWriter_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type ResponseWriter_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Size() *ResponseWriter_Size_Call {
	return &ResponseWriter_Size_Call{Call: _e.mock.On("Size")}
}

func (_c *ResponseWriter_Size_Call) Run(run func()) *ResponseWriter_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Size_Call) Return(_a0 int) *ResponseWriter_Size_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_Size_Call) RunAndReturn(run func() int) *ResponseWriter_Size_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields:
func (_m *ResponseWriter) Status() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ResponseWriter_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type ResponseWriter_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Status() *ResponseWriter_Status_Call {
	return &ResponseWriter_Status_Call{Call: _e.mock.On("Status")}
}

func (_c *ResponseWriter_Status_Call) Run(run func()) *ResponseWriter_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Status_Call) Return(_a0 int) *ResponseWriter_Status_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_Status_Call) RunAndReturn(run func() int) *ResponseWriter_Status_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: _a0
func (_m *ResponseWriter) Write(_a0 []byte) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResponseWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type ResponseWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - _a0 []byte
func (_e *ResponseWriter_Expecter) Write(_a0 interface{}) *ResponseWriter_Write_Call {
	return &ResponseWriter_Write_Call{Call: _e.mock.On("Write", _a0)}
}

func (_c *ResponseWriter_Write_Call) Run(run func(_a0 []byte)) *ResponseWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ResponseWriter_Write_Call) Return(_a0 int, _a1 error) *ResponseWriter_Write_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResponseWriter_Write_Call) RunAndReturn(run func([]byte) (int, error)) *ResponseWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// WriteHeader provides a mock function with given fields: statusCode
func (_m *ResponseWriter) WriteHeader(statusCode int) {
	_m.Called(statusCode)
}

// ResponseWriter_WriteHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteHeader'
type ResponseWriter_WriteHeader_Call struct {
	*mock.Call
}

// WriteHeader is a helper method to define mock.On call
//   - statusCode int
func (_e *ResponseWriter_Expecter) WriteHeader(statusCode interface{}) *ResponseWriter_WriteHeader_Call {
	return &ResponseWriter_WriteHeader_Call{Call: _e.mock.On("WriteHeader", statusCode)}
}

func (_c *ResponseWriter_WriteHeader_Call) Run(run func(statusCode int)) *ResponseWriter_WriteHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ResponseWriter_WriteHeader_Call) Return() *ResponseWriter_WriteHeader_Call {
	_c.Call.Return()
	return _c
}

func (_c *ResponseWriter_WriteHeader_Call) RunAndReturn(run func(int)) *ResponseWriter_WriteHeader_Call {
	_c.Call.Return(run)
	return _c
}

// WriteHeaderNow provides a mock function with given fields:
func (_m *ResponseWriter) WriteHeaderNow() {
	_m.Called()
}

// ResponseWriter_WriteHeaderNow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteHeaderNow'
type ResponseWriter_WriteHeaderNow_Call struct {
	*mock.Call
}

// WriteHeaderNow is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) WriteHeaderNow() *ResponseWriter_WriteHeaderNow_Call {
	return &ResponseWriter_WriteHeaderNow_Call{Call: _e.mock.On("WriteHeaderNow")}
}

func (_c *ResponseWriter_WriteHeaderNow_Call) Run(run func()) *ResponseWriter_WriteHeaderNow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_WriteHeaderNow_Call) Return() *ResponseWriter_WriteHeaderNow_Call {
	_c.Call.Return()
	return _c
}

func (_c *ResponseWriter_WriteHeaderNow_Call) RunAndReturn(run func()) *ResponseWriter_WriteHeaderNow_Call {
	_c.Call.Return(run)
	return _c
}

// WriteString provides a mock function with given fields: _a0
func (_m *ResponseWriter) WriteString(_a0 string) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResponseWriter_WriteString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteString'
type ResponseWriter_WriteString_Call struct {
	*mock.Call
}

// WriteString is a helper method to define mock.On call
//   - _a0 string
func (_e *ResponseWriter_Expecter) WriteString(_a0 interface{}) *ResponseWriter_WriteString_Call {
	return &ResponseWriter_WriteString_Call{Call: _e.mock.On("WriteString", _a0)}
}

func (_c *ResponseWriter_WriteString_Call) Run(run func(_a0 string)) *ResponseWriter_WriteString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ResponseWriter_WriteString_Call) Return(_a0 int, _a1 error) *ResponseWriter_WriteString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResponseWriter_WriteString_Call) RunAndReturn(run func(string) (int, error)) *ResponseWriter_WriteString_Call {
	_c.Call.Return(run)
	return _c
}

// Written provides a mock function with given fields:
func (_m *ResponseWriter) Written() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResponseWriter_Written_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Written'
type ResponseWriter_Written_Call struct {
	*mock.Call
}

// Written is a helper method to define mock.On call
func (_e *ResponseWriter_Expecter) Written() *ResponseWriter_Written_Call {
	return &ResponseWriter_Written_Call{Call: _e.mock.On("Written")}
}

func (_c *ResponseWriter_Written_Call) Run(run func()) *ResponseWriter_Written_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ResponseWriter_Written_Call) Return(_a0 bool) *ResponseWriter_Written_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseWriter_Written_Call) RunAndReturn(run func() bool) *ResponseWriter_Written_Call {
	_c.Call.Return(run)
	return _c
}

// NewResponseWriter creates a new instance of ResponseWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponseWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResponseWriter {
	mock := &ResponseWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
