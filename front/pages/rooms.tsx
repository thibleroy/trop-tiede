import React from "react";
import { useGetRoomsQuery } from "redux/middlewares/api/rooms";
import TTRooms from "@/components/rooms";
import TTError from "@/components/error";
import TTLoader from "@/components/loader";
import { Divider, Typography } from "@material-ui/core";

const Rooms = () => {
    const { data, error, isLoading } = useGetRoomsQuery();
    return (<>
        <Typography variant="h3">
            Rooms
        </Typography>

        <Divider />

        {error ? (<TTError error={error} />) :
            isLoading ? (<TTLoader />) :
                data ?
                    (<TTRooms rooms={data.Rooms} />)
                    : null}
    </>
    );
}

export default Rooms;
