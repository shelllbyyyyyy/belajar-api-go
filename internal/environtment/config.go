package environtment

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	App AppConfig 
	DB  DBConfig  
}

type AppConfig struct {
	Name       string          
	Port       string           
	Host       string
	DomainName string          
	Encryption EncryptionConfig 
}

type EncryptionConfig struct {
	Salt      uint8  
	JWTSecret string 
}

type DBConfig struct {
	Host           string                
	Port           string                 
	User           string                 
	Password       string                 
	Name           string                 
	ConnectionPool DBConnectionPoolConfig 
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 
	MaxOpenConnetcion     uint8 
	MaxLifetimeConnection uint8 
	MaxIdletimeConnection uint8
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Port: os.Getenv("PORT"),
			Host: os.Getenv("HOST"),
			DomainName: os.Getenv("DOMAIN_NAME"),
			Encryption: EncryptionConfig{
				Salt: 10,
				JWTSecret: os.Getenv("JWT_SECRET"),
			},
		},
		DB: DBConfig{
			Host: os.Getenv("DATABASE_HOST"),
			Port: os.Getenv("DATABASE_PORT"),
			User: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name: os.Getenv("DATABASE_NAME"),
			ConnectionPool: DBConnectionPoolConfig{
				MaxIdleConnection: 10,
				MaxOpenConnetcion: 30,
				MaxLifetimeConnection: 60,
				MaxIdletimeConnection: 60,
			},
		},
	}

	return cfg, nil
}