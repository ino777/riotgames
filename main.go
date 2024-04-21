package main

import (
	"fmt"
	riotgames "lol/api"
)

func main() {
	// engine := gin.Default()
	// engine.LoadHTMLGlob("templates/*")
	// engine.Static("/static", "./static")
	// engine.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"message": "hello world!!!",
	// 	})
	// })
	// engine.Run(":3000")
	sample()
}

func sample() error {
	apiKey := "xxxxx-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	client := riotgames.New(apiKey)
	championInfo, _ := client.GetChampionRotations()
	fmt.Printf("%+v", *championInfo)
	return nil
}
