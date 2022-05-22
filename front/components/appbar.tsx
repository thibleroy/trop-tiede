import { AppBar, Toolbar, Typography } from "@material-ui/core";
import TTDrawer from "./drawer";
const TTAppBar = () => {
    return (<AppBar position="static">
        <Toolbar>
            <TTDrawer />
            <Typography align="center" variant="h4"> Trop Tiède</Typography>
        </Toolbar>
    </AppBar>
    )
}

export default TTAppBar;