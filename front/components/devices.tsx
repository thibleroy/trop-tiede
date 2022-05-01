import {IDevice, IDevicesProps} from "../lib/types";
import {List} from "@material-ui/core";
import Device from "./device";
import React from "react";

const TTDevices = ({devices}: IDevicesProps) => {
    console.log("devices", devices);
    return (
        <>
            <List>
                {devices.map((device: IDevice) => (
                        <Device device={device} key={device.Resource.ID}/>
                    )
                )}
            </List>
            </>
    );
}

export default TTDevices;
