import Head from 'next/head'
import styles from '../styles/Home.module.css'
import React from "react";
import 'fontsource-roboto';
import {useRouter} from "next/router";
import {Button} from "@material-ui/core";

const Index = () => {
    const router = useRouter();
    const goHome = async (e: any) => {
        e.preventDefault();
        await router.push("/home");
    }

    return(
        <div className={styles.container}>
            <Head>
                <title>Trop tiede</title>
                <link rel="icon" href="/favicon.ico"/>
            </Head>
            <main className={styles.main}>
                <a> trop tiède </a>
            </main>
        </div>
    )
}

export default Index
