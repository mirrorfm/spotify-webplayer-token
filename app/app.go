package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

func GetAccessToken() (string, error) {
	spDc, exists := os.LookupEnv("SPOTIFY_DC")
	if !exists {
		fmt.Println("missing spotify_key")
		return "", nil
	}

	spKey, exists := os.LookupEnv("SPOTIFY_KEY")
	if !exists {
		fmt.Println("missing spotify_key")
		return "", nil
	}

	req, _ := http.NewRequest("GET", "https://open.spotify.com/get_access_token?reason=transport&productType=web_player", nil)

	req.Header.Set("user-agent", UserAgent)

	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:  "sp_dc",
		Value: spDc,
	}
	cookies = append(cookies, cookie)

	cookie = &http.Cookie{
		Name:  "sp_key",
		Value: spKey,
	}

	cookies = append(cookies, cookie)

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during GET request")
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return "", err
	}

	_ = resp.Body.Close()

	token := string(body)

	return token, nil
}
