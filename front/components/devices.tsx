import {IDevice, IDevicesResponse} from "../lib/types";
import {List} from "@material-ui/core";
import Device from "./device";
import React from "react";

const TTDevices = ({Devices}:  IDevicesResponse) => {
    console.log("devices", Devices);
    return (<>
            <List>
                {Devices.map((device: IDevice) => (
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
