import {GetServerSideProps} from "next";
import {IBody, IDevice, IDevicesResponse, IResponse} from "../lib/types";
import {List} from "@material-ui/core";
import React from "react";
import Device from "../components/device";
import TTDevices from "../components/devices";

const Devices = ({Devices}:  IDevicesResponse) => {
    console.log("devices", Devices);
    return (<div>
            <h1>Devices</h1>
            <TTDevices Devices={Devices}/>
        </div>
    );
}

export const getServerSideProps: GetServerSideProps = async() => {
    const res = await fetch(process.env.NEXT_PUBLIC_API + '/devices');
    const response: IBody = await res.json();
    console.log("resp devices", response);
    const devices: IDevicesResponse = response.Value;
    return {
        props: devices
    }
}

export default Devices;
