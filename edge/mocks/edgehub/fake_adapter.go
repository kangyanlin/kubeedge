// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeedge/kubeedge/pkg/edgehub/clients (interfaces: Adapter)

// Package edgehub is a generated GoMock package.
package edgehub

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/kubeedge/beehive/pkg/core/model"
)

// MockAdapter is a mock of Adapter interface
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockAdapter) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockAdapterMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAdapter)(nil).Init))
}

// Notify mocks base method
func (m *MockAdapter) Notify(arg0 map[string]string) {
	m.ctrl.Call(m, "Notify", arg0)
}

// Notify indicates an expected call of Notify
func (mr *MockAdapterMockRecorder) Notify(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockAdapter)(nil).Notify), arg0)
}

// Receive mocks base method
func (m *MockAdapter) Receive() (model.Message, error) {
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive
func (mr *MockAdapterMockRecorder) Receive() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockAdapter)(nil).Receive))
}

// Send mocks base method
func (m *MockAdapter) Send(arg0 model.Message) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockAdapterMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAdapter)(nil).Send), arg0)
}

// Uninit mocks base method
func (m *MockAdapter) Uninit() {
	m.ctrl.Call(m, "Uninit")
}

// Uninit indicates an expected call of Uninit
func (mr *MockAdapterMockRecorder) Uninit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninit", reflect.TypeOf((*MockAdapter)(nil).Uninit))
}
