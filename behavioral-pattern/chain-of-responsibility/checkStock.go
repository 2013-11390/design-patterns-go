package main

import "fmt"

type checkStock struct {
	next transactionHandler
}

func (d *checkStock) execute(r *transactionRequest) {
	if r.checkStockDone {
		fmt.Println("CheckStock already Done")
		d.next.execute(r)
		return
	}
	fmt.Println("Execute checkStock Handler")
	r.checkStockDone = true
	d.next.execute(r)
}

func (d *checkStock) setNext(next transactionHandler) {
	d.next = next
}
