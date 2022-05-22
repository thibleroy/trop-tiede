import { IDevice, IDevicesProps } from "../lib/types";
import { List } from "@material-ui/core";
import React from "react";
import TTDevice from "./device";

const TTDevices = ({ devices }: IDevicesProps) => {
    return (
            <List>
                {devices.map((device: IDevice) => (
                    <TTDevice device={device} key={device.Resource.ID} />
                )
                )}
            </List>
    );
}

export default TTDevices;
