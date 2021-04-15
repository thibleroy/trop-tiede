import loadable from "@loadable/component";
import {IData, IDeviceData, IDeviceDatas} from "../lib/types";
import {ApexOptions} from "apexcharts";
const ReactApexChart = loadable(() => import('react-apexcharts'), {ssr: false});

function Chart({Data}: IDeviceDatas) {
    console.log("data in chart component", Data);
    const chartOptions: ApexOptions = {title: {text: "Temperature"}, series: []};
    const series: ApexAxisChartSeries = [];
    if (Data) {
        Data.forEach((deviceData: IDeviceData) => {
            const _s: number[] = [];
            deviceData.Data.forEach((data: IData) => {
                _s.push(data.Temperature);
                chartOptions.xaxis?.categories.push(data.Time);
            });
            series.push({data: _s, name: deviceData.Device.DeviceDescription.Description.Name});
        });
    }
    chartOptions.series = series;
    return(
        <div>
            <ReactApexChart options={chartOptions} series={chartOptions.series} type="line" width={500}
                            height={320}/>
        </div>
        );
}

export default Chart;
