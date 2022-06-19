package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DataFile struct {
	Status struct {
		Water int
		Wind  int
	} `json:"status"`
}

func GenerateData() {
	for {
		fmt.Println("Generating data...")
		data := DataFile{}
		dataRand := fmt.Sprintf(`{"status":{"water":%d,"wind":%d}}`, rand.Intn(15), rand.Intn(20))
		fmt.Println(dataRand)
		err := json.Unmarshal([]byte(dataRand), &data)
		if err != nil {
			panic(err)
		}
		newJson, _ := json.Marshal(data)
		err = ioutil.WriteFile("data.json", newJson, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sleep for 15 seconds for generate new data!")
		time.Sleep(time.Second * 15)
	}
}

func main() {
	go GenerateData()
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) {
		file, err := ioutil.ReadFile("data.json")
		var statusWater, statusWind string
		if err != nil {
			panic(err)
		}
		data := DataFile{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			panic(err)
		}
		water := data.Status.Water
		wind := data.Status.Wind
		if water <= 5 {
			statusWater = "aman"
		} else if water <= 8 {
			statusWater = "siaga"
		} else {
			statusWater = "bahaya"
		}
		if wind <= 6 {
			statusWind = "aman"
		} else if water <= 15 {
			statusWind = "siaga"
		} else {
			statusWind = "bahaya"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"water":       water,
			"wind":        wind,
			"statusWater": statusWater,
			"statusWind":  statusWind,
		})

	})
	r.Run(":1111")
}
