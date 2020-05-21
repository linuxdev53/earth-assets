package controllers

import (
	"fmt"
	"strconv"
	"time"

	"earth-assets/common"
	"earth-assets/models"
	"earth-assets/views"

	"github.com/kataras/iris/v12"
)

func EarthAssets(ctx iris.Context) {
	params := models.NewEarthAssetsParams()

	// parse earth assets parameters and set param for nasa api
	lat := ctx.FormValue("lat")
	if len(lat) > 0 {
		fLatitude, err := strconv.ParseFloat(lat, 32)
		if err != nil {
			fmt.Printf("Invalid latitude field:%s\n", lat)
			views.SendResponse(ctx, iris.StatusBadRequest, common.ErrUnknown, nil)
			return
		}
		params.Latitude = float32(fLatitude)
	}

	lon := ctx.FormValue("lon")
	if len(lon) > 0 {
		fLongitude, err := strconv.ParseFloat(lon, 32)
		if err != nil {
			fmt.Printf("Invalid longitude field:%s\n", lon)
			views.SendResponse(ctx, iris.StatusBadRequest, common.ErrUnknown, nil)
			return
		}
		params.Longitude = float32(fLongitude)
	}

	date := ctx.FormValue("date")
	if len(date) > 0 {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Printf("Invalid date field:%s with err:%v\n", date, err)
			views.SendResponse(ctx, iris.StatusBadRequest, common.ErrUnknown, nil)
			return
		}
		params.Date = date
	}

	dim := ctx.FormValue("dim")
	if len(dim) > 0 {
		fDim, err := strconv.ParseFloat(dim, 32)
		if err != nil {
			fmt.Printf("Invalid degrees field:%s\n", dim)
			views.SendResponse(ctx, iris.StatusBadRequest, common.ErrUnknown, nil)
			return
		}
		params.Degrees = float32(fDim)
	}

	apiKey := ctx.FormValue("api_key")
	if len(apiKey) > 0 {
		params.ApiKey = apiKey
	}

	// call Nasa API
	resp, err := models.GetNasaApiResponse(params)
	if err != nil {
		fmt.Printf("Failed to get nasa API response with err:%v\n", err)
		views.SendResponse(ctx, iris.StatusOK, common.ErrNasaApi, nil)
		return
	}
	fmt.Sprintf("resp:%v\n", resp)

	// generate HTML file
	err = views.ParseTemplate(resp)
	if err != nil {
		fmt.Printf("Failed to generate HTML content with err:%v\n", err)
		views.SendResponse(ctx, iris.StatusOK, common.ErrHTMLInvalid, nil)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.ServeFile(common.IndexHTML, false)
}
