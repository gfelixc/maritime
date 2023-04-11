// Code generated by mockery v2.20.0. DO NOT EDIT.

package internal

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) Save(_a0 context.Context, _a1 *Port) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Port) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *Port
func (_e *MockRepository_Expecter) Save(_a0 interface{}, _a1 interface{}) *MockRepository_Save_Call {
	return &MockRepository_Save_Call{Call: _e.mock.On("Save", _a0, _a1)}
}

func (_c *MockRepository_Save_Call) Run(run func(_a0 context.Context, _a1 *Port)) *MockRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Port))
	})
	return _c
}

func (_c *MockRepository_Save_Call) Return(_a0 error) *MockRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Save_Call) RunAndReturn(run func(context.Context, *Port) error) *MockRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}