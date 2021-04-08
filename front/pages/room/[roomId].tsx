import { useRouter } from 'next/router'
import Chart from '../../components/chart';
import useSWR from "swr";
import {fetchRoomData, swrParams} from "../../lib/fetch";
import {IRoomData} from "../../lib/types";
import App from "../_app";

const buildQuery = (id: string) => {
    const today = new Date();
    const lastWeek = new Date(today.getFullYear(), today.getMonth(), today.getDate() - 7, today.getHours(), today.getMinutes(), today.getSeconds());
    const queryParams = {
        startDate: Math.round(lastWeek.getTime()/1000),
        endDate: Math.round(today.getTime()/1000)
    };
    return [process.env.NEXT_PUBLIC_API + "/room/" + id + "/temperature", queryParams];
}

const Room = () => {
    const router = useRouter();
    const { roomId } = router.query;
    const roomIdStr  = roomId as string;
    console.log("room page id", roomIdStr);
    const { data } = useSWR(buildQuery(roomIdStr), fetchRoomData, swrParams);
    const roomData: IRoomData[] = data as IRoomData[];
    console.log("data", roomData);
    return <div>
        <p>Room {roomId}</p>
        <Chart roomData={roomData}/>
    </div>;
};

export default Room
