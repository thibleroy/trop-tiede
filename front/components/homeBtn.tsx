import React from 'react';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import {useRouter} from "next/router";
import {useDispatch} from "react-redux";
import {toggleMenu} from "../redux/actions/menuActions";

const HomeBtn = () => {
    const dispatch = useDispatch();
    const router = useRouter();
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const navigate = async (e: React.MouseEvent<HTMLLIElement>, route: string) => {
        e.preventDefault();
        await router.push(route);
        handleClose();
        dispatch(toggleMenu());
    }

    return (
        <div>
            <Button aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
                Open Menu
            </Button>
            <Menu
                id="simple-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
            >
                <MenuItem onClick={e => navigate(e, "/")}>Index</MenuItem>
                <MenuItem onClick={e => navigate(e, "/home")}>Home</MenuItem>
                <MenuItem onClick={e => navigate(e, "/devices")}>Devices</MenuItem>
            </Menu>
        </div>
    );
}

export default HomeBtn
