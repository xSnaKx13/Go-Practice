package main

import "fmt"

func main(){
	var a, b string
	var c float64
	var m = map[string]map[string]float64{
		"rub":{
			"usd": 0.011,
			"eur": 0.010,
		},
		"usd":{
			"rub": 91.65,
			"eur" : 0.92,
		},
	}
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	currency := map[a][b]
}
