package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

type Token struct {
	ClientID                         string `json:"clientId"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs int    `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      bool   `json:"isAnonymous"`
}

func GetAccessTokenFromEnv() (*Token, error) {
	spDc, exists := os.LookupEnv("SPOTIFY_DC")
	if !exists {
		fmt.Println("missing spotify_key")
		return nil, nil
	}

	spKey, exists := os.LookupEnv("SPOTIFY_KEY")
	if !exists {
		fmt.Println("missing spotify_key")
		return nil, nil
	}

	return GetAccessToken(spDc, spKey)
}

func GetAccessToken(spDc, spKey string) (*Token, error) {
	req, _ := http.NewRequest("GET", "https://open.spotify.com/get_access_token?reason=transport&productType=web_player", nil)

	req.Header.Set("user-agent", UserAgent)
	// can it use cookiejar? https://gist.github.com/stephanebruckert/d35dfc39276ca742a2e1b96fe9e97df4
	req.Header.Set("cookie", fmt.Sprintf("sp_dc=%s; sp_key=%s", spDc, spKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during GET request")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return nil, err
	}

	_ = resp.Body.Close()

	token := Token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("could not unmarshal JSON")
		return nil, err
	}

	return &token, nil
}
