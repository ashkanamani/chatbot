// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package repository

import (
	"context"

	"github.com/ashkanamani/chatbot/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// NewMockCommonBehaviour creates a new instance of MockCommonBehaviour. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommonBehaviour[T entity.Entity](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommonBehaviour[T] {
	mock := &MockCommonBehaviour[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockCommonBehaviour is an autogenerated mock type for the CommonBehaviour type
type MockCommonBehaviour[T entity.Entity] struct {
	mock.Mock
}

type MockCommonBehaviour_Expecter[T entity.Entity] struct {
	mock *mock.Mock
}

func (_m *MockCommonBehaviour[T]) EXPECT() *MockCommonBehaviour_Expecter[T] {
	return &MockCommonBehaviour_Expecter[T]{mock: &_m.Mock}
}

// Get provides a mock function for the type MockCommonBehaviour
func (_mock *MockCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	ret := _mock.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 T
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, entity.ID) (T, error)); ok {
		return returnFunc(ctx, id)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, entity.ID) T); ok {
		r0 = returnFunc(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(T)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, entity.ID) error); ok {
		r1 = returnFunc(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockCommonBehaviour_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCommonBehaviour_Get_Call[T entity.Entity] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id entity.ID
func (_e *MockCommonBehaviour_Expecter[T]) Get(ctx interface{}, id interface{}) *MockCommonBehaviour_Get_Call[T] {
	return &MockCommonBehaviour_Get_Call[T]{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockCommonBehaviour_Get_Call[T]) Run(run func(ctx context.Context, id entity.ID)) *MockCommonBehaviour_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 entity.ID
		if args[1] != nil {
			arg1 = args[1].(entity.ID)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockCommonBehaviour_Get_Call[T]) Return(v T, err error) *MockCommonBehaviour_Get_Call[T] {
	_c.Call.Return(v, err)
	return _c
}

func (_c *MockCommonBehaviour_Get_Call[T]) RunAndReturn(run func(ctx context.Context, id entity.ID) (T, error)) *MockCommonBehaviour_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function for the type MockCommonBehaviour
func (_mock *MockCommonBehaviour[T]) Save(ctx context.Context, entity1 T) error {
	ret := _mock.Called(ctx, entity1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, T) error); ok {
		r0 = returnFunc(ctx, entity1)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockCommonBehaviour_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockCommonBehaviour_Save_Call[T entity.Entity] struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - entity1 T
func (_e *MockCommonBehaviour_Expecter[T]) Save(ctx interface{}, entity1 interface{}) *MockCommonBehaviour_Save_Call[T] {
	return &MockCommonBehaviour_Save_Call[T]{Call: _e.mock.On("Save", ctx, entity1)}
}

func (_c *MockCommonBehaviour_Save_Call[T]) Run(run func(ctx context.Context, entity1 T)) *MockCommonBehaviour_Save_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 T
		if args[1] != nil {
			arg1 = args[1].(T)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockCommonBehaviour_Save_Call[T]) Return(err error) *MockCommonBehaviour_Save_Call[T] {
	_c.Call.Return(err)
	return _c
}

func (_c *MockCommonBehaviour_Save_Call[T]) RunAndReturn(run func(ctx context.Context, entity1 T) error) *MockCommonBehaviour_Save_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountRepository creates a new instance of MockAccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountRepository {
	mock := &MockAccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockAccountRepository is an autogenerated mock type for the AccountRepository type
type MockAccountRepository struct {
	mock.Mock
}

type MockAccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepository) EXPECT() *MockAccountRepository_Expecter {
	return &MockAccountRepository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function for the type MockAccountRepository
func (_mock *MockAccountRepository) Get(ctx context.Context, id entity.ID) (entity.Account, error) {
	ret := _mock.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.Account
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, entity.ID) (entity.Account, error)); ok {
		return returnFunc(ctx, id)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, entity.ID) entity.Account); ok {
		r0 = returnFunc(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, entity.ID) error); ok {
		r1 = returnFunc(ctx, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockAccountRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - id entity.ID
func (_e *MockAccountRepository_Expecter) Get(ctx interface{}, id interface{}) *MockAccountRepository_Get_Call {
	return &MockAccountRepository_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *MockAccountRepository_Get_Call) Run(run func(ctx context.Context, id entity.ID)) *MockAccountRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 entity.ID
		if args[1] != nil {
			arg1 = args[1].(entity.ID)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockAccountRepository_Get_Call) Return(account entity.Account, err error) *MockAccountRepository_Get_Call {
	_c.Call.Return(account, err)
	return _c
}

func (_c *MockAccountRepository_Get_Call) RunAndReturn(run func(ctx context.Context, id entity.ID) (entity.Account, error)) *MockAccountRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function for the type MockAccountRepository
func (_mock *MockAccountRepository) Save(ctx context.Context, entity1 entity.Account) error {
	ret := _mock.Called(ctx, entity1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, entity.Account) error); ok {
		r0 = returnFunc(ctx, entity1)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// MockAccountRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockAccountRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - entity1 entity.Account
func (_e *MockAccountRepository_Expecter) Save(ctx interface{}, entity1 interface{}) *MockAccountRepository_Save_Call {
	return &MockAccountRepository_Save_Call{Call: _e.mock.On("Save", ctx, entity1)}
}

func (_c *MockAccountRepository_Save_Call) Run(run func(ctx context.Context, entity1 entity.Account)) *MockAccountRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 entity.Account
		if args[1] != nil {
			arg1 = args[1].(entity.Account)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockAccountRepository_Save_Call) Return(err error) *MockAccountRepository_Save_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockAccountRepository_Save_Call) RunAndReturn(run func(ctx context.Context, entity1 entity.Account) error) *MockAccountRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}
