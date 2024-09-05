// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/deploy/lattice (interfaces: TargetGroupManager)

// Package lattice is a generated GoMock package.
package lattice

import (
	context "context"
	reflect "reflect"

	core "github.com/aws/aws-application-networking-k8s/pkg/model/core"
	lattice0 "github.com/aws/aws-application-networking-k8s/pkg/model/lattice"
	vpclattice "github.com/aws/aws-sdk-go/service/vpclattice"
	gomock "github.com/golang/mock/gomock"
)

// MockTargetGroupManager is a mock of TargetGroupManager interface.
type MockTargetGroupManager struct {
	ctrl     *gomock.Controller
	recorder *MockTargetGroupManagerMockRecorder
}

// MockTargetGroupManagerMockRecorder is the mock recorder for MockTargetGroupManager.
type MockTargetGroupManagerMockRecorder struct {
	mock *MockTargetGroupManager
}

// NewMockTargetGroupManager creates a new mock instance.
func NewMockTargetGroupManager(ctrl *gomock.Controller) *MockTargetGroupManager {
	mock := &MockTargetGroupManager{ctrl: ctrl}
	mock.recorder = &MockTargetGroupManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetGroupManager) EXPECT() *MockTargetGroupManagerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTargetGroupManager) Delete(arg0 context.Context, arg1 *lattice0.TargetGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTargetGroupManagerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTargetGroupManager)(nil).Delete), arg0, arg1)
}

// IsTargetGroupMatch mocks base method.
func (m *MockTargetGroupManager) IsTargetGroupMatch(arg0 context.Context, arg1 *lattice0.TargetGroup, arg2 *vpclattice.TargetGroupSummary, arg3 *lattice0.TargetGroupTagFields) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTargetGroupMatch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsTargetGroupMatch indicates an expected call of IsTargetGroupMatch.
func (mr *MockTargetGroupManagerMockRecorder) IsTargetGroupMatch(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTargetGroupMatch", reflect.TypeOf((*MockTargetGroupManager)(nil).IsTargetGroupMatch), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockTargetGroupManager) List(arg0 context.Context) ([]tgListOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]tgListOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTargetGroupManagerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTargetGroupManager)(nil).List), arg0)
}

// ResolveRuleTgIds mocks base method.
func (m *MockTargetGroupManager) ResolveRuleTgIds(arg0 context.Context, arg1 *lattice0.RuleAction, arg2 core.Stack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveRuleTgIds", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveRuleTgIds indicates an expected call of ResolveRuleTgIds.
func (mr *MockTargetGroupManagerMockRecorder) ResolveRuleTgIds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveRuleTgIds", reflect.TypeOf((*MockTargetGroupManager)(nil).ResolveRuleTgIds), arg0, arg1, arg2)
}

// Upsert mocks base method.
func (m *MockTargetGroupManager) Upsert(arg0 context.Context, arg1 *lattice0.TargetGroup) (lattice0.TargetGroupStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(lattice0.TargetGroupStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockTargetGroupManagerMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTargetGroupManager)(nil).Upsert), arg0, arg1)
}