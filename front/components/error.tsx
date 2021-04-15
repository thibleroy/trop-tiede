import {IError} from "../lib/types";
import {Card, CardContent, Typography} from "@material-ui/core";

const TTError = ({Code, Message}: IError) => {
    return (
        <Card>
            <CardContent>
                <Typography variant="h4" color="textSecondary">
                    Error {Code}
                </Typography>
                <Typography variant="h6">
                    Message :
                    <Typography color="textSecondary">
                        {Message}
                    </Typography>
                </Typography>
            </CardContent>
        </Card>
    )
};

export default TTError;
