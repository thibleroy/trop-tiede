package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IRoomData struct {
	Resource IResource
	Temperature float64
	Time time.Time
	Device IDevice
	Room IRoom
}

type IDevice struct {
	Resource IResource
	DeviceDescription IDeviceDescription
}

type IRoom struct {
	Resource IResource
	RoomDescription IRoomDescription
}

type IRoomDescription struct {
	Position IPosition
	Description IDescription
}

type IDeviceDescription struct {
	Description IDescription
	SerialNumber string
}

type IRoomsResponse struct {
	Rooms []IRoom
}

type IRoomDataResponse struct {
	RoomData []IRoomData
}

type IPosition struct {
	Latitude float64
	Longitude float64
}

type IDescription struct {
	Name string
	Details string
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

type IResource struct {
	ID primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}
