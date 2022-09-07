// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libp2p/go-libp2p/core/connmgr (interfaces: ConnectionGater)

// Package libp2pwebtransport_test is a generated GoMock package.
package libp2pwebtransport_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	control "github.com/libp2p/go-libp2p/core/control"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// MockConnectionGater is a mock of ConnectionGater interface.
type MockConnectionGater struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionGaterMockRecorder
}

// MockConnectionGaterMockRecorder is the mock recorder for MockConnectionGater.
type MockConnectionGaterMockRecorder struct {
	mock *MockConnectionGater
}

// NewMockConnectionGater creates a new mock instance.
func NewMockConnectionGater(ctrl *gomock.Controller) *MockConnectionGater {
	mock := &MockConnectionGater{ctrl: ctrl}
	mock.recorder = &MockConnectionGaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionGater) EXPECT() *MockConnectionGaterMockRecorder {
	return m.recorder
}

// InterceptAccept mocks base method.
func (m *MockConnectionGater) InterceptAccept(arg0 network.ConnMultiaddrs) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptAccept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InterceptAccept indicates an expected call of InterceptAccept.
func (mr *MockConnectionGaterMockRecorder) InterceptAccept(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptAccept", reflect.TypeOf((*MockConnectionGater)(nil).InterceptAccept), arg0)
}

// InterceptAddrDial mocks base method.
func (m *MockConnectionGater) InterceptAddrDial(arg0 peer.ID, arg1 multiaddr.Multiaddr) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptAddrDial", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InterceptAddrDial indicates an expected call of InterceptAddrDial.
func (mr *MockConnectionGaterMockRecorder) InterceptAddrDial(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptAddrDial", reflect.TypeOf((*MockConnectionGater)(nil).InterceptAddrDial), arg0, arg1)
}

// InterceptPeerDial mocks base method.
func (m *MockConnectionGater) InterceptPeerDial(arg0 peer.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptPeerDial", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InterceptPeerDial indicates an expected call of InterceptPeerDial.
func (mr *MockConnectionGaterMockRecorder) InterceptPeerDial(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptPeerDial", reflect.TypeOf((*MockConnectionGater)(nil).InterceptPeerDial), arg0)
}

// InterceptSecured mocks base method.
func (m *MockConnectionGater) InterceptSecured(arg0 network.Direction, arg1 peer.ID, arg2 network.ConnMultiaddrs) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptSecured", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// InterceptSecured indicates an expected call of InterceptSecured.
func (mr *MockConnectionGaterMockRecorder) InterceptSecured(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptSecured", reflect.TypeOf((*MockConnectionGater)(nil).InterceptSecured), arg0, arg1, arg2)
}

// InterceptUpgraded mocks base method.
func (m *MockConnectionGater) InterceptUpgraded(arg0 network.Conn) (bool, control.DisconnectReason) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InterceptUpgraded", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(control.DisconnectReason)
	return ret0, ret1
}

// InterceptUpgraded indicates an expected call of InterceptUpgraded.
func (mr *MockConnectionGaterMockRecorder) InterceptUpgraded(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterceptUpgraded", reflect.TypeOf((*MockConnectionGater)(nil).InterceptUpgraded), arg0)
}
