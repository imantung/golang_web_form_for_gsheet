// Code generated by MockGen. DO NOT EDIT.
// Source: student_repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repo "github.com/imantung/golang_webform_for_gsheet/internal/app/repo"
)

// MockStudentRepo is a mock of StudentRepo interface.
type MockStudentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStudentRepoMockRecorder
}

// MockStudentRepoMockRecorder is the mock recorder for MockStudentRepo.
type MockStudentRepoMockRecorder struct {
	mock *MockStudentRepo
}

// NewMockStudentRepo creates a new mock instance.
func NewMockStudentRepo(ctrl *gomock.Controller) *MockStudentRepo {
	mock := &MockStudentRepo{ctrl: ctrl}
	mock.recorder = &MockStudentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentRepo) EXPECT() *MockStudentRepoMockRecorder {
	return m.recorder
}

// FindOne mocks base method.
func (m *MockStudentRepo) FindOne(gsheet string, row int) (*repo.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", gsheet, row)
	ret0, _ := ret[0].(*repo.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockStudentRepoMockRecorder) FindOne(gsheet, row interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockStudentRepo)(nil).FindOne), gsheet, row)
}

// Update mocks base method.
func (m *MockStudentRepo) Update(gsheet string, row int, student *repo.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", gsheet, row, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStudentRepoMockRecorder) Update(gsheet, row, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStudentRepo)(nil).Update), gsheet, row, student)
}
