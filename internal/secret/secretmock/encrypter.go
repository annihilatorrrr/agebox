// Code generated by mockery v2.5.1. DO NOT EDIT.

package secretmock

import (
	context "context"

	model "github.com/slok/agebox/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Encrypter is an autogenerated mock type for the Encrypter type
type Encrypter struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: ctx, _a1, key
func (_m *Encrypter) Decrypt(ctx context.Context, _a1 model.Secret, key model.PrivateKey) (*model.Secret, error) {
	ret := _m.Called(ctx, _a1, key)

	var r0 *model.Secret
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, model.PrivateKey) *model.Secret); ok {
		r0 = rf(ctx, _a1, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Secret, model.PrivateKey) error); ok {
		r1 = rf(ctx, _a1, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, _a1, keys
func (_m *Encrypter) Encrypt(ctx context.Context, _a1 model.Secret, keys []model.PublicKey) (*model.Secret, error) {
	ret := _m.Called(ctx, _a1, keys)

	var r0 *model.Secret
	if rf, ok := ret.Get(0).(func(context.Context, model.Secret, []model.PublicKey) *model.Secret); ok {
		r0 = rf(ctx, _a1, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Secret, []model.PublicKey) error); ok {
		r1 = rf(ctx, _a1, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
