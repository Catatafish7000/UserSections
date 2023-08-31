// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package handlers_test is a generated GoMock package.
package handlers_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockrepo is a mock of repo interface.
type Mockrepo struct {
	ctrl     *gomock.Controller
	recorder *MockrepoMockRecorder
}

// MockrepoMockRecorder is the mock recorder for Mockrepo.
type MockrepoMockRecorder struct {
	mock *Mockrepo
}

// NewMockrepo creates a new mock instance.
func NewMockrepo(ctrl *gomock.Controller) *Mockrepo {
	mock := &Mockrepo{ctrl: ctrl}
	mock.recorder = &MockrepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepo) EXPECT() *MockrepoMockRecorder {
	return m.recorder
}

// AddHistory mocks base method.
func (m *Mockrepo) AddHistory(userID int, sectionName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHistory", userID, sectionName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHistory indicates an expected call of AddHistory.
func (mr *MockrepoMockRecorder) AddHistory(userID, sectionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHistory", reflect.TypeOf((*Mockrepo)(nil).AddHistory), userID, sectionName)
}

// AddSection mocks base method.
func (m *Mockrepo) AddSection(userID int, sectionName string, ttl int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSection", userID, sectionName, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSection indicates an expected call of AddSection.
func (mr *MockrepoMockRecorder) AddSection(userID, sectionName, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSection", reflect.TypeOf((*Mockrepo)(nil).AddSection), userID, sectionName, ttl)
}

// CheckExistence mocks base method.
func (m *Mockrepo) CheckExistence(userID int, sectionName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistence", userID, sectionName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckExistence indicates an expected call of CheckExistence.
func (mr *MockrepoMockRecorder) CheckExistence(userID, sectionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistence", reflect.TypeOf((*Mockrepo)(nil).CheckExistence), userID, sectionName)
}

// CreateSection mocks base method.
func (m *Mockrepo) CreateSection(sectionName string, percentage int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSection", sectionName, percentage)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSection indicates an expected call of CreateSection.
func (mr *MockrepoMockRecorder) CreateSection(sectionName, percentage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSection", reflect.TypeOf((*Mockrepo)(nil).CreateSection), sectionName, percentage)
}

// DeleteHistory mocks base method.
func (m *Mockrepo) DeleteHistory(userID int, sectionName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHistory", userID, sectionName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHistory indicates an expected call of DeleteHistory.
func (mr *MockrepoMockRecorder) DeleteHistory(userID, sectionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHistory", reflect.TypeOf((*Mockrepo)(nil).DeleteHistory), userID, sectionName)
}

// DeleteSection mocks base method.
func (m *Mockrepo) DeleteSection(sectionName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSection", sectionName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSection indicates an expected call of DeleteSection.
func (mr *MockrepoMockRecorder) DeleteSection(sectionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSection", reflect.TypeOf((*Mockrepo)(nil).DeleteSection), sectionName)
}

// GetSectionList mocks base method.
func (m *Mockrepo) GetSectionList() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSectionList")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSectionList indicates an expected call of GetSectionList.
func (mr *MockrepoMockRecorder) GetSectionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSectionList", reflect.TypeOf((*Mockrepo)(nil).GetSectionList))
}

// RemoveSection mocks base method.
func (m *Mockrepo) RemoveSection(userID int, sectionName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSection", userID, sectionName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSection indicates an expected call of RemoveSection.
func (mr *MockrepoMockRecorder) RemoveSection(userID, sectionName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSection", reflect.TypeOf((*Mockrepo)(nil).RemoveSection), userID, sectionName)
}

// ShowHistory mocks base method.
func (m *Mockrepo) ShowHistory(year, month int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowHistory", year, month)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowHistory indicates an expected call of ShowHistory.
func (mr *MockrepoMockRecorder) ShowHistory(year, month interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowHistory", reflect.TypeOf((*Mockrepo)(nil).ShowHistory), year, month)
}

// ShowSections mocks base method.
func (m *Mockrepo) ShowSections(userID int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowSections", userID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowSections indicates an expected call of ShowSections.
func (mr *MockrepoMockRecorder) ShowSections(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowSections", reflect.TypeOf((*Mockrepo)(nil).ShowSections), userID)
}