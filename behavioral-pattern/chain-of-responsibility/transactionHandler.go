package main

type transactionHandler interface {
	execute(*transactionRequest)
	setNext(transactionHandler)
}
