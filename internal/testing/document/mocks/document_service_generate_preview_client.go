// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	v1 "github.com/kinneko-de/api-contract/golang/kinnekode/restaurant/document/v1"
)

// DocumentService_GeneratePreviewClient is an autogenerated mock type for the DocumentService_GeneratePreviewClient type
type DocumentService_GeneratePreviewClient struct {
	mock.Mock
}

type DocumentService_GeneratePreviewClient_Expecter struct {
	mock *mock.Mock
}

func (_m *DocumentService_GeneratePreviewClient) EXPECT() *DocumentService_GeneratePreviewClient_Expecter {
	return &DocumentService_GeneratePreviewClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *DocumentService_GeneratePreviewClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DocumentService_GeneratePreviewClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type DocumentService_GeneratePreviewClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *DocumentService_GeneratePreviewClient_Expecter) CloseSend() *DocumentService_GeneratePreviewClient_CloseSend_Call {
	return &DocumentService_GeneratePreviewClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *DocumentService_GeneratePreviewClient_CloseSend_Call) Run(run func()) *DocumentService_GeneratePreviewClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_CloseSend_Call) Return(_a0 error) *DocumentService_GeneratePreviewClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_CloseSend_Call) RunAndReturn(run func() error) *DocumentService_GeneratePreviewClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *DocumentService_GeneratePreviewClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// DocumentService_GeneratePreviewClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type DocumentService_GeneratePreviewClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *DocumentService_GeneratePreviewClient_Expecter) Context() *DocumentService_GeneratePreviewClient_Context_Call {
	return &DocumentService_GeneratePreviewClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *DocumentService_GeneratePreviewClient_Context_Call) Run(run func()) *DocumentService_GeneratePreviewClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Context_Call) Return(_a0 context.Context) *DocumentService_GeneratePreviewClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Context_Call) RunAndReturn(run func() context.Context) *DocumentService_GeneratePreviewClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *DocumentService_GeneratePreviewClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DocumentService_GeneratePreviewClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type DocumentService_GeneratePreviewClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *DocumentService_GeneratePreviewClient_Expecter) Header() *DocumentService_GeneratePreviewClient_Header_Call {
	return &DocumentService_GeneratePreviewClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *DocumentService_GeneratePreviewClient_Header_Call) Run(run func()) *DocumentService_GeneratePreviewClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *DocumentService_GeneratePreviewClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *DocumentService_GeneratePreviewClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *DocumentService_GeneratePreviewClient) Recv() (*v1.GeneratePreviewResponse, error) {
	ret := _m.Called()

	var r0 *v1.GeneratePreviewResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.GeneratePreviewResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.GeneratePreviewResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.GeneratePreviewResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DocumentService_GeneratePreviewClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type DocumentService_GeneratePreviewClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *DocumentService_GeneratePreviewClient_Expecter) Recv() *DocumentService_GeneratePreviewClient_Recv_Call {
	return &DocumentService_GeneratePreviewClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *DocumentService_GeneratePreviewClient_Recv_Call) Run(run func()) *DocumentService_GeneratePreviewClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Recv_Call) Return(_a0 *v1.GeneratePreviewResponse, _a1 error) *DocumentService_GeneratePreviewClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Recv_Call) RunAndReturn(run func() (*v1.GeneratePreviewResponse, error)) *DocumentService_GeneratePreviewClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *DocumentService_GeneratePreviewClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DocumentService_GeneratePreviewClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type DocumentService_GeneratePreviewClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *DocumentService_GeneratePreviewClient_Expecter) RecvMsg(m interface{}) *DocumentService_GeneratePreviewClient_RecvMsg_Call {
	return &DocumentService_GeneratePreviewClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *DocumentService_GeneratePreviewClient_RecvMsg_Call) Run(run func(m interface{})) *DocumentService_GeneratePreviewClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_RecvMsg_Call) Return(_a0 error) *DocumentService_GeneratePreviewClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *DocumentService_GeneratePreviewClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *DocumentService_GeneratePreviewClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DocumentService_GeneratePreviewClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type DocumentService_GeneratePreviewClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *DocumentService_GeneratePreviewClient_Expecter) SendMsg(m interface{}) *DocumentService_GeneratePreviewClient_SendMsg_Call {
	return &DocumentService_GeneratePreviewClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *DocumentService_GeneratePreviewClient_SendMsg_Call) Run(run func(m interface{})) *DocumentService_GeneratePreviewClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_SendMsg_Call) Return(_a0 error) *DocumentService_GeneratePreviewClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *DocumentService_GeneratePreviewClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *DocumentService_GeneratePreviewClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// DocumentService_GeneratePreviewClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type DocumentService_GeneratePreviewClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *DocumentService_GeneratePreviewClient_Expecter) Trailer() *DocumentService_GeneratePreviewClient_Trailer_Call {
	return &DocumentService_GeneratePreviewClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *DocumentService_GeneratePreviewClient_Trailer_Call) Run(run func()) *DocumentService_GeneratePreviewClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Trailer_Call) Return(_a0 metadata.MD) *DocumentService_GeneratePreviewClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DocumentService_GeneratePreviewClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *DocumentService_GeneratePreviewClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewDocumentService_GeneratePreviewClient creates a new instance of DocumentService_GeneratePreviewClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocumentService_GeneratePreviewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DocumentService_GeneratePreviewClient {
	mock := &DocumentService_GeneratePreviewClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
