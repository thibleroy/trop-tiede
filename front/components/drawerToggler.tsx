import React from "react";
import { Dispatch } from "redux";
import { useDispatch } from "react-redux";
import { toggleMenu } from "../redux/actions/menuActions";
import { IconButton } from "@material-ui/core";
import {Menu} from "@mui/icons-material";

const DrawerToggler = () => {
    const dispatch: Dispatch = useDispatch();
    const toggleDrawer = (e: any) => {
        e.preventDefault()
        dispatch(toggleMenu());
    }
    return (
        <IconButton aria-label="menu" onClick={toggleDrawer} color="inherit" edge="start">
            <Menu/>
        </IconButton>
    )
}

export default DrawerToggler;
