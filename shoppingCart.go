package main

import (
	"fmt"
)


func NewShoppingCart () ShoppingCart{
	return ShoppingCart{products:map [string] invoiceLine{}}
}

type Product struct {
	name  string
	price int
}

type invoiceLine struct {
	product Product
	amount int
}

type ShoppingCart struct {
	products  map[string]invoiceLine
}


func main2() {

	product1 := createProduct("shoes", 200)
	product2 := createProduct("jacket", 500)
	product3 := createProduct("T-shirt", 100)
	product4 := createProduct("socks", 30)

	shoppingCart1 := ShoppingCart{}
	shoppingCart1.addProduct(product1)
	shoppingCart1.addProduct(product4)

	shoppingCart2 := ShoppingCart{}

	shoppingCart2.addProduct(product2)
	shoppingCart2.addProduct(product3)
	shoppingCart2.addProduct(product1)

	fmt.Println("Shopping cart 1 : ", shoppingCart1)
	fmt.Println("Shopping cart 2 : ", shoppingCart2)


		sum1 := shoppingCart1.sumProductsPricesInTheCart()
		fmt.Println("Products in shopping cart 1 cost: ", sum1)

		sum2 := shoppingCart2.sumProductsPricesInTheCart()
		fmt.Println("Products in shopping cart 2 cost: ", sum2)

}

func (sc *ShoppingCart) addProduct(product Product) {

	il, ok := sc.products[product.name]
	if !ok {
		il = invoiceLine{product, 0}
	}

	il.amount++

	sc.products[product.name] = il


}

func createProduct(name string, price int) (product Product) {

	product = Product{name, price}

	return product
}

func (sc *ShoppingCart) sumProductsPricesInTheCart() int {

	sum := 0

	for _, il := range sc.products {
	sum += (il.product.price * il.amount)
}
	return sum
}

func  sumProductsPricesInTheCart2(sc *ShoppingCart) int {

	sum := 0

	for _, il := range sc.products {
		sum += (il.product.price * il.amount)
	}
	return sum
}

func (sc *ShoppingCart) checkTheNumberOfProducts () int {

	sum := 0

	for _, il := range sc.products {
		sum +=  il.amount
	}
	return sum
}