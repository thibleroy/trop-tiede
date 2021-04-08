import loadable from "@loadable/component";
import {IRoomData} from "../lib/types";
import {ApexOptions} from "apexcharts";

const ReactApexChart = loadable(() => import('react-apexcharts'), {ssr: false});

function Chart({roomData}: {roomData: IRoomData[]}) {
    console.log("data in chart component", roomData);
    const _s: number[] = [];
    const chartOptions: ApexOptions = {title: {text: "Temperature"}};
    chartOptions.series = [{data: [], name: "Temperature"}];
    if (roomData) {
        roomData.filter((dataValue: IRoomData) => {
            _s.push(dataValue.Temperature);
            chartOptions.xaxis?.categories.push(dataValue.Time);
        });
        chartOptions.series = [{data: _s, name: "Temperature"}];
    }
    return(
        <div>
            <ReactApexChart options={chartOptions} series={chartOptions.series} type="line" width={500}
                            height={320}/>
        </div>
        );
}

export default Chart;
