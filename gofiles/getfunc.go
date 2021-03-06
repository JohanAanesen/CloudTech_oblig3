package gofiles

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// GetCurrency gets the currency from string URL
func GetCurrency() Data2d {

	//currency codes
	s1 := codes
	//initialize the map
	var data2d Data2d
	data2d.Data = make(map[string]map[string]float64)

	//sets date
	data2d.Date = time.Now().Format("2006-01-02")

	for i := 0; i < len(s1); i++ {
		//gets currencies from Fixer with the BASE currency
		json1, err := http.Get("http://api.fixer.io/latest?base=" + s1[i]) //+ "," + s2)
		if err != nil {
			log.Printf("fixer.io is not responding, %s\n", err.Error())
			panic(err)
		}

		//data object
		var data Data

		//json decoder
		err = json.NewDecoder(json1.Body).Decode(&data)
		if err != nil { //err handler
			log.Printf("Error: %s\n", err.Error())
			panic(err)
		}

		//loops through all currency codes and adds them to 2d data map
		for j := 0; j < len(s1); j++ {
			//skip identical currency codes
			if s1[i] != s1[j] {
				//add data to 2d map
				Add2d(data2d.Data, s1[i], s1[j], data.Rates[s1[j]])
			}
		}
		//limit requests to fixer.io to 5 requests per second
		<-time.After(200 * time.Millisecond)
	}

	return data2d
}

//Add2d adds base, target and value of a currency to the 2d map
func Add2d(m map[string]map[string]float64, base string, target string, value float64) {
	mm, ok := m[base]
	//initializes the child maps
	if !ok {
		mm = make(map[string]float64)
		m[base] = mm
	}
	mm[target] = value
}
