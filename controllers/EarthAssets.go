package controllers

import (
	"fmt"
	"strconv"
	"time"

	"earth-assets/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// EarthAssets represents API controller
type EarthAssets struct {
	Ctx iris.Context
}

// Get handles GET: http://localhost:8080/api/EarthAssets
func (e *EarthAssets) Get(ctx iris.Context) mvc.Result {
	var (
		lat          = ctx.FormValue("lat")
		lon          = ctx.FormValue("lon")
		date         = ctx.FormValue("date")
		dim          = ctx.FormValue("dim")
		apiKey       = ctx.FormValue("api_key")
		err    error = nil
	)

	// set params from form value
	params := middleware.NewEarthAssetsParams()
	if len(lat) > 0 {
		fLatitude, err := strconv.ParseFloat(lat, 32)
		if err == nil {
			params.Latitude = float32(fLatitude)
		}
	}

	if len(lon) > 0 {
		fLongitude, err := strconv.ParseFloat(lon, 32)
		if err == nil {
			params.Longitude = float32(fLongitude)
		}
	}

	if len(date) > 0 {
		_, err := time.Parse("2006-01-02", date)
		if err == nil {
			params.Date = date
		}
	}

	if len(dim) > 0 {
		fDim, err := strconv.ParseFloat(dim, 32)
		if err == nil {
			params.Degrees = float32(fDim)
		}
	}

	if len(apiKey) > 0 {
		params.ApiKey = apiKey
	}

	if err != nil {
		fmt.Printf("Failed to parse parameter with err:%v\n", err)
		return mvc.View{
			Name: "assets/error.html",
			Data: iris.Map{
				"Message": fmt.Sprintf("%v", err),
			},
		}
	}

	// call Nasa API
	resp, err := middleware.GetNasaApiResponse(params)
	if err != nil {
		fmt.Printf("Failed to get nasa API response with err:%v\n", err)
		return mvc.View{
			Name: "assets/error.html",
			Data: iris.Map{
				"Message": fmt.Sprintf("%v", err),
			},
		}
	}

	// when image isn't exist, then return eror message
	if len(resp.URL) == 0 {
		return mvc.View{
			Name: "assets/error.html",
			Data: iris.Map{
				"Message": resp.Msg,
			},
		}
	}

	return mvc.View{
		Name: "assets/index.html",
		Data: iris.Map{
			"URL": resp.URL,
		},
	}
}
