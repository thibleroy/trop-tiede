import { IErrorProps } from "@/lib/types";
import { Card, CardContent, Typography } from "@material-ui/core";

const TTError = ({ error }: IErrorProps) => {
    console.log('err', error)
    
    if ('status' in error) {
        // you can access all properties of `FetchBaseQueryError` here
        const errMsg = 'error' in error ? error.error : JSON.stringify(error.data)
        return (
            <>
                <Typography variant="h4" color="textSecondary">
                    Error {error.status}
                </Typography>
                <Typography variant="h6">
                    Message :
                    <Typography color="textSecondary">
                        {errMsg}
                    </Typography>
                </Typography>
            </>
        )
    }
    else {
        // you can access all properties of `SerializedError` here
        return (
            <Card>
                <CardContent>
                    <Typography variant="h4" color="textSecondary">
                        Error {error.message}
                    </Typography>
                </CardContent>
            </Card>
        )
    }
};

export default TTError;
