package main

type transactionRequest struct {
	productName    string
	decryptDone    bool
	authDone       bool
	checkStockDone bool
	paymentDone    bool
}
