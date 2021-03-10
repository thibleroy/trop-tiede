import React from "react";
import {Dispatch} from "redux";
import {useDispatch} from "react-redux";
import {toggleMenu} from "../redux/actions/menuActions";

const DrawerToggler = () => {
    const dispatch: Dispatch = useDispatch();
    const toggleDrawer = (e: any) => {

        e.preventDefault()
        dispatch(toggleMenu());
    }
    return (
        <button onClick={toggleDrawer}> Toggle Menu</button>
    )
}

export default DrawerToggler;
