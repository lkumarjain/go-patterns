package main

import "github.com/lkumarjain/go-patterns/src/service"

func main() {
	s := service.GetService()
	s.OperationOne()
	s.OperationTwo("Test")
}
