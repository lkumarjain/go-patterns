package model

//Service is an interface implemented by service layer of this
type Service interface {
	OperationOne() error
	OperationTwo(message string) error
}

//Dal is an interface implemnted by Dal layer
type Dal interface {
	OperationOne() error
	OperationTwo(message string) error
}
