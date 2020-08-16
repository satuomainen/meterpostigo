import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

export interface Summary {
  id: number,
  createdAt: Date,
  updatedAt: Date,
  dataSeriesId: number,
  dataSeriesLabel: string,
  dataSeriesName: string,
  dataSeriesDescription: string,
  currentValue: string,
  maxValue: string,
  minValue: string
}

export interface SummaryState {
  dataseries: Summary[] | null
}

const initialState: SummaryState = {
  dataseries: null,
};

const fetchSummaries = createAsyncThunk<Summary[], void, object>(
  'summaries/fetchSummaries',
  async () => {
    const response = await axios.get('http://localhost:9000/dataseries/summaries');
    return response.data;
  });

const summarySlice = createSlice({
  name: 'summaries',
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder.addCase(fetchSummaries.fulfilled, (state, action) => {
      state.dataseries = action.payload;
    });

    builder.addCase(fetchSummaries.rejected, (state, action) => {
      state.dataseries = null;
    });
  },
});


export {
  fetchSummaries,
};

export default summarySlice.reducer;
