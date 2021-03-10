import React from "react";
import {useRouter} from "next/router";
import {Dispatch} from "redux";
import {useDispatch} from "react-redux";
import {hideMenu} from "../redux/actions/menuActions"

const HomeBtn = () => {
    const router = useRouter();
    const dispatch: Dispatch = useDispatch();
    const goHome = async (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault()
        await router.push("/home");
        dispatch(hideMenu());
    }

    return (
        <button onClick={goHome}>Go Home</button>
    )
}
export default HomeBtn;
