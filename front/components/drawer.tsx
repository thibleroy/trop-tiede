import {Drawer, ListItem, List, ClickAwayListener} from "@material-ui/core";
import React from "react";
import {useSelector, useDispatch} from 'react-redux'
import {Dispatch} from "redux";
import {toggleMenu, hideMenu} from "../redux/actions/menuActions";
import {RootState} from "../redux/reducers/rootReducer";
import {IRoom, IRoomsResponse} from "../lib/types";
import {useRouter} from "next/router";
import HomeBtn from "../components/homeBtn";

const TTDrawer = ({Rooms}: IRoomsResponse) => {
    const router = useRouter();
    const drawerState = useSelector((state: RootState) => state.menu);
    const dispatch: Dispatch = useDispatch();
    const toggleDrawer = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        dispatch(toggleMenu());
    }
    const clickAwayHandler = (e: React.MouseEvent<Document>) => {
        e.preventDefault();
        dispatch(hideMenu());
    }
    const navigate = async (e: React.MouseEvent<HTMLButtonElement>, id: string) => {
        e.preventDefault();
        await router.push("/room/" + id);
        dispatch(toggleMenu());
    };
    return (
        <Drawer open={drawerState.value} onEscapeKeyDown={clickAwayHandler} onBackdropClick={clickAwayHandler}>
            <button onClick={toggleDrawer}> Toggle Menu</button>
            <List>
                <HomeBtn/>
                {Rooms.map((room: IRoom) => (
                    <ListItem key={room.Data.Time.toString()}>
                       <button onClick={e => navigate(e, room.Resource.ID)}>{room.Data.Temperature}</button>
                    </ListItem>
                    )
                )}
            </List>
        </Drawer>
    )
}
export default TTDrawer;
