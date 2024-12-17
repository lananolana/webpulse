import { useDispatch, useSelector } from 'react-redux'
export type RootState = ReturnType<typeof toolkitStore.getState>
export type AppDispatch = typeof toolkitStore.dispatch
import { IMainState, mainReducer } from '../services/mainSlice'
import { configureStore } from '@reduxjs/toolkit'
import type { TypedUseSelectorHook } from 'react-redux'

export const toolkitStore = configureStore<StateSchema>({
  reducer: {
      data: mainReducer
  },
  //@ts-ignore
  middleware: (getDefaultMiddleware) =>
  getDefaultMiddleware({
    serializableCheck: {
      ignoredActions: ['APP_SET_WINDOW_STACK'],
    },
  }),
})


export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector

export interface StateSchema {
  data: IMainState
}