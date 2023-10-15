package business

import "errors"

var (
	//ErrInternalServerError Error caused by system error
	ErrInternalServerError = errors.New("internal Server Error")

	//ErrHasBeenModified Error when update item that has been modified
	ErrHasBeenModified = errors.New("data has been modified")

	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("data was not found")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("given spec is not valid")

	//ErrInvalidCommand Error when the message is formatted like a command but it's not a valid one
	ErrInvalidCommand = errors.New("this is not a valid command")
)
