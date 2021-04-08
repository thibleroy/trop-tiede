import {IDevice} from "../lib/types";
import {ListItem, ListItemText} from "@material-ui/core";

const Device = (Device: IDevice) => {
    return (
        <ListItem>
            <ListItemText primary={Device.DeviceDescription.Description.Name}
                          secondary={Device.DeviceDescription.Description.Details}/>
            <ListItemText secondary={Device.DeviceDescription.SerialNumber}/>
        </ListItem>
    )
}

export default Device;
