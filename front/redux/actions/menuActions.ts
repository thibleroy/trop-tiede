import {createAction} from "@reduxjs/toolkit"
import { SHOW_MENU, HIDE_MENU, TOGGLE_MENU } from "redux/types/menuTypes";

export const showMenu = createAction(SHOW_MENU);
export const toggleMenu = createAction(TOGGLE_MENU);
export const hideMenu = createAction(HIDE_MENU);

