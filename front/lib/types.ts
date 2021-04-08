export interface  IResource {
    ID: string
    CreatedAt: Date
    UpdatedAt: Date
}

export interface  IRoomData {
    Temperature: number
    Time: Date
    Room_ID: string
}

export interface IDevice {
    Resource: IResource;
    DeviceDescription: IDeviceDescription;
}

export interface  IRoom {
    Resource?: IResource;
    RoomDescription: IRoomDescription;
}

export interface IRoomDescription {
    Description: IDescription;
}

export interface IDeviceDescription  {
    SerialNumber: string;
    Position?: IPosition;
    Description: IDescription
}

interface IDescription {
    Name: string;
    Details: string;
}

interface IPosition {
    Latitude: number;
    Longitude: number;
}

export interface IResponse {
    StatusCode: number;
    Result: any;
}

export interface IRoomResponse {
    Room: IRoom
}

export interface IRoomsResponse {
    Rooms: IRoom[]
}

export interface IDevicesResponse {
    Devices: IDevice[]
}

export interface IError {
    Message: string
    Code:    number
}

export interface  IEnvironment {
    WebServerPort: number
    MongoURL: string
    MongoPort: number
    JwtSecret: string
    MqttBrokerURL: string
    MqttBrokerPort: number
    MqttClientId: string
    MqttUsername: string
    MqttPassword: string
    MqttTemperatureTopic: string
}

export interface IPostReturn {
    ID: string
}

export interface  IUser {
    Resource:  IResource
    FirstName: string
    LastName:  string
    Email:     string
    Password:  string
}

export interface  IError {
    Error:   Error
    Message: string
    Code:    number
}
