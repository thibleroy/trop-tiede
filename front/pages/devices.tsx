import {Divider, Typography} from "@material-ui/core";
import React from "react";
import TTDevices from "../components/devices";
import TTError from "@/components/error";
import TTLoader from "@/components/loader";
import { useGetDevicesQuery } from "redux/middlewares/api/devices";

const Devices = () => {
    const { data, error, isLoading } = useGetDevicesQuery();

    return (<>
        <Typography variant="h3">
            Devices
        </Typography>

        <Divider />

        {error ? (<TTError error={error} />) :
            isLoading ? (<TTLoader />) :
                data ?
                    (<TTDevices devices={data.Devices} />)
                    : null}
    </>
    );
}

export default Devices;
