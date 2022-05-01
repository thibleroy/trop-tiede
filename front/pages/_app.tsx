import '../styles/globals.css';
import React from "react";
import {Provider} from "react-redux";
import {store} from "../redux/store";
import TTDrawer from "../components/drawer";
import DrawerToggler from "../components/drawerToggler";
import { AppProps } from 'next/app';


const MyApp = ({Component, pageProps}: AppProps) => {
    return (
        <Provider store={store}>
            <DrawerToggler/>
            <TTDrawer/>
            <Component {...pageProps} />
        </Provider>
    )
}

export default MyApp
