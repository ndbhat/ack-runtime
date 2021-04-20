// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	manager "sigs.k8s.io/controller-runtime/pkg/manager"

	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1 "k8s.io/api/core/v1"
)

// Reconciler is an autogenerated mock type for the Reconciler type
type Reconciler struct {
	mock.Mock
}

// BindControllerManager provides a mock function with given fields: _a0
func (_m *Reconciler) BindControllerManager(_a0 manager.Manager) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(manager.Manager) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reconcile provides a mock function with given fields: _a0
func (_m *Reconciler) Reconcile(_a0 reconcile.Request) (reconcile.Result, error) {
	ret := _m.Called(_a0)

	var r0 reconcile.Result
	if rf, ok := ret.Get(0).(func(reconcile.Request) reconcile.Result); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(reconcile.Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecretValueFromReference provides a mock function with given fields: _a0
func (_m *Reconciler) SecretValueFromReference(_a0 *v1.SecretReference) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*v1.SecretReference) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.SecretReference) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}