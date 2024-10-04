// Code generated by mockery v2.46.2. DO NOT EDIT.

package processmock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IDProcessor is an autogenerated mock type for the IDProcessor type
type IDProcessor struct {
	mock.Mock
}

// ProcessID provides a mock function with given fields: ctx, secretID
func (_m *IDProcessor) ProcessID(ctx context.Context, secretID string) (string, error) {
	ret := _m.Called(ctx, secretID)

	if len(ret) == 0 {
		panic("no return value specified for ProcessID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, secretID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, secretID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, secretID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIDProcessor creates a new instance of IDProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDProcessor {
	mock := &IDProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
