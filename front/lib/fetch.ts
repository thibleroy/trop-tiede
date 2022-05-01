import { IRoom, IBody } from "./types";

export const retrieveRooms = async(): Promise<IRoom[]> => {
    const res = await fetch(process.env.NEXT_PUBLIC_API + '/rooms');
    const body: IBody = await res.json();
    return body.Value.Rooms;
}

