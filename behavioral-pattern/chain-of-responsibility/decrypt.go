package main

import "fmt"

type decrypt struct {
	next transactionHandler
}

func (d *decrypt) execute(r *transactionRequest) {
	if r.decryptDone {
		fmt.Println("Decrypt already Done")
		d.next.execute(r)
		return
	}
	fmt.Println("Execute Decrypt Handler")
	r.decryptDone = true
	d.next.execute(r)
}

func (d *decrypt) setNext(next transactionHandler) {
	d.next = next
}
