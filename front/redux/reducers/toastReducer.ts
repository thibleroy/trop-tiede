import {createReducer} from "@reduxjs/toolkit";
import { hideToast, showToast, toggleToast } from "redux/actions/toastActions";

interface ToastState {
    open: boolean;
    message?: string;
  }

const ToastState: ToastState = {open: false}

export const toastReducer = createReducer(ToastState, (builder) => {
    builder
      .addCase(showToast, (state, action) => {
        state.open = true;
        state.message = action.payload
      })
      .addCase(hideToast, (state, action) => {
        state.open = false;
      })
      .addCase(toggleToast, (state, action) => {
        state.open = !state.open;
      })
  });
