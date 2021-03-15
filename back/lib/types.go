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
	Resource IResource
	Temperature float64
	Time time.Time
	Device IDevice
}

type IDevice struct {
	Resource IResource
	DeviceDescription IDeviceDescription
	RoomID primitive.ObjectID
}

type IRoom struct {
	Resource IResource
	RoomDescription IRoomDescription
}

type IRoomDescription struct {
	Description IDescription
}

type IDeviceDescription struct {
	Description IDescription
	SerialNumber string
}

type IDescription struct {
	Name string
	Details string
}

type IRoomsResponse struct {
	Rooms []IRoom
}

type IRoomDataResponse struct {
	RoomData []IRoomData
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
	MqttTemperatureTopic string
}

type IUser struct {
	Resource  IResource
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type IError struct {
	Error   error
	Message string
	Code    int
}
