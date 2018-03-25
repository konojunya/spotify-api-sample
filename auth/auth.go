package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	jwt "github.com/makki0205/gojwt"
	"golang.org/x/oauth2"
)

var (
	config oauth2.Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config = oauth2.Config{
		ClientID:     os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
		RedirectURL: "http://localhost:9000/oauth",
		Scopes:      []string{"playlist-modify", "user-read-private", "user-library-read"},
	}
}

// GetRedirectURL リダイレクトする
func GetRedirectURL() string {
	// CSRF対策のコードを抜いている
	return config.AuthCodeURL("spotify")
}

func GetJWTToken(code string) (string, error) {
	jwt.SetSalt("D79998A7-3F2B-4505-BE2B-6E68500AAE37")
	jwt.SetExp(60 * 60 * 24)

	token, err := config.Exchange(oauth2.NoContext, code)
	claims := map[string]string{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
	}

	return jwt.Generate(claims), err
}

// GetClient http clientを取得
func GetClient(token *oauth2.Token) *http.Client {
	return config.Client(oauth2.NoContext, token)
}
