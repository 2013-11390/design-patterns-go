package main

import "fmt"

type auth struct {
	next transactionHandler
}

func (d *auth) execute(r *transactionRequest) {
	if r.checkStockDone {
		fmt.Println("CheckStock already Done")
		d.next.execute(r)
		return
	}
	fmt.Println("Execute checkStock Handler")
	r.checkStockDone = true
	d.next.execute(r)
}

func (d *auth) setNext(next transactionHandler) {
	d.next = next
}
