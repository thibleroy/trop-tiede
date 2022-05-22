import { Drawer, List } from "@material-ui/core";
import React from "react";
import { useSelector, useDispatch } from 'react-redux'
import { Dispatch } from "redux";
import { hideMenu } from "../redux/actions/menuActions";
import { RootState } from "../redux/reducers/rootReducer";
import { AppBar, Toolbar, Typography } from "@material-ui/core";
import DrawerToggler from "./drawerToggler";
import { IDrawerItemProps } from "@/lib/types";
import TTDrawerItem from "./drawerItem";

const TTDrawer = () => {
    const drawerState = useSelector((state: RootState) => state.menu);
    const dispatch: Dispatch = useDispatch();

    const handleClose = () => {
        dispatch(hideMenu());
    };

    const clickAwayHandler = (e: React.MouseEvent<Document>, reason: "backdropClick" | "escapeKeyDown") => {
        console.log('reason', reason);
        e.preventDefault();
        handleClose();
    }

    const routes: IDrawerItemProps[] =
        [
            {
                label: 'Dashboard',
                route: '/'
            },
            {
                label: 'Home',
                route: '/home',
            },
            {
                label: 'Rooms',
                route: '/rooms'
            },
            {
                label: 'Devices',
                route: '/devices'
            }
        ];

    return (
        <>
            <DrawerToggler />
            <Drawer open={drawerState.open} onClose={clickAwayHandler}>
                <AppBar position="static">
                    <Toolbar>
                        <DrawerToggler />
                        <Typography variant="h4">
                            Trop tiède
                        </Typography>
                    </Toolbar>
                </AppBar>
                <List>
                    {routes.map((drawerItemProps: IDrawerItemProps, id: number) => (
                        <TTDrawerItem key={id} label={drawerItemProps.label} route={drawerItemProps.route} />
                    ))}
                </List>
            </Drawer>
        </>
    )
}
export default TTDrawer;
