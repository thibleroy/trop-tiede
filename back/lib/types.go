package lib

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IRoomData struct {
	Resource    IResource
	Temperature float64
	Time        time.Time
	RoomId      primitive.ObjectID
	DeviceId 	primitive.ObjectID
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
	Description IDescription
}

type IDeviceDescription struct {
	Position IPosition
	SerialNumber string
	Description IDescription
}

type IResponse struct {
	StatusCode int
	Result interface{}
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
	Message string
	Code    int
}

type IRoomResponse struct {
	Room IRoom
}

type IRoomsResponse struct {
	Rooms []IRoom
}

type IDeviceResponse struct {
	Device IDevice
}

type IDevicesResponse struct {
	Devices []IDevice
}

type IResource struct {
	ID primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}
