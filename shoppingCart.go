package main

import (
	"fmt"
)

type Product struct {
	name  string
	price int
}

type ShoppingCart struct {
	products []Product
}

func main2() {

	product1 := createProduct("shoes", 200)
	product2 := createProduct("jacket", 500)
	product3 := createProduct("T-shirt", 100)

	shoppingCart1 := ShoppingCart{}
	shoppingCart1.addProduct(product1)

	shoppingCart2 := ShoppingCart{}

	shoppingCart2.addProduct(product2)
	shoppingCart2.addProduct(product3)
	shoppingCart2.addProduct(product1)

	fmt.Println("Shopping cart 1 : ", shoppingCart1)
	fmt.Println("Shopping cart 2 : ", shoppingCart2)


		sum1 := shoppingCart1.sumProductsPricesInTheCart(shoppingCart1.products)
		fmt.Println("Products in shopping cart 1 cost: ", sum1)

		sum2 := shoppingCart2.sumProductsPricesInTheCart(shoppingCart2.products)
		fmt.Println("Products in shopping cart 2 cost: ", sum2)

}

func (sc *ShoppingCart) addProduct(product Product) {
	sc.products = append(sc.products, product)

}

func createProduct(name string, price int) (product Product) {

	product = Product{name, price}

	return product
}

func (sc *ShoppingCart) sumProductsPricesInTheCart(products [] Product) int {

	sum := 0

	for i := range products {
	sum += products[i].price
}
	return sum
}