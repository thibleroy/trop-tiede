import {menuReducer} from './menuReducer';
import {combineReducers} from '@reduxjs/toolkit';
import {roomsApi} from "../middlewares/api/rooms";
import {devicesApi} from "../middlewares/api/devices";
import { toastReducer } from './toastReducer';

const rootReducer = combineReducers({
    menu: menuReducer,
    toast: toastReducer,
    [roomsApi.reducerPath]: roomsApi.reducer,
    [devicesApi.reducerPath]: devicesApi.reducer
});

export type RootState = ReturnType<typeof rootReducer>

export default rootReducer;
