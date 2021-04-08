import '../styles/globals.css';
import React from "react";
import {Provider} from "react-redux";
import {store} from "../redux/store";
import MyDrawer from "../components/drawer";
import {IResponse, IRoomsResponse} from "../lib/types";
import DrawerToggler from "../components/drawerToggler";
import {AppProps} from "next/app";

interface TTAppProps extends IRoomsResponse, AppProps{}

const MyApp = ({Component, pageProps, Rooms}: TTAppProps) => {
    return (
        <Provider store={store}>
            <DrawerToggler/>
            <MyDrawer Rooms={Rooms}/>
            <Component {...pageProps} />
        </Provider>
    )
}
MyApp.getInitialProps = async (): Promise<IRoomsResponse> => {
    const res = await fetch(process.env.NEXT_PUBLIC_API + '/rooms');
    const json: IResponse = await res.json();
    return json.Result;
}
export default MyApp
