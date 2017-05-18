package dal

import (
	"errors"
	"fmt"

	"github.com/lkumarjain/go-patterns/src/model"
)

//GetDal : Package level method to create Dal instance
func GetDal() model.Dal {
	return dalImpl{}
}

type dalImpl struct {
}

func (dalImpl) OperationOne() error {
	fmt.Println("Operation One Called !!")
	return nil
}

func (dalImpl) OperationTwo(message string) error {
	if message != "Test" {
		return errors.New("Unknown Message")
	}
	fmt.Printf("Operation 2 called with %s message ", message)
	return nil
}
