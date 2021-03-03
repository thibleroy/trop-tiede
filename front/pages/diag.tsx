import React, {Component} from "react";
import loadable from '@loadable/component';
import {IRoomsResponse} from '../lib/types'
const ReactApexChart = loadable(() => import('react-apexcharts'), {ssr: false});
import {ApexOptions} from 'apexcharts';
const s1 = [
    {
        name: 'series-1',
        data: [30, 40, 45, 50, 49, 60, 70, 91]
    }];
interface c {
    a: ApexOptions,
    toggle: string
}
class Diag extends Component<{}, c> {
    constructor(props: any) {
        super(props);
        this.state = {
            a: {
                xaxis: {
                    categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
                },
                series: [],
            },
            toggle: 'Start'
        }
        this.toggle = this.toggle.bind(this);
    }

    async toggle() {
        const res = await fetch('http://localhost:8081/rooms');
        const rooms: IRoomsResponse = await res.json();
        const _s: number[] = []
        const _dates: Date[] =[]
        rooms.Rooms.filter((room) => {
            _s.push(room.Data.Temperature);
            _dates.push(room.Data.Time)
        })
        console.log('rooms', rooms);
            if (this.state.toggle === 'Start') {
                this.setState({
                    toggle: "Stop",
                    a: {
                        xaxis: {
                            categories: _dates
                        },
                        series: [{data: _s}]

                    }
                });
            }
            else {
                this.setState({
                    toggle: "Start",
                    a: {
                        xaxis: {
                            categories: []
                        },
                        series: []
                    }
                });
            }
    }

    render() {
        return (
            <div>
                <button onClick={this.toggle}>{this.state.toggle}</button>
                <ReactApexChart options={this.state.a} series={this.state.a.series} type="line" width={500}
                                height={320}/>
            </div>
        )
    }
}

export default Diag
