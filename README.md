# spotify-webplayer-token

Helper to retrieve a webplayer Spotify token that is valid for 1 year.

Inspired from https://github.com/enriquegh/spotify-webplayer-token and adapted to golang.

## To obtain the cookies (valid for 1 year):

 - Open a new Incognito window in Chrome (or another browser) at https://accounts.spotify.com/en/login?continue=https:%2F%2Fopen.spotify.com%2F
 - Open Developer Tools in your browser (might require developer menu to be enabled in some browsers)
 - In the Network tab, enable "Preserve log"
 - Login to Spotify.
 - In the Network tab, search/Filter for `password`
 - Under cookies for the request save the values for `sp_dc` and `sp_key`.
 - Close the window without logging out (Otherwise the cookies are made invalid).

## To obtain a Spotify token:

### Programmatically

1. Store `sp_dc` and `sp_key` respectively as environment variables `SPOTIFY_DC` and `SPOTIFY_KEY`,
2. Do:

    ```golang
    import (
        "github.com/mirrorfm/spotify-webplayer-token/app"
    )

    func main() {
        token, err := app.GetAccessTokenFromEnv()
        // use token.AccessToken
    }
    ```

### From the command line

1. In the current repo, create `env.json` such as:

       {
           "SPOTIFY_DC": "AQAJJvY80mY--nNZ_vWVMyWrvrSse...",
           "SPOTIFY_KEY": "478fafec-ed..."
       }

2. Run:

       make setup && make run
