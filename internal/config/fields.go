package config

import "time"

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Logger struct {
	LogInFile bool   `yaml:"logInFile"`
	OutputDir string `yaml:"outputDir"`
}

type Database struct {
	Database         string `yaml:"database"`
	PasswordFilePath string `yaml:"passwordFilePath"`
	UserName         string `yaml:"userName"`
	Url              string `yaml:"url"`
	SslMode          string `yaml:"sslMode"`

	DriverName string `yaml:"driverName"`
}

type Jwt struct {
	Iss               string        `yaml:"iss"`
	SigningFilePath   string        `yaml:"signingFilePath"`
	SessionTokenLen   int           `yaml:"sessionTokenLen"`
	SessionSigningLen int           `yaml:"sessionSigningLen"`
	AccessTokenExp    time.Duration `yaml:"accessTokenExp"`
	RefreshTokenExp   time.Duration `yaml:"refreshTokenExp"`
}

type Mail struct {
	ServerName string `yaml:"serverName"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}
