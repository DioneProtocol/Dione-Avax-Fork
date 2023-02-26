// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dioneprotocol/dionego/message (interfaces: OutboundMessage)

// Package message is a generated GoMock package.
package message

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOutboundMessage is a mock of OutboundMessage interface.
type MockOutboundMessage struct {
	ctrl     *gomock.Controller
	recorder *MockOutboundMessageMockRecorder
}

// MockOutboundMessageMockRecorder is the mock recorder for MockOutboundMessage.
type MockOutboundMessageMockRecorder struct {
	mock *MockOutboundMessage
}

// NewMockOutboundMessage creates a new mock instance.
func NewMockOutboundMessage(ctrl *gomock.Controller) *MockOutboundMessage {
	mock := &MockOutboundMessage{ctrl: ctrl}
	mock.recorder = &MockOutboundMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutboundMessage) EXPECT() *MockOutboundMessageMockRecorder {
	return m.recorder
}

// BypassThrottling mocks base method.
func (m *MockOutboundMessage) BypassThrottling() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BypassThrottling")
	ret0, _ := ret[0].(bool)
	return ret0
}

// BypassThrottling indicates an expected call of BypassThrottling.
func (mr *MockOutboundMessageMockRecorder) BypassThrottling() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BypassThrottling", reflect.TypeOf((*MockOutboundMessage)(nil).BypassThrottling))
}

// Bytes mocks base method.
func (m *MockOutboundMessage) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes.
func (mr *MockOutboundMessageMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockOutboundMessage)(nil).Bytes))
}

// BytesSavedCompression mocks base method.
func (m *MockOutboundMessage) BytesSavedCompression() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BytesSavedCompression")
	ret0, _ := ret[0].(int)
	return ret0
}

// BytesSavedCompression indicates an expected call of BytesSavedCompression.
func (mr *MockOutboundMessageMockRecorder) BytesSavedCompression() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BytesSavedCompression", reflect.TypeOf((*MockOutboundMessage)(nil).BytesSavedCompression))
}

// Op mocks base method.
func (m *MockOutboundMessage) Op() Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Op")
	ret0, _ := ret[0].(Op)
	return ret0
}

// Op indicates an expected call of Op.
func (mr *MockOutboundMessageMockRecorder) Op() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Op", reflect.TypeOf((*MockOutboundMessage)(nil).Op))
}
