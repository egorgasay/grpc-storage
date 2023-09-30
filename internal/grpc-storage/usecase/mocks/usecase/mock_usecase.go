// Code generated by MockGen. DO NOT EDIT.
// Source: itisadb/internal/grpc-storage/core (interfaces: IUseCase)

// Package mocks is a generated GoMock package.
package usecase_mocks

import (
	usecase "itisadb/internal/grpc-storage/usecase"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUseCase is a mock of IUseCase interface.
type MockIUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIUseCaseMockRecorder
}

// MockIUseCaseMockRecorder is the mock recorder for MockIUseCase.
type MockIUseCaseMockRecorder struct {
	mock *MockIUseCase
}

// NewMockIUseCase creates a new mock instance.
func NewMockIUseCase(ctrl *gomock.Controller) *MockIUseCase {
	mock := &MockIUseCase{ctrl: ctrl}
	mock.recorder = &MockIUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUseCase) EXPECT() *MockIUseCaseMockRecorder {
	return m.recorder
}

// AttachToObject mocks base method.
func (m *MockIUseCase) AttachToObject(arg0, arg1 string) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachToObject", arg0, arg1)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachToObject indicates an expected call of AttachToObject.
func (mr *MockIUseCaseMockRecorder) AttachToObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachToObject", reflect.TypeOf((*MockIUseCase)(nil).AttachToObject), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIUseCase) Delete(arg0 string) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIUseCaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUseCase)(nil).Delete), arg0)
}

// DeleteAttr mocks base method.
func (m *MockIUseCase) DeleteAttr(arg0, arg1 string) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAttr", arg0, arg1)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAttr indicates an expected call of DeleteAttr.
func (mr *MockIUseCaseMockRecorder) DeleteAttr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAttr", reflect.TypeOf((*MockIUseCase)(nil).DeleteAttr), arg0, arg1)
}

// DeleteIfExists mocks base method.
func (m *MockIUseCase) DeleteIfExists(arg0 string) usecase.RAM {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIfExists", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	return ret0
}

// DeleteIfExists indicates an expected call of DeleteIfExists.
func (mr *MockIUseCaseMockRecorder) DeleteIfExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIfExists", reflect.TypeOf((*MockIUseCase)(nil).DeleteIfExists), arg0)
}

// DeleteObject mocks base method.
func (m *MockIUseCase) DeleteObject(arg0 string) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockIUseCaseMockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockIUseCase)(nil).DeleteObject), arg0)
}

// Get mocks base method.
func (m *MockIUseCase) Get(arg0 string) (usecase.RAM, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockIUseCaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUseCase)(nil).Get), arg0)
}

// GetFromObject mocks base method.
func (m *MockIUseCase) GetFromObject(arg0, arg1 string) (usecase.RAM, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromObject", arg0, arg1)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFromObject indicates an expected call of GetFromObject.
func (mr *MockIUseCaseMockRecorder) GetFromObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromObject", reflect.TypeOf((*MockIUseCase)(nil).GetFromObject), arg0, arg1)
}

// ObjectToJSON mocks base method.
func (m *MockIUseCase) ObjectToJSON(arg0 string) (usecase.RAM, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectToJSON", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ObjectToJSON indicates an expected call of ObjectToJSON.
func (mr *MockIUseCaseMockRecorder) ObjectToJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectToJSON", reflect.TypeOf((*MockIUseCase)(nil).ObjectToJSON), arg0)
}

// NewObject mocks base method.
func (m *MockIUseCase) NewObject(arg0 string) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewObject", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewObject indicates an expected call of NewObject.
func (mr *MockIUseCaseMockRecorder) NewObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewObject", reflect.TypeOf((*MockIUseCase)(nil).NewObject), arg0)
}

// Save mocks base method.
func (m *MockIUseCase) Save() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Save")
}

// Save indicates an expected call of Save.
func (mr *MockIUseCaseMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUseCase)(nil).Save))
}

// Set mocks base method.
func (m *MockIUseCase) Set(arg0, arg1 string, arg2 bool) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockIUseCaseMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIUseCase)(nil).Set), arg0, arg1, arg2)
}

// SetToObject mocks base method.
func (m *MockIUseCase) SetToObject(arg0, arg1, arg2 string, arg3 bool) (usecase.RAM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetToObject", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetToObject indicates an expected call of SetToObject.
func (mr *MockIUseCaseMockRecorder) SetToObject(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToObject", reflect.TypeOf((*MockIUseCase)(nil).SetToObject), arg0, arg1, arg2, arg3)
}

// Size mocks base method.
func (m *MockIUseCase) Size(arg0 string) (usecase.RAM, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size", arg0)
	ret0, _ := ret[0].(usecase.RAM)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Size indicates an expected call of Size.
func (mr *MockIUseCaseMockRecorder) Size(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockIUseCase)(nil).Size), arg0)
}
