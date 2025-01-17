// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/repository (interfaces: Orders)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/models"
)

// MockOrders is a mock of Orders interface.
type MockOrders struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersMockRecorder
}

// MockOrdersMockRecorder is the mock recorder for MockOrders.
type MockOrdersMockRecorder struct {
	mock *MockOrders
}

// NewMockOrders creates a new mock instance.
func NewMockOrders(ctrl *gomock.Controller) *MockOrders {
	mock := &MockOrders{ctrl: ctrl}
	mock.recorder = &MockOrdersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrders) EXPECT() *MockOrdersMockRecorder {
	return m.recorder
}

// ChangeStatusAndSum mocks base method.
func (m *MockOrders) ChangeStatusAndSum(arg0 float64, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatusAndSum", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeStatusAndSum indicates an expected call of ChangeStatusAndSum.
func (mr *MockOrdersMockRecorder) ChangeStatusAndSum(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatusAndSum", reflect.TypeOf((*MockOrders)(nil).ChangeStatusAndSum), arg0, arg1, arg2)
}

// CreateOrder mocks base method.
func (m *MockOrders) CreateOrder(arg0 int, arg1, arg2 string) (int, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrdersMockRecorder) CreateOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrders)(nil).CreateOrder), arg0, arg1, arg2)
}

// GetOrders mocks base method.
func (m *MockOrders) GetOrders(arg0 int) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrdersMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrders)(nil).GetOrders), arg0)
}

// GetOrdersWithStatus mocks base method.
func (m *MockOrders) GetOrdersWithStatus() ([]models.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersWithStatus")
	ret0, _ := ret[0].([]models.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersWithStatus indicates an expected call of GetOrdersWithStatus.
func (mr *MockOrdersMockRecorder) GetOrdersWithStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersWithStatus", reflect.TypeOf((*MockOrders)(nil).GetOrdersWithStatus))
}
