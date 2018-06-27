package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Cod     string  `json:"cod"`
	Message float32 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp     float32 `json:"temp"`
			Temp_Min float32 `json:"temp_min"`
			Temp_Max float32 `json:"temp_max"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Dt_txt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}

func main() {
	const apiKey = "APIKEY"
	const id = "1863627"
	response, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=" + id + "&appid=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body)
	var data Result
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Println("Located is " + data.City.Name)
	for _, v := range data.List {
		fmt.Printf("%v : %v : %v \n", v.Dt_txt, v.Weather[0].Main, v.Main.Temp-273.15)
	}
}
