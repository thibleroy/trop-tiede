import {createReducer} from "@reduxjs/toolkit";
import { hideMenu, showMenu, toggleMenu } from "redux/actions/menuActions";

interface MenuState {
    open: boolean
  }

const initialMenuState: MenuState = {open: false};

export const menuReducer = createReducer(initialMenuState, (builder) => {
    builder
      .addCase(showMenu, (state, action) => {
        state.open = true;
      })
      .addCase(hideMenu, (state, action) => {
        state.open = false;
      })
      .addCase(toggleMenu, (state, action) => {
        state.open = !state.open;
      })
  });
