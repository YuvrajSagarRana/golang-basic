package main

type WeatherData struct {
	locationName string `json:"locationname"`
	Weather string `json:"weather"`
	Temperature int `json:"temperature"`
	Clesius bool `json:"clesius"`
	TempForcast []int `json:"temp_forcast"`
	WindData windData `json:"wind_data"`
}

type windData struct {
	Direction string `json:"direction"`
	Speed int `json:"speed"`

}
func main()  {

}
