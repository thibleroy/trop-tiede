import { configureStore, EnhancedStore } from '@reduxjs/toolkit';
import rootReducer from './reducers/rootReducer';
import middlewares from './middlewares';

export const store: EnhancedStore = configureStore(
{
    reducer: rootReducer,
    middleware: middlewares,
    devTools: true
}
);

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
