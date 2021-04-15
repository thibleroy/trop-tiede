import fetch from 'isomorphic-unfetch';
import {IData, IDevice, IDeviceData, IDeviceResponse} from "./types";
import {ConfigInterface} from "swr";
import {IDeviceDataParams} from "../pages/room/[roomId]";

export const fetchRoomData = async(input: RequestInfo, init: RequestInit): Promise<IDeviceData[]> => {
    console.log("req info", input);
    console.log("init", init);
    const deviceDataParams: IDeviceDataParams = init as IDeviceDataParams;
    // @ts-ignore
    const queryParams = new URLSearchParams(deviceDataParams.queryStrings).toString();
    const devicesDataToReturn: IDeviceData[] = [];
    for (const id of deviceDataParams.deviceIds) {
        const deviceRes = await fetch(input + id + "/temperature"+ "?" + queryParams);
        const deviceResp: IDeviceResponse = await deviceRes.json()
        const device: IDevice = deviceResp.Device;
        console.log("device", device);
    }
    return devicesDataToReturn;
    //const res = await fetch(input + id + "/temperature"+ "?" + queryParams);
    //return await res.json();
}

export const swrParams: ConfigInterface = {
    revalidateOnFocus: false,
    revalidateOnMount:false,
    revalidateOnReconnect: false,
    refreshWhenOffline: false,
    refreshWhenHidden: false,
    refreshInterval: 3000
};
