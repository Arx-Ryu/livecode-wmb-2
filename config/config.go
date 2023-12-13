package config

import (
	"fmt"
	"os"
	"time"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type GrpcConfig struct {
	Url string
}

type ApiConfig struct {
	Url string
}

type DBConfig struct {
	DataSourceName string
}

type TokenConfig struct {
	ApplicationName  		string
	JwtSigningMethod 		*jwt.SigningMethodHMAC
	JwtSignatureKey 		string
	AccessTokenLifeTime 	time.Duration
	Client 					*redis.Client
}

type Config struct {
	ApiConfig
	DBConfig
	GrpcConfig
	TokenConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.DBConfig = DBConfig{DataSourceName: dsn}

	c.ApiConfig = ApiConfig{Url: api}

	grpcUrl := os.Getenv("GRPC_URL")
	c.GrpcConfig = GrpcConfig{Url: grpcUrl}

	c.TokenConfig = TokenConfig{
		ApplicationName: "ENIGMA",
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtSignatureKey: "3N!GM4",
		AccessTokenLifeTime: 60*time.Second,
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		}),
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}

/*
//CMD TERMINAL, NOT POWERSHELL
set DB_HOST=localhost
set DB_USER=postgres
set DB_PASSWORD=h0twh33l5
set DB_NAME=livecode_wmb_rest_api
set DB_PORT=5432

set API_URL=localhost:8080
set GRPC_URL=localhost:8000

set ENV=dev
set ENV=migration
*/