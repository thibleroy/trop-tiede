import Head from 'next/head'
import styles from '../styles/Home.module.css'
import React from "react";
import Diag from "./diag"

const a = "Ifécho";
export default function Home() {
    return (
        <div className={styles.container}>
            <Head>
                <title>{a}</title>
                <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico"/>
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    <a> trop tiède </a>
                </h1>

                <div className="mixed-chart">
                    <Diag/>
                </div>
            </main>
        </div>
    )
}
