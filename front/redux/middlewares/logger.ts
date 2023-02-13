import { Action, Middleware } from "@reduxjs/toolkit";
import { RootState, store } from "@/store";
import { showToast } from "redux/actions/toastActions";

const loggerMiddleware: Middleware<
  {}, // Most middleware do not modify the dispatch return value
  RootState
> = store => (next: any) => (action: Action) => {
  console.log('dispatching', action)
  if ('error' in action) { }
  switch (action.type) {
    case 'devicesApi/executeQuery/rejected':
      store.dispatch(showToast("Fetch devices error"));
      break;
    case 'roomsApi/executeQuery/rejected':
      store.dispatch(showToast("Fetch rooms error"));
      break;
    default: break;
  }
  const result = next(action)
  console.log('next state', store.getState())
  return result
};


export default loggerMiddleware;