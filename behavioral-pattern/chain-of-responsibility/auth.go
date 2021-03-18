package main

import "fmt"

type auth struct {
	next transactionHandler
}

func (d *auth) execute(r *transactionRequest) {
	if r.authDone {
		fmt.Println("Auth already Done")
		d.next.execute(r)
		return
	}
	fmt.Println("Execute auth Handler")
	r.authDone = true
	d.next.execute(r)
}

func (d *auth) setNext(next transactionHandler) {
	d.next = next
}
