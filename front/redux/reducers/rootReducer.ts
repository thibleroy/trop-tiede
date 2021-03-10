import {menuReducer} from './menuReducer';
import {combineReducers} from 'redux';

const rootReducer = combineReducers({
    menu: menuReducer
});

export type RootState = ReturnType<typeof rootReducer>

export default rootReducer;
