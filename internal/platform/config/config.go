package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	DatabaseURL              string `env:"DATABASE_URL" validate:"required"`
	JwtSecret                string `env:"JWT_SECRET" validate:"required"`
	JwtExpires               string `env:"JWT_EXPIRES" validate:"required"`
	RedisUrl                 string `env:"REDIS_URL" validate:"required"`
	Port                     string `env:"PORT" validate:"required"`
	EncryptionSecretKey      string `env:"ENCRYPTION_SECRET_KEY" validate:"required"`
	EncryptionSecretIV       string `env:"ENCRYPTION_SECRET_IV" validate:"required"`
	EncryptionMethod         string `env:"ENCRYPTION_METHOD" validate:"required"`
	AppEnv                   string `env:"APP_ENV" validate:"required"`
	PaystackSecretKey        string `env:"PAYSTACK_SECRET_KEY" validate:"required"`
	PaystackBaseURL          string `env:"PAYSTACK_BASE_URL" validate:"required"`
	EmailAPIKey              string `env:"EMAIL_API_KEY" validate:"required"`
	MonoSecretKey            string `env:"MONO_SECRET_KEY" validate:"required"`
	MonoBaseURL              string `env:"MONO_BASE_URL" validate:"required"`
	ResendAPIKey             string `env:"RESEND_API_KEY" validate:"required"`
	ResendAPIURL             string `env:"RESEND_API_URL" validate:"required"`
	KafkaBroker              string `env:"KAFKA_BROKER" validate:"required"`
	KafkaClientID            string `env:"KAFKA_CLIENT_ID" validate:"required"`
	KafkaConsumerGroupPrefix string `env:"KAFKA_CONSUMER_GROUP_PREFIX" validate:"required"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
