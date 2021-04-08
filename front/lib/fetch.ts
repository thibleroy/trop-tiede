import fetch from 'isomorphic-unfetch';
import {IRoomData} from "./types";
import {ConfigInterface} from "swr";

export const fetchRoomData = async(input: RequestInfo, init: RequestInit): Promise<IRoomData> => {
    console.log("req info", input);
    console.log("init", init);
    // @ts-ignore
    const queryParams = new URLSearchParams(init).toString();
    const res = await fetch(input + "?" + queryParams);
    return await res.json();
}

export const swrParams: ConfigInterface = {
    revalidateOnFocus: false,
    revalidateOnMount:false,
    revalidateOnReconnect: false,
    refreshWhenOffline: false,
    refreshWhenHidden: false,
    refreshInterval: 3000
};
