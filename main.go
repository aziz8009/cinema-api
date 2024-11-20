package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/aziz8009/cinema-app/config"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

var appName = "cinema-app"
var appVersion = "v1.0.0 "

func main() {

	config.LoadEnv()

	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			fmt.Printf("error loading location '%s': %v\n", tz, err)
		}
	}

	fmt.Println("========================")
	fmt.Println(appName, appVersion)
	fmt.Println("========================", os.Getenv("DB_USER"))

	runtime.GOMAXPROCS(runtime.NumCPU())

	e := echo.New()
	// Start the server
	config.StartServer(e)
}
