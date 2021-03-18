package main

import "fmt"

type checkStock struct {
	next transactionHandler
}

func (d *checkStock) execute(r *transactionRequest) {
	fmt.Println("Execute checkStock Handler")
	r.checkStockDone = true
	d.next.execute(r)
}

func (d *checkStock) setNext(next transactionHandler) {
	d.next = next
}
