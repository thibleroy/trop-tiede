import { SerializedError } from "@reduxjs/toolkit";
import { FetchBaseQueryError } from "@reduxjs/toolkit/dist/query";

export interface  IResource {
    ID: string;
    CreatedAt: Date;
    UpdatedAt: Date;
}

export interface IData {
    Resource?: IResource;
    Temperature: number;
    Time: number;
    DeviceId?: string;
}

export interface IDeviceData {
    device?: IDevice;
    data: IData[];
}

export interface IDeviceDataResponse {
    DeviceData: IDeviceData[];
}

export interface IDevice {
    Resource: IResource;
    DeviceDescription: IDeviceDescription;
    RoomId: string;
}

export interface IDevicesProps {
    devices: IDevice[];
}

export interface IRoomsProps {
    rooms: IRoom[];
}
export interface IDrawerItemProps {
    label: string;
    route: string;
}

export interface IDeviceProps {
    device: IDevice;
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

export interface IRoomProps {
    room: IRoom;
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
    Rooms: IRoom[];
}

export interface IDrawerProps {
    rooms: IRoom[]
}

export interface IDevicesResponse {
    Devices: IDevice[];
}

export interface IDeviceResponse {
    Device: IDevice
}

export interface IRoomProps {
    room: IRoom
}

export interface IErrorProps {
    error: FetchBaseQueryError | SerializedError;
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
