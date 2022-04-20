package lib

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDevice struct {
	Resource          IResource
	DeviceDescription IDeviceDescription
	RoomId            primitive.ObjectID
}

type IDeviceData struct {
	Resource    IResource
	Temperature float64
	Time        time.Time
	DeviceId    primitive.ObjectID
}

type IRoom struct {
	Resource        IResource
	RoomDescription IRoomDescription
}

type IID struct {
	Id string
}

type IRoomDescription struct {
	Description IDescription
}

type IDeviceDescription struct {
	Position     IPosition
	SerialNumber string
	Description  IDescription
}

type IResponse struct {
	Status  IStatus
	Body    IBody
	Headers []IHeader
}

type IBody struct {
	Value   interface{}
	Message string
}

type IHeader struct {
	Key   string
	Value string
}

type IDeviceDataResponse struct {
	DeviceData []IDeviceData
}

type IPosition struct {
	Latitude  float64
	Longitude float64
}

type IDescription struct {
	Name    string
	Details string
}

type IEnvironment struct {
	WebServerPort  int
	MongoURL       string
	MongoPort      int
	JwtSecret      string
	MqttBrokerURL  string
	MqttBrokerPort int
	MqttClientId   string
	MqttUsername   string
	MqttPassword   string
}

type IUser struct {
	Resource  IResource
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type IStatus struct {
	Message    string
	Code       int
	RawMessage string
}

type IRoomResponse struct {
	Room IRoom
}

type IRoomsResponse struct {
	Rooms  []IRoom
	Length int
}

type IDeviceResponse struct {
	Device IDevice
}

type IDevicesResponse struct {
	Devices []IDevice
	Length  int
}

type IUserResponse struct {
	User IUser
}

type IUsersResponse struct {
	Users  []IUser
	Length int
}

type IResource struct {
	ID        primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}
