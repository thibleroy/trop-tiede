import {Drawer, ListItem, List, ClickAwayListener, Modal} from "@material-ui/core";
import React from "react";
import {useSelector, useDispatch} from 'react-redux'
import {Dispatch} from "redux";
import {toggleMenu, hideMenu} from "../redux/actions/menuActions";
import {RootState} from "../redux/reducers/rootReducer";
import {IRoom, IRoomsResponse} from "../lib/types";
import {useRouter} from "next/router";
import HomeBtn from "../components/homeBtn";
import { Button } from '@material-ui/core';
import {Add} from "@material-ui/icons";
import DrawerToggler from "./drawerToggler";

function rand() {
    return Math.round(Math.random() * 20) - 10;
}

function getModalStyle() {
    const top = 50 + rand();
    const left = 50 + rand();

    return {
        top: `${top}%`,
        left: `${left}%`,
        transform: `translate(-${top}%, -${left}%)`,
    };
}

const TTDrawer = ({Rooms}: IRoomsResponse) => {
    const router = useRouter();
    const drawerState = useSelector((state: RootState) => state.menu);
    const dispatch: Dispatch = useDispatch();
    const clickAwayHandler = (e: React.MouseEvent<Document>) => {
        e.preventDefault();
        dispatch(hideMenu());
    }
    const navigate = async (e: React.MouseEvent<HTMLButtonElement>, id: string) => {
        e.preventDefault();
        await router.push("/room/" + id);
        dispatch(toggleMenu());
    };
    const [modalStyle] = React.useState(getModalStyle);
    const [open, setOpen] = React.useState(false);

    const handleOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const body = (
        <div style={modalStyle}>
            <h2 id="simple-modal-title">Text in a modal</h2>
            <p id="simple-modal-description">
                Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
            </p>
        </div>
    );
    return (
        <Drawer open={drawerState.value} onEscapeKeyDown={clickAwayHandler} onBackdropClick={clickAwayHandler}>
            <DrawerToggler/>
            <HomeBtn/>
            <Button color="primary" size="large" onClick={handleOpen}><Add/></Button>
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
