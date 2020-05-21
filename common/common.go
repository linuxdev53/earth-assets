package common

const (
	ListenAddr = ":8080"

	IndexHTML      = "./assets/index.html"
	IndexHTMLTempl = "./assets/index.html.tmpl"

	NasaApiURL  = "https://api.nasa.gov"
	NasaApiPath = "/planetary/earth/assets"

	DefaultLatitude  = 1.5
	DefaultLongitude = 100.78
	DefaultDegrees   = 0.025
)

type ErrCode int

const (
	ErrSuccess     ErrCode = 0x0000
	ErrNasaApi     ErrCode = 0x0001
	ErrHTMLInvalid ErrCode = 0x0002
	ErrUnknown     ErrCode = 0x0100
)
