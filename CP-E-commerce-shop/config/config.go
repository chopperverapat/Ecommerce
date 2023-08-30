package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// --------------------Load Config-----------------------//

func LoadConfig(path string) Iconfig {
	envConfig, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("failed to get ENV , error : %v", err)
	}
	return &config{
		app: &app{
			host: envConfig["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envConfig["APP_PORT"])
				if err != nil {
					log.Fatalf("failed to get port , error : %v", err)
				}
				return p
			}(),
			name:    envConfig["APP_NAME"],
			version: envConfig["APP_VERSION"],
			readTimeout: func() time.Duration {
				r, err := strconv.Atoi(envConfig["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("failed get read time out , error : %v", err)
				}
				return time.Duration(int64(r) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				w, err := strconv.Atoi(envConfig["APP_WRTIE_TIMEOUT"])
				if err != nil {
					log.Fatalf("failed to get wrtie time outerror : %v", err)
				}
				return time.Duration(int64(w) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envConfig["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("failed to get bofu limit")
				}
				return b
			}(),
			fileLimit: func() int {
				b, err := strconv.Atoi(envConfig["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("failed get file limit : %v", err)
				}
				return b
			}(),
			gcpbucket: envConfig["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envConfig["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envConfig["DB_PORT"])
				if err != nil {
					log.Fatalf("failed get port db , error : %v", err)
				}
				return p
			}(),
			protocol: envConfig["DB_PROTOCOL"],
			username: envConfig["DB_USERNAME"],
			password: envConfig["DB_PASSWORD"],
			databse:  envConfig["DB_DATABASE"],
			sslMode:  envConfig["DB_SSL_MODE"],
			maxConnection: func() int {
				m, err := strconv.Atoi(envConfig["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("failed get port max connection , error : %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			adminkey: envConfig["JWT_ADMIN_KEY"],
			secret:   envConfig["JWT_SECRET_KEY"],
			apikeys:  envConfig["JWT_API_KEY"],
			accessExpiredAt: func() int {
				a, err := strconv.Atoi(envConfig["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("failed get port max connection , error : %v", err)
				}
				return a

			}(),
			refreshExpiredAt: func() int {
				r, err := strconv.Atoi(envConfig["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("failed get port max connection , error : %v", err)
				}
				return r

			}(),
		},
	}
}

// --------------------CONFIG-----------------------//
type config struct {
	app *app
	db  *db
	jwt *jwt
}

type Iconfig interface {
	App() Iapp
	Db() Idb
	Jwt() Ijwt
}

// --------------------APP-----------------------//
type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //byte
	fileLimit    int //byte
	gcpbucket    string
}

type Iapp interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
	Host() string
	Port() int
}

func (c *config) App() Iapp {
	return c.app
}

func (a *app) Url() string                 { return fmt.Sprintf("%v:%v", a.host, a.port) }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.writeTimeout }
func (a *app) WriteTimeout() time.Duration { return a.readTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GCPBucket() string           { return a.gcpbucket }
func (a *app) Host() string                { return a.host }
func (a *app) Port() int                   { return a.port }

// --------------------DB-----------------------//
type db struct {
	host          string
	port          int
	protocol      string
	username      string
	password      string
	databse       string
	sslMode       string
	maxConnection int
}

type Idb interface {
	Url() string
	MaxOpenConns() int
}

func (c *config) Db() Idb {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		d.host,
		d.port,
		d.username,
		d.password,
		d.databse,
		d.sslMode,
	)
}

func (d *db) MaxOpenConns() int { return d.maxConnection }

// --------------------JWT-----------------------//
type jwt struct {
	adminkey         string
	secret           string
	apikeys          string
	accessExpiredAt  int
	refreshExpiredAt int
}

type Ijwt interface {
	SecretKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetAccessExpiresAt(t int) int
	SetRefreshExpiresAt(t int) int
}

func (c *config) Jwt() Ijwt {
	return c.jwt
}

func (j *jwt) SecretKey() []byte     { return []byte(j.secret) }
func (j *jwt) AdminKey() []byte      { return []byte(j.adminkey) }
func (j *jwt) ApiKey() []byte        { return []byte(j.apikeys) }
func (j *jwt) AccessExpiresAt() int  { return j.accessExpiredAt }
func (j *jwt) RefreshExpiresAt() int { return j.refreshExpiredAt }
func (j *jwt) SetAccessExpiresAt(jin int) int {
	j.accessExpiredAt = jin
	return j.accessExpiredAt
}
func (j *jwt) SetRefreshExpiresAt(jin int) int {
	j.refreshExpiredAt = jin
	return j.refreshExpiredAt
}
