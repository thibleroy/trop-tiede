import {IDeviceData, IDeviceProps} from "../lib/types";
import {ListItem, ListItemText} from "@material-ui/core";
import Chart from './chart';
export interface IDeviceDataParams {
    deviceIds: string[];
    queryStrings: {
        startDate: number;
        endDate: number;
    }
}
const TTDevice = ({device}: IDeviceProps) => {

    const data: IDeviceData = {
        device,
        data: [{Temperature: 21, Time: 1},
            {Temperature: 23, Time: 3},
            {Temperature: 25, Time: 5}]
    }
    return (<div>
            <ListItem divider>
                <ListItemText primary={device.DeviceDescription.Description.Name}
                              secondary={device.DeviceDescription.Description.Details}/>
                <ListItemText secondary={device.DeviceDescription.SerialNumber}/>
            </ListItem>
            <Chart data={data.data}></Chart>
    </div>
    )
}

export default TTDevice;
