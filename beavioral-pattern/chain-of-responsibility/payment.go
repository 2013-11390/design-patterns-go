package main

import "fmt"

type payment struct {
	next transactionHandler
}

func (d *payment) execute(r *transactionRequest) {
	if r.paymentDone {
		fmt.Println("Payment already Done")
		d.next.execute(r)
		return
	}
	r.paymentDone = true
	fmt.Println("Execute payment Handler")
}

func (d *payment) setNext(next transactionHandler) {
	d.next = next
}
