import {
  createAsyncThunk,
  createSlice,
  PayloadAction,
  SerializedError,
} from '@reduxjs/toolkit'
import { api } from './api';

export interface IMainState {
  url: string | null,
  data: any | null,
  isLoading: boolean,
  isError: boolean,
  errorMessage: string | null,
}

const initialState: IMainState = {
  url: null,
  data: null,
  isLoading: false,
  isError: false,
  errorMessage: null,
}

// Получение ответа
const getInfo = createAsyncThunk(
  'main/getInfo',
  async (data: string, { rejectWithValue }) => {
    try {
      return await api.getInfo(data)
    } catch (error) {
      return rejectWithValue(error)
    }
  }
)

export const mainSlice = createSlice({
  name: 'main',
  initialState,
  reducers: {},
  extraReducers: builder =>
    builder
      .addCase(getInfo.pending, state => {
        state.isLoading = true
      })
      .addCase(
        getInfo.fulfilled,
        (state, action: PayloadAction<any | null>) => {
            state.data = action.payload
            state.isLoading = false
            state.isError= false
            state.errorMessage = ''
        }
      )
      .addCase(getInfo.rejected, (state, action) => {
        state.isLoading = false
        state.isError = true
        state.errorMessage =
            (action.payload as SerializedError)?.message ||
            'Произошла ошибка при получении ответа'
      })
})

export const { reducer: mainReducer } = mainSlice

export { getInfo }


