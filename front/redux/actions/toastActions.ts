import {createAction} from "@reduxjs/toolkit"
import { HIDE_TOAST, SHOW_TOAST, TOGGLE_TOAST } from "redux/types/toastTypes";

export const showToast = createAction<string | undefined>(SHOW_TOAST);
export const toggleToast = createAction(TOGGLE_TOAST);
export const hideToast = createAction(HIDE_TOAST);
