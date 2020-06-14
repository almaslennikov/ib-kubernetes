// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import runtime "k8s.io/apimachinery/pkg/runtime"
import utils "github.com/Mellanox/ib-kubernetes/pkg/utils"

// ResourceEventHandler is an autogenerated mock type for the ResourceEventHandler type
type ResourceEventHandler struct {
	mock.Mock
}

// GetResourceObject provides a mock function with given fields:
func (_m *ResourceEventHandler) GetResourceObject() runtime.Object {
	ret := _m.Called()

	var r0 runtime.Object
	if rf, ok := ret.Get(0).(func() runtime.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Object)
		}
	}

	return r0
}

// GetResults provides a mock function with given fields:
func (_m *ResourceEventHandler) GetResults() (*utils.SynchronizedMap, *utils.SynchronizedMap) {
	ret := _m.Called()

	var r0 *utils.SynchronizedMap
	if rf, ok := ret.Get(0).(func() *utils.SynchronizedMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utils.SynchronizedMap)
		}
	}

	var r1 *utils.SynchronizedMap
	if rf, ok := ret.Get(1).(func() *utils.SynchronizedMap); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*utils.SynchronizedMap)
		}
	}

	return r0, r1
}

// OnAdd provides a mock function with given fields: obj
func (_m *ResourceEventHandler) OnAdd(obj interface{}) {
	_m.Called(obj)
}

// OnDelete provides a mock function with given fields: obj
func (_m *ResourceEventHandler) OnDelete(obj interface{}) {
	_m.Called(obj)
}

// OnUpdate provides a mock function with given fields: oldObj, newObj
func (_m *ResourceEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	_m.Called(oldObj, newObj)
}