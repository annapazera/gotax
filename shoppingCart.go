package main

import "fmt"

var (
	shoppingCartAni  []Product
	shoppingCartKasi []Product
	shoppingCart     []Product
)

type Product struct {
	name  string
	price int
}

func main() {

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

}

func addProduct(product Product) []Product {

	shoppingCart = append(shoppingCart, product)

	return shoppingCart
}

func createProduct(name string, price int) (product Product) {

	product = Product{name, price}

	return product
}
