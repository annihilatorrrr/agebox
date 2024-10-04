// Code generated by mockery v2.46.2. DO NOT EDIT.

package encryptmock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/slok/agebox/internal/model"
)

// Encrypter is an autogenerated mock type for the Encrypter type
type Encrypter struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: ctx, secret, keys
func (_m *Encrypter) Decrypt(ctx context.Context, secret model.Secret, keys []model.PrivateKey) (*model.Secret, error) {
	ret := _m.Called(ctx, secret, keys)

	if len(ret) == 0 {
		panic("no return value specified for Decrypt")
	}

	var r0 *model.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, []model.PrivateKey) (*model.Secret, error)); ok {
		return rf(ctx, secret, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, []model.PrivateKey) *model.Secret); ok {
		r0 = rf(ctx, secret, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Secret, []model.PrivateKey) error); ok {
		r1 = rf(ctx, secret, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, secret, keys
func (_m *Encrypter) Encrypt(ctx context.Context, secret model.Secret, keys []model.PublicKey) (*model.Secret, error) {
	ret := _m.Called(ctx, secret, keys)

	if len(ret) == 0 {
		panic("no return value specified for Encrypt")
	}

	var r0 *model.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, []model.PublicKey) (*model.Secret, error)); ok {
		return rf(ctx, secret, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, []model.PublicKey) *model.Secret); ok {
		r0 = rf(ctx, secret, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Secret, []model.PublicKey) error); ok {
		r1 = rf(ctx, secret, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEncrypter creates a new instance of Encrypter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEncrypter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Encrypter {
	mock := &Encrypter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
