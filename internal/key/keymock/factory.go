// Code generated by mockery v2.46.2. DO NOT EDIT.

package keymock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/slok/agebox/internal/model"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// GetPrivateKey provides a mock function with given fields: ctx, data
func (_m *Factory) GetPrivateKey(ctx context.Context, data []byte) (model.PrivateKey, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for GetPrivateKey")
	}

	var r0 model.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (model.PrivateKey, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) model.PrivateKey); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicKey provides a mock function with given fields: ctx, data
func (_m *Factory) GetPublicKey(ctx context.Context, data []byte) (model.PublicKey, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for GetPublicKey")
	}

	var r0 model.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (model.PublicKey, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) model.PublicKey); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
