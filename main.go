package main

import (
	"fmt"
	"github.com/mirrorfm/spotify-webplayer-token-go/app"
	"os"
)

func main() {
	token, err := app.GetAccessToken()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(token)
}
