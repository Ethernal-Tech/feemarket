// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper is an autogenerated mock type for the AccountKeeper type
type AccountKeeper struct {
	mock.Mock
}

// GetAccount provides a mock function with given fields: ctx, addr
func (_m *AccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) authtypes.AccountI {
	ret := _m.Called(ctx, addr)

	var r0 authtypes.AccountI
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress) authtypes.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.AccountI)
		}
	}

	return r0
}

// GetModuleAccount provides a mock function with given fields: ctx, name
func (_m *AccountKeeper) GetModuleAccount(ctx types.Context, name string) authtypes.ModuleAccountI {
	ret := _m.Called(ctx, name)

	var r0 authtypes.ModuleAccountI
	if rf, ok := ret.Get(0).(func(types.Context, string) authtypes.ModuleAccountI); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.ModuleAccountI)
		}
	}

	return r0
}

// GetModuleAddress provides a mock function with given fields: moduleName
func (_m *AccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	ret := _m.Called(moduleName)

	var r0 types.AccAddress
	if rf, ok := ret.Get(0).(func(string) types.AccAddress); ok {
		r0 = rf(moduleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccAddress)
		}
	}

	return r0
}

// GetParams provides a mock function with given fields: ctx
func (_m *AccountKeeper) GetParams(ctx types.Context) authtypes.Params {
	ret := _m.Called(ctx)

	var r0 authtypes.Params
	if rf, ok := ret.Get(0).(func(types.Context) authtypes.Params); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(authtypes.Params)
	}

	return r0
}

// NewAccountWithAddress provides a mock function with given fields: ctx, addr
func (_m *AccountKeeper) NewAccountWithAddress(ctx types.Context, addr types.AccAddress) authtypes.AccountI {
	ret := _m.Called(ctx, addr)

	var r0 authtypes.AccountI
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress) authtypes.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.AccountI)
		}
	}

	return r0
}

// SetAccount provides a mock function with given fields: ctx, acc
func (_m *AccountKeeper) SetAccount(ctx types.Context, acc authtypes.AccountI) {
	_m.Called(ctx, acc)
}

// NewAccountKeeper creates a new instance of AccountKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountKeeper {
	mock := &AccountKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
