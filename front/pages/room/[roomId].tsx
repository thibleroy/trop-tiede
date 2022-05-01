import {
    IBody,
    IDevicesResponse,
    IStatus,
    IRoomProps
} from "../../lib/types";
import {GetServerSideProps, GetServerSidePropsContext} from "next";
import TTDevices from "../../components/devices";
import React from "react";

const Room = ({room, devices}: IRoomProps) => {
    console.log("devices room", devices)
    console.log("room", room)
        return (<div>
            <h1>Room {room.Resource?.ID}</h1>
            <h2>{room.RoomDescription.Description.Name}</h2>
            <h2>{room.RoomDescription.Description.Details}</h2>
            <h2 >Devices</h2>
            <TTDevices devices={devices}/>
        </div>);
};

export const getServerSideProps: GetServerSideProps = async (context: GetServerSidePropsContext) => {
    const res1 = await fetch(process.env.NEXT_PUBLIC_API + '/room/' + context.params?.roomId);
    const response: IBody = await res1.json();
    const room = response.Value.Room;
    const res2 = await fetch(process.env.NEXT_PUBLIC_API + '/devices?roomId=' + context.params?.roomId);
    const response2: IBody = await res2.json();
    const devicesResponse: IDevicesResponse = response2.Value;
    const devices = devicesResponse.Devices;
    return {
        props: {
            room,
            devices
        }
    };
}

export default Room
