// Code generated by MockGen. DO NOT EDIT.
// Source: internal/lsp/handler.go
//
// Generated by this command:
//
//	mockgen -source internal/lsp/handler.go -destination internal/lsp/testutil/handler.go -package testutil
//

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/armsnyder/openapi-language-server/internal/lsp/types"
	gomock "go.uber.org/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockHandler) Capabilities() types.ServerCapabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(types.ServerCapabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockHandlerMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockHandler)(nil).Capabilities))
}

// HandleChange mocks base method.
func (m *MockHandler) HandleChange(params types.DidChangeTextDocumentParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleChange", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleChange indicates an expected call of HandleChange.
func (mr *MockHandlerMockRecorder) HandleChange(params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleChange", reflect.TypeOf((*MockHandler)(nil).HandleChange), params)
}

// HandleClose mocks base method.
func (m *MockHandler) HandleClose(params types.DidCloseTextDocumentParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleClose", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleClose indicates an expected call of HandleClose.
func (mr *MockHandlerMockRecorder) HandleClose(params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleClose", reflect.TypeOf((*MockHandler)(nil).HandleClose), params)
}

// HandleDefinition mocks base method.
func (m *MockHandler) HandleDefinition(params types.DefinitionParams) ([]types.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDefinition", params)
	ret0, _ := ret[0].([]types.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleDefinition indicates an expected call of HandleDefinition.
func (mr *MockHandlerMockRecorder) HandleDefinition(params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDefinition", reflect.TypeOf((*MockHandler)(nil).HandleDefinition), params)
}

// HandleOpen mocks base method.
func (m *MockHandler) HandleOpen(params types.DidOpenTextDocumentParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleOpen", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleOpen indicates an expected call of HandleOpen.
func (mr *MockHandlerMockRecorder) HandleOpen(params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleOpen", reflect.TypeOf((*MockHandler)(nil).HandleOpen), params)
}

// HandleReferences mocks base method.
func (m *MockHandler) HandleReferences(params types.ReferenceParams) ([]types.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleReferences", params)
	ret0, _ := ret[0].([]types.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleReferences indicates an expected call of HandleReferences.
func (mr *MockHandlerMockRecorder) HandleReferences(params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleReferences", reflect.TypeOf((*MockHandler)(nil).HandleReferences), params)
}
