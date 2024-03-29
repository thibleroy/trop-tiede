package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IEnvironment struct {
	RabbitMQBrokerUrl      string
	RabbitMQBrokerPort     string
	RabbitMQBrokerUsername string
	RabbitMQBrokerPassword string
}

func dotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetServerEnv() IEnvironment {
	return IEnvironment{
		RabbitMQBrokerUrl:      dotEnvVariable("RABBITMQ_BROKER_URL"),
		RabbitMQBrokerPort:     dotEnvVariable("RABBITMQ_BROKER_PORT"),
		RabbitMQBrokerUsername: dotEnvVariable("RABBITMQ_BROKER_USERNAME"),
		RabbitMQBrokerPassword: dotEnvVariable("RABBITMQ_BROKER_PASSWORD"),
	}
}
