package main

import (
	"encoding/json"

	"io/ioutil"

	"log"

	"net/http"
)

type WeatherInfoJson struct {
	Weatherinfo WeatherinfoObject
}

type WeatherinfoObject struct {
	City string

	Temp string

	WD string

	WS string

	SD string

	Time string
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	resp, err := http.Get("http://www.weather.com.cn/data/sk/101160101.html")

	if err != nil {

		log.Fatal(err)

	}

	defer resp.Body.Close()

	input, err := ioutil.ReadAll(resp.Body)

	var jsonWeather WeatherInfoJson

	json.Unmarshal(input, &jsonWeather)

	log.Printf("城市:%v\n", jsonWeather.Weatherinfo.City)
	log.Printf("温度:%v\n", jsonWeather.Weatherinfo.Temp)
	log.Printf("风向:%v\n", jsonWeather.Weatherinfo.WD)
	log.Printf("风力:%v\n", jsonWeather.Weatherinfo.WS)
	log.Printf("湿度:%v\n", jsonWeather.Weatherinfo.SD)
	log.Printf("时间:%v\n", jsonWeather.Weatherinfo.Time)

}
