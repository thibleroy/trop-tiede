import { IRoom, IRoomsProps } from "../lib/types";
import { List } from "@material-ui/core";
import React from "react";
import TTRoom from "./room";

const TTRooms = ({ rooms }: IRoomsProps) => {
    return (
        <List>
            {rooms.map((room: IRoom) => (
                <TTRoom room={room} key={room.Resource?.ID} />
            )
            )}
        </List>
    );
}

export default TTRooms;
