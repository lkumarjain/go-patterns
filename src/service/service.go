package service

import (
	"fmt"

	"github.com/lkumarjain/go-patterns/src/dal"
	"github.com/lkumarjain/go-patterns/src/model"
)

//GetService : Package level method to create Service instance
func GetService() model.Service {
	return serviceImpl{dal: dal.GetDal()}
}

type serviceImpl struct {
	dal model.Dal
}

func (s serviceImpl) OperationOne() error {
	err := s.dal.OperationOne()
	if err != nil {
		fmt.Printf("Error %v in Operation One\n", err)
		return err
	}
	fmt.Println("Operation One Sucessfully executed !!")
	return nil
}

func (s serviceImpl) OperationTwo(message string) error {
	err := s.dal.OperationTwo(message)
	if err != nil {
		fmt.Printf("Error %v in Operation Two\n", err)
		return err
	}
	fmt.Println("Operation Two Sucessfully executed !!")
	return nil
}
