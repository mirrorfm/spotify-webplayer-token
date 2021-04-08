package main

import (
	"fmt"
	"github.com/mirrorfm/spotify-webplayer-token/app"
	"os"
)

func main() {
	token, err := app.GetAccessTokenFromEnv()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(token)
}
