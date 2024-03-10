package main

import (
	"github.com/ak-yudha/crud-gin/routers"
	"golang.org/x/oauth2"
)

func main() {
	oauthConfig := &oauth2.Config{}

	r := routers.SetupRouter(oauthConfig)
	r.Run(":8080")
}
