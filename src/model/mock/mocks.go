// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/lkumarjain/go-patterns/src/model (interfaces: Service,Dal)

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceRecorder
}

// Recorder for MockService (not exported)
type _MockServiceRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &_MockServiceRecorder{mock}
	return mock
}

func (_m *MockService) EXPECT() *_MockServiceRecorder {
	return _m.recorder
}

func (_m *MockService) OperationOne() error {
	ret := _m.ctrl.Call(_m, "OperationOne")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) OperationOne() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OperationOne")
}

func (_m *MockService) OperationTwo(_param0 string) error {
	ret := _m.ctrl.Call(_m, "OperationTwo", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) OperationTwo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OperationTwo", arg0)
}

// Mock of Dal interface
type MockDal struct {
	ctrl     *gomock.Controller
	recorder *_MockDalRecorder
}

// Recorder for MockDal (not exported)
type _MockDalRecorder struct {
	mock *MockDal
}

func NewMockDal(ctrl *gomock.Controller) *MockDal {
	mock := &MockDal{ctrl: ctrl}
	mock.recorder = &_MockDalRecorder{mock}
	return mock
}

func (_m *MockDal) EXPECT() *_MockDalRecorder {
	return _m.recorder
}

func (_m *MockDal) OperationOne() error {
	ret := _m.ctrl.Call(_m, "OperationOne")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDalRecorder) OperationOne() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OperationOne")
}

func (_m *MockDal) OperationTwo(_param0 string) error {
	ret := _m.ctrl.Call(_m, "OperationTwo", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDalRecorder) OperationTwo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OperationTwo", arg0)
}
