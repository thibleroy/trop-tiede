import {IDevice} from "../lib/types";
import {List} from "@material-ui/core";
import Device from "./device";
import React from "react";

const TTDevices = (devices:  IDevice[]) => {
    console.log("devices", devices);
    return (
        <>
            <List>
                {devices.map((device: IDevice) => (
                        <Device key={device.Resource.ID}
                                DeviceDescription={device.DeviceDescription}
                                Resource={device.Resource}
                                RoomId={device.RoomId}
                        />
                    )
                )}
            </List>
            </>
    );
}

export default TTDevices;
