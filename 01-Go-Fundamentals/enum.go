package main

import "fmt"

func main() {

	type OrderStatus int

	const (
		Created        OrderStatus = 0
		Processing     OrderStatus = 1
		WaitForPayment OrderStatus = 2
		PaymentDone    OrderStatus = 3
		Issue          OrderStatus = 4
		Refund         OrderStatus = 5
	)

	type Product struct {
		Id     uint32
		Price  uint64
		Status OrderStatus
	}

	p1 := Product{Id: 1, Price: 49, Status: Processing}
	p2 := Product{Id: 2, Price: 149, Status: Issue}
	p3 := Product{Id: 3, Price: 79, Status: Refund}

	fmt.Printf("Id : %d , Price : %d , Status : %d\n", p1.Id, p1.Price, p1.Status)
	fmt.Printf("Id : %d , Price : %d , Status : %d\n", p2.Id, p2.Price, p2.Status)
	fmt.Printf("Id : %d , Price : %d , Status : %d\n", p3.Id, p3.Price, p3.Status)
}
