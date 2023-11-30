package main

/*
Copyright 2023 Rachel Sheikh

	Permission is hereby granted, free of charge, to any person obtaining a copy of
	this software and associated documentation files (the “Software”), to deal in the
	Software without restriction, including without limitation the rights to use,
	copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
	Software, and to permit persons to whom the Software is furnished to do so,
	subject to the following conditions:

	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
	PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
	OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
	SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/sheikhrachel/workbench/api_common/call"
	"github.com/sheikhrachel/workbench/api_common/router"
	"github.com/sheikhrachel/workbench/api_common/utils/errutil"
	"github.com/sheikhrachel/workbench/api_common/utils/httpUtil"
	"github.com/sheikhrachel/workbench/handlers"
)

func main() {
	// set the cpu cores config
	runtime.GOMAXPROCS(runtime.NumCPU())
	// setting up the router
	r := router.CreateRouterWithMiddleware()
	cc := router.InitCall()
	// setting up the routes
	// - this will set up the routes from
	//   the handlers package, within handlers/routes.go
	handlers.SetRoutes(r, cc)
	// setting up swagger
	r.Static("/swagger", "./swaggerui")

	cc.InfoF("The weather in Seattle, WA is %+v°F", getWeather(cc).Current.TempF)

	// starting the router
	// - this will start accepting results on port 8080
	router.StartRouter(r, cc)
}

const weatherURL = "http://api.weatherapi.com/v1"

// getWeather returns the weather for Seattle, WA using the Weather API
func getWeather(cc call.Call) (weather WeatherAPIResponse) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return WeatherAPIResponse{}
	}
	body, err := httpUtil.GetRespBody(cc, apiKey, weatherURL+"/current.json?key="+apiKey+"&q=Seattle", nil)
	if errutil.HandleError(cc, err) {
		return WeatherAPIResponse{}
	}
	if err = json.Unmarshal(body, &weather); errutil.HandleError(cc, err) {
		return WeatherAPIResponse{}
	}
	return weather
}

type WeatherAPIResponse struct {
	Current WeatherAPIResponseCurrent `json:"current"`
}

type WeatherAPIResponseCurrent struct {
	TempF float64 `json:"temp_f"`
}
