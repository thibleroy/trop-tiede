import {Drawer, ListItem, List, Modal} from "@material-ui/core";
import React from "react";
import {useSelector, useDispatch} from 'react-redux'
import {Dispatch} from "redux";
import {toggleMenu, hideMenu} from "../redux/actions/menuActions";
import {RootState} from "../redux/reducers/rootReducer";
import {IRoom, IRoomsResponse} from "../lib/types";
import {useRouter} from "next/router";
import HomeBtn from "../components/homeBtn";
import {Button} from '@material-ui/core';
import {Add} from "@material-ui/icons";
import DrawerToggler from "./drawerToggler";

const TTDrawer = ({Rooms}: IRoomsResponse) => {
    const router = useRouter();
    const drawerState = useSelector((state: RootState) => state.menu);
    const dispatch: Dispatch = useDispatch();
    const clickAwayHandler = (e: React.MouseEvent<Document>) => {
        e.preventDefault();
        dispatch(hideMenu());
    }
    const navigate = async (e: React.MouseEvent<HTMLButtonElement>, id?: string) => {
        e.preventDefault();
        await router.push("/room/" + id);
        dispatch(toggleMenu());
    };
    const [open, setOpen] = React.useState(false);

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const addRoom = async (e: React.MouseEvent<HTMLFormElement>) => {
        e.preventDefault();
        const room: IRoom = {
            RoomDescription: {
                Description: {Details: "test details", Name: "test name"}
            }
        };
        const res = await fetch(
            process.env.NEXT_PUBLIC_API + '/rooms',
            {
                body: JSON.stringify(room),
                headers: {
                    'Content-Type': 'application/json'
                },
                method: 'POST'
            }
        );
        console.log('post create', res.headers.entries());
        handleClose();
    };

    const body = (
        <form onSubmit={addRoom}>
            <label htmlFor="name">Name</label>
            <input id="name" type="text" autoComplete="name" required/>
            <button type="submit">Register</button>
        </form>
    );
    return (
        <Drawer open={drawerState.value} onEscapeKeyDown={clickAwayHandler} onBackdropClick={clickAwayHandler}>
            <DrawerToggler/>
            <HomeBtn/>
            <Button color="default" variant="contained" size="large" onClick={handleOpen}><Add/></Button>
            <Modal
                open={open}
                onClose={handleClose}
                aria-labelledby="simple-modal-title"
                aria-describedby="simple-modal-description"
            >
                {body}
            </Modal>
            <List>
                {Rooms.map((room: IRoom) => (
                        <ListItem key={room.Resource?.ID}>
                            <Button
                                onClick={e => navigate(e, room.Resource?.ID)}>{room.RoomDescription.Description.Name}</Button>
                        </ListItem>
                    )
                )}
            </List>
        </Drawer>
    )
}
export default TTDrawer;
