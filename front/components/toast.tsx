import { RootState } from '@/store';
import { Snackbar } from '@material-ui/core';
import { useSelector, useDispatch } from 'react-redux';
import { hideToast } from 'redux/actions/toastActions';

const TTToast = () => {
    const toastState = useSelector((state: RootState) => state.toast);
    const dispatch = useDispatch();
    const handleClose = () => {
        dispatch(hideToast());
    }

    return (
        <Snackbar
            autoHideDuration={5000}
            anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
            open={toastState.open}
            onClose={handleClose}
            message={toastState.message}
        />
    )
}

export default TTToast;
