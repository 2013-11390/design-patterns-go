package main

func main() {
	decrypt := &decrypt{}
	auth := &auth{}
	checkStock := &checkStock{}
	payment := &payment{}

	decrypt.setNext(auth)
	auth.setNext(checkStock)
	checkStock.setNext(payment)

	transactionRequest := &transactionRequest{productName: "sekomdalkom"}

	decrypt.execute(transactionRequest)
}
