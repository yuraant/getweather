package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	apihost   = "api.openweathermap.org"
	apischeme = "https"
	apipath   = "data/2.5/weather"
)

var (
	apicity string
	apikey  string
	units   = "metric"
	lang    = "en"
)

type apiResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		Pressure  float64 `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func init() {
	var err error
	apikey, err = getMandatoryEnv("OPENWEATHER_API_KEY")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	apicity, err = getMandatoryEnv("CITY_NAME")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func main() {

	url := collectURL(apischeme, apihost, apipath, apicity, getOptionalEnv("UNITS", units), getOptionalEnv("LANG_NAME", lang), apikey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsonData, _ := parseResponse(data)
		fmt.Printf("source=openweathermap, city=\"%s\", description=\"%s\", temp=%v, humidity=%d\n", jsonData.Name, jsonData.Weather[0].Description, jsonData.Main.Temp, jsonData.Main.Humidity)
	}
}

func collectURL(scheme string, host string, path string, city string, units string, lang string, apiKey string) string {
	url := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	q := url.Query()
	q.Add("q", city)
	q.Add("units", units)
	q.Add("lang", lang)
	q.Add("appid", apiKey)
	url.RawQuery = q.Encode()
	return url.String()
}

func getMandatoryEnv(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", errors.New("Environment variable " + key + " is not set. Please add it one")
	}
	return value, nil
}

func getOptionalEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}

func parseResponse(data []byte) (*apiResponse, error) {
	var parsedData = new(apiResponse)
	err := json.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Println("error:", err)
	}
	return parsedData, err
}
