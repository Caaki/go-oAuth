package auth

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

const (
	key    = "SomethingRandom"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))

	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd
	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://localhost:6969/auth/google"))

}
