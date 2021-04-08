import {GetServerSideProps} from "next";
import {IDevice, IDevicesResponse, IResponse} from "../lib/types";
import {List} from "@material-ui/core";
import React from "react";
import Device from "../components/device";

const Devices = ({devices}: {devices: IDevicesResponse}) => {
    return (<div>
            <h1>Devices</h1>
            <List>
                {devices.Devices.map((device: IDevice) => (
                   <Device key={device.Resource.ID}
                           DeviceDescription={device.DeviceDescription}
                           Resource={device.Resource}/>
                    )
                )}
            </List>
        </div>
    );
}

export const getServerSideProps: GetServerSideProps = async() => {
    const res = await fetch(process.env.NEXT_PUBLIC_API + '/devices');
    const response: IResponse = await res.json();
    console.log("resp devices", response);
    const devices: IDevicesResponse = response.Result;
    return {
        props: {
            devices
        }
    }
}

export default Devices;
