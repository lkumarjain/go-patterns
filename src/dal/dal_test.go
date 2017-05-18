package dal

import (
	"reflect"
	"testing"
)

func TestGetDal(t *testing.T) {
	d := GetDal()
	dType, ok := d.(dalImpl)
	if !ok {
		t.Errorf("Expected type is dalImpl but got %v", reflect.TypeOf(dType))
	}
}

func TestOperationOne(t *testing.T) {
	d := dalImpl{}
	err := d.OperationOne()
	if err != nil {
		t.Errorf("Expected nil but got error %v", err)
	}
}

func TestOperationTwo(t *testing.T) {
	d := dalImpl{}
	err := d.OperationTwo("Test")
	if err != nil {
		t.Errorf("Expected nil but got error %v", err)
	}
}

func TestOperationTwoError(t *testing.T) {
	d := dalImpl{}
	err := d.OperationTwo("Error")
	if err == nil {
		t.Errorf("Expected Error but got nil")
	}
}
