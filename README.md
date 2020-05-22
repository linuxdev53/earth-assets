# Amazing visualization of the acquisition pattern for Landsat 8 imagery

[![Build Status](https://travis-ci.com/linuxdev53/earth-assets.svg?branch=master)](https://travis-ci.com/linuxdev53/earth-assets)

This endpoint retrieves the date-times and asset names for closest available imagery for a supplied location and date. The satellite passes over each point on earth roughly once every sixteen days. This is an amazing visualization of the acquisition pattern for Landsat 8 imagery. The objective of this endpoint is primarily to support the use of the imagery endpoint.

It uses nasa API endpoint for Earth->Assets.(https://api.nasa.gov/)

## How to build and run

Install `golang`:

```sh
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

Build the project:

```sh
git clone git@github.com:linuxdev53/earth-assets.git
cd earth-assets/
go build
./earth-assets
```

To run by docker, you should install `docker-compose` utility.

```sh
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

Run a docker:

```sh
sudo docker-compose up --build
```

## How to test

### HTTP Request

```
http://localhost:8080/api/EarthAssets
```

### Query Parameters

 Parameter | Type | Default | Description
 --------- | -------- | ----- | ----------
 lat | float | n/a | Latitude
 lon | float | n/a | Longitude
 date | YYYY-MM-DD | n/a | beginning of 30 day date range that will be used to look for closest image to that date
 dim | float | 0.025 | width and height of image in degrees
 api_key | string | DEMO_KEY | api.nasa.gov key for expanded usage

### Example query

```
http://localhost:8080/api/EarthAssets?lon=-95.33&lat=29.78&date=2018-01-01&&dim=0.10&api_key=DEMO_KEY
```
