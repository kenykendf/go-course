package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type DataFile struct {
	Weather struct {
		Water       string `json:"water"`
		Wind        string `json:"wind"`
		StatusWater string `json:"statusWater"`
		StatusWind  string `json:"statusWind"`
	} `json:"status"`
}

func main() {
	for {
		data := DataFile{}

		water := rand.Intn(100)
		wind := rand.Intn(100)

		switch wr := water; {
		case wr < 5:
			data.Weather.StatusWater = "aman"
		case wr >= 6 && wr <= 8:
			data.Weather.StatusWater = "siaga"
		case wr > 8:
			data.Weather.StatusWater = "bahaya"
		}

		switch wi := wind; {
		case wi < 6:
			data.Weather.StatusWind = "aman"
		case wi >= 7 && wi <= 15:
			data.Weather.StatusWind = "siaga"
		case wi > 15:
			data.Weather.StatusWind = "bahaya"
		}

		data.Weather.Water += fmt.Sprintf("%d meter", water)
		data.Weather.Wind += fmt.Sprintf("%d m/s", wind)

		newJson, _ := json.Marshal(data)
		ioutil.WriteFile("file.json", newJson, 0644)

		time.Sleep(time.Second * 15)
	}
}
