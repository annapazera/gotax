package main

import (
"fmt"
"os"
"strconv"
)

func main() {


type Próg struct {
	próg float64
}



var stawki = []Próg{
    {próg: 6600},
    {próg: 11000},
    {próg: 85528},
    {próg: 127000},
    {próg: 0},
	}
    
    
    
    
var tax = func(salary float64) float64 {
    
            fmt.Println("sraka")
            
    var podatek float64 = 0;
    if(salary<stawki[0].próg){
        podatek=0.18*salary-1188;
        if(podatek<0){
            podatek=0;
        }
    }
    if(salary>=stawki[0].próg && salary<stawki[1].próg){
            fmt.Println("dupa")
        podatek=0.18*salary-1188-631.98*(salary-6600)/4400;
        }
    if(salary>=stawki[1].próg && salary<stawki[2].próg){
        podatek=0.18*salary-556.02;
        }
    if(salary>=stawki[2].próg && salary<stawki[3].próg){
        podatek=15395.04+0.32*salary-15395.04-(556.02-556.02*(salary-85.528)/41472);
        }
    if(salary>=stawki[3].próg){
        podatek=0.32*salary;
        }
    
    return podatek;
};


param1 := os.Args[1]

fmt.Println(param1)


salary, _ := strconv.Atoi(param1)

var podatek = tax(float64(salary))

fmt.Println(podatek);

}




