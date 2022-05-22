import { ListItem, Button, Typography } from "@material-ui/core";
import React from "react";
import { Dispatch } from "redux";
import { hideMenu } from "../redux/actions/menuActions";
import { useRouter } from "next/router";
import { IDrawerItemProps } from "@/lib/types";
import { useDispatch } from "react-redux";

const TTDrawerItem = (drawerItemParam: IDrawerItemProps) => {
    const router = useRouter();
    const dispatch: Dispatch = useDispatch();

    const handleClose = () => {
        dispatch(hideMenu());
    };

    const navigate = async (e: React.MouseEvent<HTMLButtonElement>, route: string) => {
        e.preventDefault();
        await router.push(route);
        handleClose();
    }

    return (
        <ListItem>
            <Button onClick={e => navigate(e, drawerItemParam.route)}>
                <Typography variant="h6">
                {drawerItemParam.label}
                </Typography>
            </Button>
        </ListItem>
    )
}
export default TTDrawerItem;
