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



	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		salary, _ := strconv.Atoi(r.URL.Query().Get("pensja"))

		var podatek = tax(float64(salary))
		var procentPensji = tax(float64(salary))/float64(salary)*100

		pisz(w, fmt.Sprintf("Twoj podatek to %v", podatek))
		pisz(w, fmt.Sprintf(".Twoj podatek jako procent pensji to %v", procentPensji))

		pisz(w, "<form><input name='pensja'></form>")

	})


	http.HandleFunc("/odejmij", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		odjemna, _ := strconv.Atoi(r.URL.Query().Get("odjemna"))
		odjemnik, _ := strconv.Atoi(r.URL.Query().Get("odjemnik"))

		wynik := odjemna - odjemnik;

		pisz(w, fmt.Sprintf("Twoj wynik to %v", wynik))

		pisz(w, "<form><input name='odjemna'>-<input name='odjemnik'><input type='submit' value='Odejmij'></form>")

	})

	http.HandleFunc("/dodaj", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		s1, _ := strconv.Atoi(r.URL.Query().Get("skladnik1"))
		s2, _ := strconv.Atoi(r.URL.Query().Get("skladnik2"))

		wynik := s1+s2

		pisz(w, fmt.Sprintf("Twoj wynik to %v", wynik))

		pisz(w, "<form><input name='skladnik1'>+<input name='skladnik2'><input type='submit' value='Dodaj'></form>")

	})

	http.HandleFunc("/podziel", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		dzielna, _ := strconv.Atoi(r.URL.Query().Get("dzielna"))
		dzielnik, _ := strconv.Atoi(r.URL.Query().Get("dzielnik"))

		wynik := dzielna/dzielnik

		pisz(w, fmt.Sprintf("Twoj wynik to %v", wynik))

		pisz(w, "<form><input name='dzielna'>/<input name='dzielnik'><input type='submit' value='Podziel'></form>")

	})

	http.HandleFunc("/pomnóż", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		mnożna, _ := strconv.Atoi(r.URL.Query().Get("mnożna"))
		mnożnik, _ := strconv.Atoi(r.URL.Query().Get("mnożnik"))

		wynik := mnożna*mnożnik

		pisz(w, fmt.Sprintf("Twoj wynik to %v", wynik))

		pisz(w, "<form><input name='mnożna'>*<input name='mnożnik'><input type='submit' value='Pomnóż'></form>")

	})

	http.ListenAndServe("0.0.0.0:9999", nil)

}

func pisz(w http.ResponseWriter, tekst string) {
	w.Write([]byte(tekst))
}
