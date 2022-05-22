import React from "react";
import { useGetRoomQuery } from "redux/middlewares/api/rooms";
import { useRouter } from "next/router";
import { Divider, Typography } from "@material-ui/core";
import TTError from "@/components/error";
import TTLoader from "@/components/loader";

const Room = () => {
    const router = useRouter();
    const {roomId} = router.query;
    const { data, error, isLoading } = useGetRoomQuery(roomId as string, {skip: roomId === '' || roomId === undefined});
    return (<>
        <Typography variant="h3">
            Room
        </Typography>
        <Divider />

        {error ? (<TTError error={error} />) :
            isLoading ? (<TTLoader />) :
                data ?
                    (<> {JSON.stringify(data.Room)} </>)
                    : null}
    </>
    );


};

export default Room
