import loggerMiddleware from "./logger";
import {devicesApi} from "./api/devices";
import {roomsApi} from "./api/rooms";
import { CurriedGetDefaultMiddleware } from "@reduxjs/toolkit/dist/getDefaultMiddleware";
import { MiddlewareArray } from "@reduxjs/toolkit";
import { RootState } from "@/store";

const middlewares = (getDefaultMiddleware: CurriedGetDefaultMiddleware): MiddlewareArray<RootState> => {
    return getDefaultMiddleware().concat(loggerMiddleware, roomsApi.middleware, devicesApi.middleware)
}

export default middlewares;
