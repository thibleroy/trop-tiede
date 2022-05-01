import { Action, Middleware } from "@reduxjs/toolkit";
import { RootState } from "@/store";

const loggerMiddleware: Middleware<
{}, // Most middleware do not modify the dispatch return value
RootState
> = store => (next: any) => (action: Action) => {
    console.log('dispatching', action)
    let result = next(action)
    console.log('next state', store.getState())
    return result
  };

  
  export default loggerMiddleware;