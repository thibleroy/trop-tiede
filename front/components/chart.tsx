import { IDeviceData } from '@/lib/types';
import {Line} from 'react-chartjs-2';
import { 
    ChartDataset, 
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend } from 'chart.js';

function Chart({data}: IDeviceData) {
    ChartJS.register(
        CategoryScale, 
        LinearScale, 
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend);
        const chartDataset: ChartDataset = {
            type: 'line',
            label: 'temperature',
            data: []
        };
        for (let d of data) {
            chartDataset.data.push({
                x: d.Time,
                y: d.Temperature
            });
        }
    const chartData = {
        datasets: [chartDataset],
        labels: data.map(data => data.Time)
    }
    return (
        <>
        <Line data={chartData} ></Line>
        </>
    )
}

export default Chart;
