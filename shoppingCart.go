package main

import (
	"fmt"
)

var (
	shoppingCartAni  []Product
	shoppingCartKasi []Product
	shoppingCart     []Product
	clients [] Client


)

type Product struct {
	name  string
	price int
}

type Client struct {
	name string
	nick string
}

func main2() {

	product1 := createProduct("shoes", 200)
	product2 := createProduct("jacket", 500)
	product3 := createProduct("T-shirt", 100)

	shoppingCartAni := addProduct(product1)
	fmt.Println("Zawartość koszyka Ani", shoppingCartAni)

	shoppingCartAni = addProduct(product2)
	fmt.Println("Zawartość koszyka Ani", shoppingCartAni)

	shoppingCartAni = addProduct(product3)
	fmt.Println("Zawartość koszyka Ani", shoppingCartAni)

	shoppingCartKasi = addProduct(product3)
	fmt.Println("Zawartość koszyka Kasi", shoppingCartKasi)


	client1 := createClient("Jan Kowalski", "kowal")
	client2 := createClient("Edmund Świerkowski", "świrek")

	clients := addClient(client1)
	clients = addClient(client2)

	fmt.Println("Lista klientów: ", clients)

	for _, v := range clients {
		carts :=createShoppingCartForClient(v)

	fmt.Println("ShoppingCart for Clients", carts)
	}

}

func addProduct(product Product) []Product {

	shoppingCart = append(shoppingCart, product)

	return shoppingCart
}

func addClient(client Client) []Client {

	clients = append(clients, client)

	return clients
}

func createProduct(name string, price int) (product Product) {

	product = Product{name, price}

	return product
}

func createClient(name string, nick string) (client Client) {

	client = Client{name, nick}

	return client
}

func createShoppingCartForClient  (client Client) ( shoppingCartForClient []Product) {

	shoppingCartForClient = []Product {}

	return shoppingCartForClient
}