package mocks

import (
	mock "github.com/stretchr/testify/mock"
	common "main/common"
)

type Controller struct {
	mock.Mock
}

func (_m *Controller) Routes() []common.Route {
	ret := _m.Called()

	var r0 []common.Route
	if rf, ok := ret.Get(0).(func() []common.Route); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Route)
		}
	}

	return r0
}
