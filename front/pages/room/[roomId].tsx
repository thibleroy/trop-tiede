import {useRouter} from 'next/router'
import {
    IBody,
    IDevice,
    IDevicesResponse,
    IResponse,
    IRoom,
    IRoomResponse,
    IStatus
} from "../../lib/types";
import {GetServerSideProps, GetServerSidePropsContext} from "next";
import TTDevices from "../../components/devices";
import React from "react";
import TTError from "../../components/error";
interface room extends IResponse {
    Devices: IDevice[];
}

export interface IDeviceDataParams {
    deviceIds: string[];
    queryStrings: {
        startDate: number;
        endDate: number;
    }
}

const Room = ({Body, Status, Devices}: room) => {
    const deviceIds = Devices.map((device) => {return device.Resource.ID})
    console.log("deviceIds", deviceIds)
    if (400 <= Status.Code && Status.Code <= 405) {
        return <TTError Code={Status.Code} Message={Status.Message}/>
    } else {
        const router = useRouter();
        const {roomId} = router.query;
        const roomIdStr = roomId as string
        const roomResponse: IRoomResponse = Body.Value;
        const room: IRoom = roomResponse.Room;
        return (<div>
            <h1>Room</h1>
            <h2>Id</h2>{roomId}
            <h2>Name</h2>{room.RoomDescription.Description.Name}
            <h2>Description</h2>{room.RoomDescription.Description.Details}
            <h2>Devices</h2>
            <TTDevices {Devices}/>
        </div>);
    }
};

export const getServerSideProps: GetServerSideProps = async (context: GetServerSidePropsContext) => {
    const res1 = await fetch(process.env.NEXT_PUBLIC_API + '/room/' + context.params?.roomId);
    const response: IBody = await res1.json();
    const roomResponse: IRoomResponse = response.Value;
    const room: IRoom = roomResponse.Room;
    const res2 = await fetch(process.env.NEXT_PUBLIC_API + '/devices?roomId=' + context.params?.roomId);
    const response2: IBody = await res2.json();
    const devicesResponse: IDevicesResponse = response2.Value;
    const devices = devicesResponse.Devices;
    const status: IStatus = {
        Code: res1.status,
        Message: response.Message
    };
    console.log("devices", devices);
    return {
        props: {
            Body: response,
            Status: status,
            Devices: devices
        }
    };
}

export default Room
