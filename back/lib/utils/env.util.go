package utils

import (
	"back/lib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func dotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetServerEnv() lib.IEnvironment {
	wspStr := dotEnvVariable("WEBSERVER_PORT")
	mongoURLStr := dotEnvVariable("MONGODB_URL")
	mongoPortStr := dotEnvVariable("MONGODB_PORT")
	jwtSecretStr := dotEnvVariable("JWT_SECRET")
	mqttURLStr := dotEnvVariable("MQTT_URL")
	mqttPortStr := dotEnvVariable("MQTT_PORT")
	mqttClientIdStr := dotEnvVariable("MQTT_CLIENTID")
	mqttUsernameStr := dotEnvVariable("MQTT_USERNAME")
	mqttPasswordStr := dotEnvVariable("MQTT_PASSWORD")
	mqttTemperatureTopic := dotEnvVariable("MQTT_TEMP_TOPIC")

	var serverEnvironment lib.IEnvironment
	wsPort, _ := strconv.Atoi(wspStr)
	mongoPort, _ := strconv.Atoi(mongoPortStr)
	mqttPort, _ := strconv.Atoi(mqttPortStr)
	serverEnvironment = lib.IEnvironment{
		WebServerPort: wsPort,
		MongoURL:      mongoURLStr,
		MongoPort:     mongoPort,
		JwtSecret:     jwtSecretStr,
		MqttBrokerURL: mqttURLStr,
		MqttBrokerPort: mqttPort,
		MqttClientId: mqttClientIdStr,
		MqttUsername: mqttUsernameStr,
		MqttPassword: mqttPasswordStr,
		MqttTemperatureTopic: mqttTemperatureTopic,
	}
	return serverEnvironment
}
