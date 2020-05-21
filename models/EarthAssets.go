package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"earth-assets/common"
)

type EarthAssetsParams struct {
	Latitude  float32
	Longitude float32
	Date      string
	Degrees   float32
	ApiKey    string
}

type EarthAssetsResource struct {
	Dataset string `json:"dataset"`
	Planet  string `json:"planet"`
}

type EarthAssetsResp struct {
	Date           string              `json:"date"`
	ID             string              `json:"id"`
	Resource       EarthAssetsResource `json:"resource"`
	ServiceVersion string              `json:"service_version"`
	URL            string              `json:"url"`
}

func NewEarthAssetsParams() *EarthAssetsParams {
	return &EarthAssetsParams{
		Latitude:  common.DefaultLatitude,
		Longitude: common.DefaultLongitude,
		Date:      time.Now().Format("2006-01-01"),
		Degrees:   0.025,
		ApiKey:    "DEMO_KEY",
	}
}

func GetNasaApiResponse(params *EarthAssetsParams) (*EarthAssetsResp, error) {
	// create URL data
	urlData := url.Values{}
	if params.Latitude > 0 {
		urlData.Set("lat", fmt.Sprintf("%f", params.Latitude))
	}

	if params.Longitude > 0 {
		urlData.Set("lon", fmt.Sprintf("%f", params.Longitude))
	}

	urlData.Set("date", params.Date)
	urlData.Set("dim", fmt.Sprintf("%f", params.Degrees))
	urlData.Set("api_key", params.ApiKey)

	url, _ := url.ParseRequestURI(common.NasaApiURL)
	url.Path = common.NasaApiPath
	url.RawQuery = urlData.Encode()
	urlStr := fmt.Sprintf("%v", url)

	// send a request to nasa API
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(urlData.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respJsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse response from nasa
	nasaResp := &EarthAssetsResp{}
	err = json.Unmarshal(respJsonBytes, nasaResp)
	if err != nil {
		return nil, err
	}

	return nasaResp, nil
}
