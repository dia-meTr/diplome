package config

import (
	"flag"
	"fmt"
	"sync"

	goconfig "github.com/Yalantis/go-config"
)

var (
	config        Config
	isInitialised bool
	mutex         = new(sync.Mutex)
)

const configFilename = "config.json"

type Config struct {
	Address   string   `json:"address" default:":8080"`
	JwtSecret string   `json:"jwt_secret" default:"my-totally-secret-key"`
	JWT       JWT      `json:"jwt"`
	Oauth     Oauth    `json:"oauth"`
	Postgres  Postgres `json:"postgres"`
	Redis     Redis    `json:"redis"`
	Log       Logger   `json:"logger"`
	AWS       AWS      `json:"aws"`
}

type Oauth struct {
	Google Google `json:"google"`
}

type Google struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url" default:"http://localhost:8080/api/v1/test"`
	Scopes       []string `json:"scopes" default:"https://www.googleapis.com/auth/userinfo.profile,https://www.googleapis.com/auth/userinfo.email"`
}

type Postgres struct {
	Host     string `json:"host" default:"localhost"`
	Port     int    `json:"port" default:"5432"`
	Database string `json:"database" default:"shop1"`
	User     string `json:"user" default:"shopadmin"`
	Password string `json:"password"`
	Log      bool   `json:"log" default:"true"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Logger struct {
	Level string `json:"level" default:"info"`
}

type AWS struct {
	Region          string `json:"region" default:"us-east-1"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"access_key_secret"`
	S3              S3     `json:"s3"`
	SES             SES    `json:"ses"`
}

type S3 struct {
	Bucket string `json:"bucket"`
}

type SES struct {
	SourceEmail string `json:"source_email"`
}

type JWT struct {
	Secret           string `json:"secret"             envconfig:"JWT_SECRET"             default:"my-totally-secret-key"`
	AccessExpiresIn  int    `json:"access_expires_in"  envconfig:"JWT_ACCESS_EXPIRES_IN"  default:"1"`
	RefreshExpiresIn int    `json:"refresh_expires_in" envconfig:"JWT_REFRESH_EXPIRES_IN" default:"168"`
}

func New() (Config, error) {
	var cfg Config

	filename := configFilename

	flag.Parse()

	if flag.Arg(0) != "" {
		filename = flag.Arg(0)
	}

	if err := goconfig.Init(&cfg, filename); err != nil {
		return Config{}, fmt.Errorf("init config: %w", err)
	}

	config = cfg
	isInitialised = true

	return cfg, nil
}

func Get() Config {
	mutex.Lock()
	if !isInitialised {
		cfg, err := New()
		if err != nil {
			panic(err)
		}
		config = cfg
	}
	mutex.Unlock()

	return config
}
