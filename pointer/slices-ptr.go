package main

import (
	"log"
	"strings"
)

/*

* When you pass a slice to a function, it gets a copy of the slice.
* It does not mean that you cannot modify the slice.
* You can just modify the elements already present in the slice.

 */ 

func main() {
	EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
    addCountries(EUcountries)
	log.Println(EUcountries)
	upper(EUcountries)
	log.Println(EUcountries)
    addCountries1(&EUcountries)
	log.Println(EUcountries)
}

func upper(countries []string) {
    for k, _ := range countries {
        countries[k] = strings.ToUpper(countries[k])
    }
}

func addCountries(countries []string) {
    countries = append(countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
}
func addCountries1(countries *[]string) {
    *countries = append(*countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
}
