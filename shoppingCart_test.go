package main

import (
	"testing"

	"fmt"
)


func TestSumProductPrices(t *testing.T) {




	products := [] Product{

		{"shoes", 400},
		{"sweater", 150},
		{"jeans", 300},
		{"dress", 250},
	}


	shoppingCart :=NewShoppingCart()

	for _, p := range products {
		shoppingCart.addProduct(p)
	}

		sum := shoppingCart.sumProductsPricesInTheCart()


	if sum != 1100 {
		t.Error("Sum was incorrect, got: %d, want: %d.", sum, 1100)
	}


	}



func TestAddProduct(t *testing.T) {


	product1 := Product{"dress", 500}
	product2 := Product{"jacket", 500}
	product3 := Product{"T-shirt", 100}
	product4 := Product{"socks", 30}

	shoppingCart :=NewShoppingCart()



	var dupa func (*ShoppingCart, product Product)

	dupa = (*ShoppingCart).addProduct

	fmt.Println(dupa)

	(*ShoppingCart).addProduct(&shoppingCart, product1)

	shoppingCart.addProduct(product1)

	shoppingCart.addProduct(product2)
	shoppingCart.addProduct(product3)
	shoppingCart.addProduct(product4)

	fmt.Println("Zawartość koszyka", shoppingCart)

	numberOfItemsInTheCart := shoppingCart.checkTheNumberOfProducts()

	if numberOfItemsInTheCart != 4 {
		t.Error("The number of products in the shopping cart was incorrect, got: %d, " +
			"want: " + "%d ", len(shoppingCart.products), 4)
	}
	//if shoppingCart.products[0].name != "dress"{
	//	t.Error("The number of products in the shopping cart was incorrect, got: %d, " +
	//		"want: " +"%d.", product1.name, "dress")
	//
	//}


}

