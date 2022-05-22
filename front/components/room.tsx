import { IRoomProps } from "../lib/types";
import { Button, ListItem, ListItemText } from "@material-ui/core";
import { useRouter } from "next/router";

const TTRoom = ({ room }: IRoomProps) => {
    const router = useRouter();
    const navigate = async (e: React.MouseEvent<HTMLButtonElement>, route: string) => {
        e.preventDefault();
        await router.push(route);
    }

    return (
        <ListItem divider>
            <ListItemText primary={room.RoomDescription.Description.Name}
                secondary={room.RoomDescription.Description.Details} 
                />
                id: <Button onClick={e => navigate(e, '/room/' + room.Resource?.ID)}>{room.Resource?.ID}</Button>
        </ListItem>
    )
}

export default TTRoom;
