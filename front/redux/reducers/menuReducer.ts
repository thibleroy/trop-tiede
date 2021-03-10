import {Action} from "redux";
import {TOGGLE_MENU, HIDE_MENU} from "../types/menuTypes";

const initialMenuState = {value: false};

export const menuReducer = (state = initialMenuState, action: Action) => {
    switch (action.type) {
        case TOGGLE_MENU:
            return {...state, value: !state.value};
        case HIDE_MENU:
            return {...state, value: false}
        default:
            return {...state};
    }
};
