package lib

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
	"time"
)

func dotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env.local file")
	}
	return os.Getenv(key)
}

func GetServerEnv() IEnvironment {
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

	var serverEnvironment IEnvironment
	wsPort, _ := strconv.Atoi(wspStr)
	mongoPort, _ := strconv.Atoi(mongoPortStr)
	mqttPort, _ := strconv.Atoi(mqttPortStr)
	serverEnvironment = IEnvironment{
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

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func NewResource() IResource {
	creationTime := time.Now()
	return IResource{
		ID:        primitive.NewObjectIDFromTimestamp(creationTime),
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
	}
}

func GenerateJWT(secret string)(string,error){
	token:= jwt.New(jwt.SigningMethodHS256)
	tokenString, err :=  token.SignedString(secret)
	if err !=nil{
		log.Println("Error in JWT token generation")
		return "",err
	}
	return tokenString, nil
}
