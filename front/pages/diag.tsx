import {GetServerSideProps, InferGetServerSidePropsType} from 'next'
import {IRoomsResponse} from "../lib/types";
import loadable from "@loadable/component";
const ReactApexChart = loadable(() => import('react-apexcharts'), {ssr: false});

export const getServerSideProps: GetServerSideProps = async () => {
    const res = await fetch(process.env.NEXT_PUBLIC_WEBSERVER_URL + '/rooms');
    const datar: IRoomsResponse = await res.json();
    const _s: number[] = []
    const _dates: Date[] =[]
    datar.Rooms.filter((room) => {
        _s.push(room.Data.Temperature);
        _dates.push(room.Data.Time)
    })
    return {
        props: {
            rooms: [{data:  _s}],
            dates: {
                xaxis: {
                    categories: _dates
                }
            }
        },
    }
}

function Diag({rooms, dates}: InferGetServerSidePropsType<typeof getServerSideProps>) {
    return(
        <div>
            <ReactApexChart options={dates} series={rooms} type="line" width={500}
                            height={320}/>
        </div>
        );
}

export default Diag;
