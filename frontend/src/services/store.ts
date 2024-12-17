import { compose } from 'redux';
import { configureStore } from '@reduxjs/toolkit';
import { mainReducer } from './mainSlice';

declare global {
  interface Window {
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: typeof compose;
  }
}

export const store = configureStore({
  reducer: {
    data: mainReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;