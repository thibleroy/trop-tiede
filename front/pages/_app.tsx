import '../styles/globals.css'
import React from "react";
import {Provider} from "react-redux";
import {store} from "../redux/store";
import MyDrawer from "../components/drawer";
import {IRoomsResponse} from "../lib/types";
import DrawerToggler from "../components/drawerToggler";


// @ts-ignore
 const MyApp = ({ Component, pageProps, Rooms }) => {
     return (
             <Provider store={store} >
                 <DrawerToggler/>
                 <MyDrawer Rooms={Rooms}/>
                 <Component {...pageProps} />
             </Provider>
     )
 }
MyApp.getInitialProps = async () => {
    const res = await fetch(process.env.NEXT_PUBLIC_WEBSERVER_URL + '/rooms');
    const rooms: IRoomsResponse = await res.json();
    console.log("rooms", rooms)
    return rooms
}
export default MyApp
