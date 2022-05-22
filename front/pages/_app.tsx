import '../styles/globals.css';
import React from "react";
import { Provider } from "react-redux";
import { store } from "../redux/store";
import { AppProps } from 'next/app';
import TTAppBar from '@/components/appbar';


const MyApp = ({ Component, pageProps }: AppProps) => {
    return (
        <Provider store={store}>
            <TTAppBar />
            <Component {...pageProps} />
        </Provider>
    )
}

export default MyApp
