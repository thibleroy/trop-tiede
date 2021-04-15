export interface  IResource {
    ID: string;
    CreatedAt: Date;
    UpdatedAt: Date;IData
}

export interface IData {
    Resource: IResource;
    Temperature: number;
    Time: Date;
    DeviceId: string;
}

export interface IDeviceData {
    DeviceId: IDevice;
    Data: IData[];
}

export interface IDeviceDatas {
    Data: IDeviceData[];
}

export interface IDeviceDataResponse {
    DeviceData: IDeviceData[];
}

export interface IDevice {
    Resource: IResource;
    DeviceDescription: IDeviceDescription;
    RoomId: string;
}

export interface  IRoom {
    Resource?: IResource;
    RoomDescription: IRoomDescription;
    DeviceIds: string[];
}

export interface IRoomDescription {
    Description: IDescription;
}

export interface IDeviceDescription  {
    SerialNumber: string;
    Position?: IPosition;
    Description: IDescription;
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
    Status: IStatus
    Body: IBody
    Headers: IHeader[];
}

export interface IBody {
    Value: any;
    Message: string;
}

export interface IHeader {
    Key: string;
    Value: string;
}

export interface IStatus {
    Message: string;
    Code:    number;
}

export interface IRoomResponse {
    Room: IRoom;
}

export interface IRoomsResponse {
    Rooms: IRoom[]
}

export interface IDevicesResponse {
    Devices: IDevice[];
}

export interface IDeviceResponse {
    Device: IDevice
}

export interface IError {
    Message: string;
    Code:    number;
}

export interface  IEnvironment {
    WebServerPort: number;
    MongoURL: string;
    MongoPort: number;
    JwtSecret: string;
    MqttBrokerURL: string;
    MqttBrokerPort: number;
    MqttClientId: string;
    MqttUsername: string;
    MqttPassword: string;
    MqttTemperatureTopic: string;
}

export interface IPostReturn {
    ID: string;
}

export interface  IUser {
    Resource:  IResource;
    FirstName: string;
    LastName:  string;
    Email:     string;
    Password:  string;
}

export interface  IError {
    Message: string;
    Code:    number;
}
