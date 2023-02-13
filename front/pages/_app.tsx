import '../styles/globals.css';
import '../styles/drawer.css';
import '../styles/Home.module.css';
import '../styles/toolbar.css';
import '../styles/footer.css';
import '../styles/title.css';
import React from "react";
import { Provider } from "react-redux";
import { store } from "../redux/store";
import { AppProps } from 'next/app';
import TTAppBar from '@/components/appbar';
import TTToast from '@/components/toast';

const MyApp = ({ Component, pageProps }: AppProps) =>
(
    <Provider store={store}>
        <TTToast/>
        <TTAppBar/>
        <Component {...pageProps} />
    </Provider>
)


export default MyApp
