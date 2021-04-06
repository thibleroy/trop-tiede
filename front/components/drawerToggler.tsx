import React from "react";
import {Dispatch} from "redux";
import {useDispatch} from "react-redux";
import {toggleMenu} from "../redux/actions/menuActions";
import {Button} from "@material-ui/core";

const DrawerToggler = () => {
    const dispatch: Dispatch = useDispatch();
    const toggleDrawer = (e: any) => {
        e.preventDefault()
        dispatch(toggleMenu());
    }
    return (
        <Button onClick={toggleDrawer}>Toggle Menu</Button>
    )
}

export default DrawerToggler;
