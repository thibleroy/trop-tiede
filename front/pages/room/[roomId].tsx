import {useRouter} from 'next/router'
import {
    IBody,
    IDevice,
    IDeviceData,
    IDeviceDatas,
    IDeviceResponse,
    IResponse,
    IRoom,
    IRoomResponse,
    IStatus
} from "../../lib/types";
import {GetServerSideProps, GetServerSidePropsContext} from "next";
import {List} from "@material-ui/core";
import Device from "../../components/device";
import React from "react";
import TTError from "../../components/error";
import Chart from "../../components/chart";
import useSWR from "swr";
import {fetchRoomData, swrParams} from "../../lib/fetch";
interface room extends IResponse {
    DeviceIds: string[];
}

export interface IDeviceDataParams {
    deviceIds: string[];
    queryStrings: {
        startDate: number;
        endDate: number;
    }
}

const buildQuery = (deviceIds: string[]) => {
    const today = new Date();
    const lastWeek = new Date(today.getFullYear(), today.getMonth(), today.getDate() - 7, today.getHours(),
        today.getMinutes(), today.getSeconds());
    const deviceDataParams: IDeviceDataParams = {
        deviceIds: deviceIds,
        queryStrings: {
            startDate: Math.round(lastWeek.getTime() / 1000),
            endDate: Math.round(today.getTime() / 1000)
        }
    };
    return [process.env.NEXT_PUBLIC_API + "/device/", deviceDataParams];
}

const Room = ({Body, Status, DeviceIds}: room) => {
    const {data} = useSWR(buildQuery(DeviceIds), fetchRoomData, swrParams);
    //const deviceData: IDeviceData[] = [{DeviceId: Device, Data: data}];
    //const dataToChart: IDeviceDatas = {
    //Data: deviceData
    //}
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
            {/*<List>*/}
            {/*    {devices.map((device: IDevice) => (*/}
            {/*            <Device key={device.Resource.ID}*/}
            {/*                    DeviceDescription={device.DeviceDescription}*/}
            {/*                    Resource={device.Resource}*/}
            {/*                    RoomId={roomIdStr}/>*/}
            {/*        )*/}
            {/*    )}*/}
            {/*</List>*/}
            {/*<Chart Data={}/>*/}
        </div>);
    }

};

export const getServerSideProps: GetServerSideProps = async (context: GetServerSidePropsContext) => {
    const res = await fetch(process.env.NEXT_PUBLIC_API + '/room/' + context.params?.roomId);
    const response: IBody = await res.json();
    const roomResponse: IRoomResponse = response.Value;
    const room: IRoom = roomResponse.Room;
    const devices: IDevice[] = [];
    for (const deviceId of room.DeviceIds) {
        const dRes = await fetch(process.env.NEXT_PUBLIC_API + '/device/' + deviceId);
        const dResponse: IBody = await dRes.json();
        const deviceResponse: IDeviceResponse = dResponse.Value;
        const device: IDevice = deviceResponse.Device;
        devices.push(device);
    }
    const status: IStatus = {
        Code: res.status,
        Message: response.Message
    }
    return {
        props: {
            Body: response,
            Status: status,
            devices: devices
        }
    }
}

export default Room
