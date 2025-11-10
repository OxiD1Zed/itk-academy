package model

import "errors"

var (
	ErrorUnknowOperationType error = errors.New("unknow operation")
	ErrorClosedConnection    error = errors.New("the connection is closed")
	ErrorAcquireTimeout      error = errors.New("the waiting time is up")
	ErrorNotFound            error = errors.New("not found")
	ErrorInsufficientFunds   error = errors.New("insufficient funds in the balance")
)
