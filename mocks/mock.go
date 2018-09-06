package mocks

import (
	config "changelog/config"
	resolver "changelog/resolver"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
)

// MockMutations is a mock of Mutations interface
type MockMutations struct {
	ctrl     *gomock.Controller
	recorder *MockMutationsMockRecorder
}

// MockMutationsMockRecorder is the mock recorder for MockMutations
type MockMutationsMockRecorder struct {
	mock *MockMutations
}

// NewMockMutations creates a new mock instance
func NewMockMutations(ctrl *gomock.Controller) *MockMutations {
	mock := &MockMutations{ctrl: ctrl}
	mock.recorder = &MockMutationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutations) EXPECT() *MockMutationsMockRecorder {
	return m.recorder
}

// CreateChange mocks base method
func (m *MockMutations) CreateChange(arg0 resolver.InputCreates) *resolver.ChangeResolver {
	ret := m.ctrl.Call(m, "CreateChange", arg0)
	ret0, _ := ret[0].(*resolver.ChangeResolver)
	return ret0
}

// CreateChange indicates an expected call of CreateChange
func (mr *MockMutationsMockRecorder) CreateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChange", reflect.TypeOf((*MockMutations)(nil).CreateChange), arg0)
}

// DeleteChange mocks base method
func (m *MockMutations) DeleteChange(arg0 resolver.InputUpdates) *resolver.ChangeResolver {
	ret := m.ctrl.Call(m, "DeleteChange", arg0)
	ret0, _ := ret[0].(*resolver.ChangeResolver)
	return ret0
}

// DeleteChange indicates an expected call of DeleteChange
func (mr *MockMutationsMockRecorder) DeleteChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChange", reflect.TypeOf((*MockMutations)(nil).DeleteChange), arg0)
}

// UpdateChange mocks base method
func (m *MockMutations) UpdateChange(arg0 resolver.InputDeletes) *resolver.ChangeResolver {
	ret := m.ctrl.Call(m, "UpdateChange", arg0)
	ret0, _ := ret[0].(*resolver.ChangeResolver)
	return ret0
}

// UpdateChange indicates an expected call of UpdateChange
func (mr *MockMutationsMockRecorder) UpdateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChange", reflect.TypeOf((*MockMutations)(nil).UpdateChange), arg0)
}

// MockDBinterface is a mock of DBinterface interface
type MockDBinterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBinterfaceMockRecorder
}

// MockDBinterfaceMockRecorder is the mock recorder for MockDBinterface
type MockDBinterfaceMockRecorder struct {
	mock *MockDBinterface
}

// NewMockDBinterface creates a new mock instance
func NewMockDBinterface(ctrl *gomock.Controller) *MockDBinterface {
	mock := &MockDBinterface{ctrl: ctrl}
	mock.recorder = &MockDBinterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBinterface) EXPECT() *MockDBinterfaceMockRecorder {
	return m.recorder
}

// OpenDB mocks base method
func (m *MockDBinterface) OpenDB(c *config.Config) (*gorm.DB, error) {
	ret := m.ctrl.Call(m, "OpenDB", c)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenDB indicates an expected call of OpenDB
func (mr *MockDBinterfaceMockRecorder) OpenDB(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDB", reflect.TypeOf((*MockDBinterface)(nil).OpenDB), c)
}

// MustOpen mocks base method
func (m *MockDBinterface) MustOpen(c *config.Config) *gorm.DB {
	ret := m.ctrl.Call(m, "MustOpen", c)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// MustOpen indicates an expected call of MustOpen
func (mr *MockDBinterfaceMockRecorder) MustOpen(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustOpen", reflect.TypeOf((*MockDBinterface)(nil).MustOpen), c)
}
