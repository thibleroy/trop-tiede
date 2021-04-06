export interface  IResource {
    ID: string
    CreatedAt: Date
    UpdatedAt: Date
}

export interface  IRoomData {
    Temperature: number
    Time: Date
}

export interface  IRoom {
    Resource?: IResource;
    RoomDescription: IRoomDescription;
}

interface IRoomDescription {
    Position: IPosition
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

export interface  IRoomsResponse {
    Rooms: IRoom[]
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
