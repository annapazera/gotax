package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	type Próg struct {
		próg float64
		stawka float64
	}

	var stawki = []Próg{
		{stawka: 0, próg: 6600},
		{stawka: 0.18, próg: 11000},
		{stawka: 0.18, próg: 85528},
		{stawka: 0.32, próg: 127000},
		{stawka: 0.32, próg: 0},
	}

	var tax = func(salary float64) float64 {

		fmt.Println("sroka")

		var podatek float64 = 0
		if salary < stawki[0].próg {
			podatek = stawki[1].stawka*salary - 1188
			if podatek < 0 {
				podatek = 0
			}
		}
		if salary >= stawki[0].próg && salary < stawki[1].próg {
			fmt.Println("zupa")
			podatek = stawki[1].stawka*salary - 1188 - 631.98*(salary-6600)/4400
		}
		if salary >= stawki[1].próg && salary < stawki[2].próg {
			podatek = stawki[1].stawka*salary - 556.02
		}
		if salary >= stawki[2].próg && salary < stawki[3].próg {
			podatek = 15395.04 + stawki[3].stawka*salary - 15395.04 - (556.02 - 556.02*(salary-85.528)/41472)
		}
		if salary >= stawki[3].próg {
			podatek = stawki[3].stawka * salary
		}

		return podatek
	}

	param1 := "12223"

	fmt.Println(param1)

	salary, _ := strconv.Atoi(param1)

	var podatek = tax(float64(salary))
	fmt.Println("twoj podatek to")
	fmt.Println(podatek)

	var procentPensji = tax(float64(salary))/float64(salary)*100
	fmt.Println("podatek jako procent pensji to")
	fmt.Println(procentPensji)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		salary, _ := strconv.Atoi(r.URL.Query().Get("pensja"))

		var podatek = tax(float64(salary))
		var procentPensji = tax(float64(salary))/float64(salary)*100

		pisz(w, fmt.Sprintf("Twoj podatek to %v", podatek))
		pisz(w, fmt.Sprintf(".Twoj podatek jako procent pensji to %v", procentPensji))

		pisz(w, "<form><input name='pensja'></form>")

	})



	http.ListenAndServe("0.0.0.0:9999", nil)

}

func pisz(w http.ResponseWriter, tekst string) {
	w.Write([]byte(tekst))
}
