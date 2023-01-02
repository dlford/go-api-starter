package constants

import (
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var GIN_MODE string
var ACCESS_TOKEN_LIFETIME time.Duration
var REDIS_HOST string
var REDIS_PASSWORD string
var REDIS_DB int
var SERVICE_NAME string
var SESSION_LIFETIME_DAYS int
var SESSION_LIFETIME_SECONDS int
var APPKEY_CACHE_LIFETIME time.Duration
var COOKIE_SAME_SITE http.SameSite
var COOKIE_PATH string
var COOKIE_DOMAIN string
var COOKIE_SECURE bool
var COOKIE_HTTP_ONLY bool
var SESSION_REAPER_SCHEDULE string
var MAIL_HOST string
var MAIL_PORT int
var MAIL_USER string
var MAIL_PASSWORD string
var MAIL_FROM string
var MAIL_FROM_NAME string
var FQDN string
var DB_HOST string
var DB_PORT int
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_SSLMODE string
var EMAIL_VERIFIED_REDIRECT_URL string
var RESET_PASSWORD_URL string

func SetupEnv() {
	oneDay, err := time.ParseDuration("24h")
	if err != nil {
		panic(err)
	}

	s := os.Getenv("ACCESS_TOKEN_LIFETIME")
	if s == "" {
		s = "15m"
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	ACCESS_TOKEN_LIFETIME = d

	s = os.Getenv("REDIS_HOST")
	if s == "" {
		s = "localhost:6379"
	}
	REDIS_HOST = s

	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

	s = os.Getenv("REDIS_DB")
	if s == "" {
		s = "0"
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	REDIS_DB = i

	s = os.Getenv("SERVICE_NAME")
	if s == "" {
		s = "Demo API"
	}
	SERVICE_NAME = s

	s = os.Getenv("COOKIE_SAME_SITE")
	if s == "" {
		s = "lax"
	}
	switch s {
	case "strict":
		COOKIE_SAME_SITE = http.SameSiteStrictMode
	case "none":
		COOKIE_SAME_SITE = http.SameSiteNoneMode
	default:
		COOKIE_SAME_SITE = http.SameSiteLaxMode
	}

	s = os.Getenv("SESSION_LIFETIME_DAYS")
	if s == "" {
		s = "30"
	}
	i, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	SESSION_LIFETIME_DAYS = i
	SESSION_LIFETIME_SECONDS = int(oneDay.Seconds() * float64(i))

	s = os.Getenv("COOKIE_PATH")
	if s == "" {
		s = "/"
	}
	COOKIE_PATH = s

	s = os.Getenv("COOKIE_DOMAIN")
	if s == "" {
		s = "localhost"
	}
	COOKIE_DOMAIN = s

	s = os.Getenv("COOKIE_SECURE")
	if s == "" {
		s = "false"
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	COOKIE_SECURE = b

	s = os.Getenv("COOKIE_HTTP_ONLY")
	if s == "" {
		s = "true"
	}
	b, err = strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	COOKIE_HTTP_ONLY = b

	s = os.Getenv("SESSION_REAPER_SCHEDULE")
	if s == "" {
		s = "0 * * * *"
	}
	SESSION_REAPER_SCHEDULE = s

	MAIL_HOST = os.Getenv("MAIL_HOST")
	MAIL_USER = os.Getenv("MAIL_USER")
	MAIL_PASSWORD = os.Getenv("MAIL_PASSWORD")
	MAIL_FROM = os.Getenv("MAIL_FROM")
	MAIL_FROM_NAME = os.Getenv("MAIL_FROM_NAME")

	s = os.Getenv("MAIL_PORT")
	if s == "" {
		panic("Missing mail port")
	}
	i, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	MAIL_PORT = i

	s = os.Getenv("FQDN")
	if s == "" {
		s = "http://localhost:8080"
	}
	FQDN = s

	if MAIL_HOST == "" || MAIL_USER == "" || MAIL_PASSWORD == "" || MAIL_FROM == "" {
		panic("Missing mail configuration")
	}

	s = os.Getenv("DB_HOST")
	if s == "" {
		s = "localhost"
	}
	DB_HOST = s

	s = os.Getenv("DB_USER")
	if s == "" {
		s = "postgres"
	}
	DB_USER = s

	s = os.Getenv("DB_PASSWORD")
	if s == "" {
		s = "postgres"
	}
	DB_PASSWORD = s

	s = os.Getenv("DB_NAME")
	if s == "" {
		s = "dev"
	}
	DB_NAME = s

	s = os.Getenv("DB_PORT")
	if s == "" {
		s = "5432"
	}
	i, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	DB_PORT = i

	s = os.Getenv("DB_SSLMODE")
	if s == "" {
		s = "disable"
	}
	DB_SSLMODE = s

	s = os.Getenv("GIN_MODE")
	if s == "" {
		s = "debug"
	}
	GIN_MODE = s

	EMAIL_VERIFIED_REDIRECT_URL = os.Getenv("EMAIL_VERIFIED_REDIRECT_URL")

	s = os.Getenv("RESET_PASSWORD_URL")
	if s == "" {
		panic("RESET_PASSWORD_URL is not set")
	}
	RESET_PASSWORD_URL = s

	s = os.Getenv("APPKEY_CACHE_LIFETIME")
	if s == "" {
		s = "12h"
	}

	d, err = time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	APPKEY_CACHE_LIFETIME = d
}
