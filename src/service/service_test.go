package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lkumarjain/go-patterns/src/model/mock"
)

func TestGetService(t *testing.T) {
	s := GetService()
	dType, ok := s.(serviceImpl)
	if !ok {
		t.Errorf("Expected type is serviceImpl but got %v", reflect.TypeOf(dType))
	}
}

func TestOperationOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := mock.NewMockDal(ctrl)
	d.EXPECT().OperationOne().Return(nil)
	s := serviceImpl{
		dal: d,
	}
	err := s.OperationOne()
	if err != nil {
		t.Errorf("Expected nil but got error %v", err)
	}
}

func TestOperationOneError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := mock.NewMockDal(ctrl)
	d.EXPECT().OperationOne().Return(errors.New("Error"))
	s := serviceImpl{
		dal: d,
	}
	err := s.OperationOne()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestOperationTwo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := mock.NewMockDal(ctrl)
	d.EXPECT().OperationTwo(gomock.Any()).Return(nil)
	s := serviceImpl{
		dal: d,
	}
	err := s.OperationTwo("test")
	if err != nil {
		t.Errorf("Expected nil but got error %v", err)
	}
}

func TestOperationTwoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := mock.NewMockDal(ctrl)
	d.EXPECT().OperationTwo(gomock.Any()).Return(errors.New("Error"))
	s := serviceImpl{
		dal: d,
	}
	err := s.OperationTwo("message")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
