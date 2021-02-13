package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IResource struct {
	ID primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IRoomData struct {
	Temperature int
}

type IRoom struct {
	Resource IResource
	Name string
	Description string
	Data IRoomData
}
type IEnvironment struct {
	WebServerPort int
	MongoURL string
	MongoPort int
	JwtSecret string
	MqttBrokerURL string
	MqttBrokerPort int
	MqttClientId string
	MqttUsername string
	MqttPassword string
}

type IPostReturn struct {
	ID primitive.ObjectID
}

type IUser struct {
	Resource IResource
	FirstName string
	LastName string
	Email string
	Password string
}

type IError struct {
	Error   error
	Message string
	Code    int
}
