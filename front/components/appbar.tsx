import { AppBar, Box, Grid, Toolbar, Typography } from "@material-ui/core";
import TTDrawer from "./drawer";

const TTAppBar = () => {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar className="TTToolbar">
                <Grid container justifyContent="space-between">  
                    <TTDrawer/>
                    <Typography className="TTTitle" align="left" variant="h4"> Trop Tiède</Typography>
                    </Grid>
                </Toolbar>
            </AppBar>
        </Box>
    )
}

export default TTAppBar;