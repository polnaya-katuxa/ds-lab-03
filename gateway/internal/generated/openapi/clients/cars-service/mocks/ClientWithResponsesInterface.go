// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cars_service "github.com/polnaya-katuxa/ds-lab-02/gateway/internal/generated/openapi/clients/cars-service"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ClientWithResponsesInterface is an autogenerated mock type for the ClientWithResponsesInterface type
type ClientWithResponsesInterface struct {
	mock.Mock
}

type ClientWithResponsesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientWithResponsesInterface) EXPECT() *ClientWithResponsesInterface_Expecter {
	return &ClientWithResponsesInterface_Expecter{mock: &_m.Mock}
}

// BookWithResponse provides a mock function with given fields: ctx, carUid, reqEditors
func (_m *ClientWithResponsesInterface) BookWithResponse(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn) (*cars_service.BookResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, carUid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BookWithResponse")
	}

	var r0 *cars_service.BookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.BookResponse, error)); ok {
		return rf(ctx, carUid, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) *cars_service.BookResponse); ok {
		r0 = rf(ctx, carUid, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cars_service.BookResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) error); ok {
		r1 = rf(ctx, carUid, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_BookWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BookWithResponse'
type ClientWithResponsesInterface_BookWithResponse_Call struct {
	*mock.Call
}

// BookWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - carUid uuid.UUID
//   - reqEditors ...cars_service.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) BookWithResponse(ctx interface{}, carUid interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_BookWithResponse_Call {
	return &ClientWithResponsesInterface_BookWithResponse_Call{Call: _e.mock.On("BookWithResponse",
		append([]interface{}{ctx, carUid}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_BookWithResponse_Call) Run(run func(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn)) *ClientWithResponsesInterface_BookWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]cars_service.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(cars_service.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(uuid.UUID), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_BookWithResponse_Call) Return(_a0 *cars_service.BookResponse, _a1 error) *ClientWithResponsesInterface_BookWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_BookWithResponse_Call) RunAndReturn(run func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.BookResponse, error)) *ClientWithResponsesInterface_BookWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// GetWithResponse provides a mock function with given fields: ctx, carUid, reqEditors
func (_m *ClientWithResponsesInterface) GetWithResponse(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn) (*cars_service.GetResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, carUid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetWithResponse")
	}

	var r0 *cars_service.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.GetResponse, error)); ok {
		return rf(ctx, carUid, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) *cars_service.GetResponse); ok {
		r0 = rf(ctx, carUid, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cars_service.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) error); ok {
		r1 = rf(ctx, carUid, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_GetWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithResponse'
type ClientWithResponsesInterface_GetWithResponse_Call struct {
	*mock.Call
}

// GetWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - carUid uuid.UUID
//   - reqEditors ...cars_service.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) GetWithResponse(ctx interface{}, carUid interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_GetWithResponse_Call {
	return &ClientWithResponsesInterface_GetWithResponse_Call{Call: _e.mock.On("GetWithResponse",
		append([]interface{}{ctx, carUid}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_GetWithResponse_Call) Run(run func(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn)) *ClientWithResponsesInterface_GetWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]cars_service.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(cars_service.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(uuid.UUID), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_GetWithResponse_Call) Return(_a0 *cars_service.GetResponse, _a1 error) *ClientWithResponsesInterface_GetWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_GetWithResponse_Call) RunAndReturn(run func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.GetResponse, error)) *ClientWithResponsesInterface_GetWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// ListWithResponse provides a mock function with given fields: ctx, params, reqEditors
func (_m *ClientWithResponsesInterface) ListWithResponse(ctx context.Context, params *cars_service.ListParams, reqEditors ...cars_service.RequestEditorFn) (*cars_service.ListResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListWithResponse")
	}

	var r0 *cars_service.ListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cars_service.ListParams, ...cars_service.RequestEditorFn) (*cars_service.ListResponse, error)); ok {
		return rf(ctx, params, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cars_service.ListParams, ...cars_service.RequestEditorFn) *cars_service.ListResponse); ok {
		r0 = rf(ctx, params, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cars_service.ListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cars_service.ListParams, ...cars_service.RequestEditorFn) error); ok {
		r1 = rf(ctx, params, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_ListWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListWithResponse'
type ClientWithResponsesInterface_ListWithResponse_Call struct {
	*mock.Call
}

// ListWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - params *cars_service.ListParams
//   - reqEditors ...cars_service.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) ListWithResponse(ctx interface{}, params interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_ListWithResponse_Call {
	return &ClientWithResponsesInterface_ListWithResponse_Call{Call: _e.mock.On("ListWithResponse",
		append([]interface{}{ctx, params}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_ListWithResponse_Call) Run(run func(ctx context.Context, params *cars_service.ListParams, reqEditors ...cars_service.RequestEditorFn)) *ClientWithResponsesInterface_ListWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]cars_service.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(cars_service.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(*cars_service.ListParams), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_ListWithResponse_Call) Return(_a0 *cars_service.ListResponse, _a1 error) *ClientWithResponsesInterface_ListWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_ListWithResponse_Call) RunAndReturn(run func(context.Context, *cars_service.ListParams, ...cars_service.RequestEditorFn) (*cars_service.ListResponse, error)) *ClientWithResponsesInterface_ListWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// LiveWithResponse provides a mock function with given fields: ctx, reqEditors
func (_m *ClientWithResponsesInterface) LiveWithResponse(ctx context.Context, reqEditors ...cars_service.RequestEditorFn) (*cars_service.LiveResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for LiveWithResponse")
	}

	var r0 *cars_service.LiveResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...cars_service.RequestEditorFn) (*cars_service.LiveResponse, error)); ok {
		return rf(ctx, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...cars_service.RequestEditorFn) *cars_service.LiveResponse); ok {
		r0 = rf(ctx, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cars_service.LiveResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...cars_service.RequestEditorFn) error); ok {
		r1 = rf(ctx, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_LiveWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LiveWithResponse'
type ClientWithResponsesInterface_LiveWithResponse_Call struct {
	*mock.Call
}

// LiveWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - reqEditors ...cars_service.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) LiveWithResponse(ctx interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_LiveWithResponse_Call {
	return &ClientWithResponsesInterface_LiveWithResponse_Call{Call: _e.mock.On("LiveWithResponse",
		append([]interface{}{ctx}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_LiveWithResponse_Call) Run(run func(ctx context.Context, reqEditors ...cars_service.RequestEditorFn)) *ClientWithResponsesInterface_LiveWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]cars_service.RequestEditorFn, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(cars_service.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_LiveWithResponse_Call) Return(_a0 *cars_service.LiveResponse, _a1 error) *ClientWithResponsesInterface_LiveWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_LiveWithResponse_Call) RunAndReturn(run func(context.Context, ...cars_service.RequestEditorFn) (*cars_service.LiveResponse, error)) *ClientWithResponsesInterface_LiveWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// UnbookWithResponse provides a mock function with given fields: ctx, carUid, reqEditors
func (_m *ClientWithResponsesInterface) UnbookWithResponse(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn) (*cars_service.UnbookResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, carUid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnbookWithResponse")
	}

	var r0 *cars_service.UnbookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.UnbookResponse, error)); ok {
		return rf(ctx, carUid, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) *cars_service.UnbookResponse); ok {
		r0 = rf(ctx, carUid, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cars_service.UnbookResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) error); ok {
		r1 = rf(ctx, carUid, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_UnbookWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnbookWithResponse'
type ClientWithResponsesInterface_UnbookWithResponse_Call struct {
	*mock.Call
}

// UnbookWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - carUid uuid.UUID
//   - reqEditors ...cars_service.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) UnbookWithResponse(ctx interface{}, carUid interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_UnbookWithResponse_Call {
	return &ClientWithResponsesInterface_UnbookWithResponse_Call{Call: _e.mock.On("UnbookWithResponse",
		append([]interface{}{ctx, carUid}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_UnbookWithResponse_Call) Run(run func(ctx context.Context, carUid uuid.UUID, reqEditors ...cars_service.RequestEditorFn)) *ClientWithResponsesInterface_UnbookWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]cars_service.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(cars_service.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(uuid.UUID), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_UnbookWithResponse_Call) Return(_a0 *cars_service.UnbookResponse, _a1 error) *ClientWithResponsesInterface_UnbookWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_UnbookWithResponse_Call) RunAndReturn(run func(context.Context, uuid.UUID, ...cars_service.RequestEditorFn) (*cars_service.UnbookResponse, error)) *ClientWithResponsesInterface_UnbookWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientWithResponsesInterface creates a new instance of ClientWithResponsesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientWithResponsesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientWithResponsesInterface {
	mock := &ClientWithResponsesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}