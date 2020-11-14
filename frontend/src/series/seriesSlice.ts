import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';
import { Summary } from '../dashboard/summarySlice';

export interface Reading {
  id: number,
  createdAt: string,
  updatedAt: string,
  value: string
}

export interface Series {
  summary: Summary | null,
  readings: Reading[] | null
}

export interface SeriesState {
  series: Series | null
}

const initialState: SeriesState = {
  series: null,
};

const READING_LIMIT = 100;

function compareReadingsByTime(left: Reading, right: Reading) {
  return left.createdAt.localeCompare(right.createdAt);
}

const fetchSeries = createAsyncThunk<Series, number, object>(
  'summaries/fetchSeries',
  async (dataSeriesId: number) => {
    const summaryUrl = `http://localhost:9000/dataseries/${dataSeriesId}/summary`;
    const summaryResponse = await axios.get(summaryUrl);

    const readingsUrl = `http://localhost:9000/dataseries/${dataSeriesId}/readings?limit=${READING_LIMIT}`;
    const readingsResponse = await axios.get(readingsUrl);
    const readings = readingsResponse.data;
    readings.sort(compareReadingsByTime);

    return {
      summary: summaryResponse.data,
      readings
    };
  });

const seriesSlice = createSlice({
  name: 'series',
  initialState,
  reducers: {},
  extraReducers: builder => {
    builder.addCase(fetchSeries.fulfilled, (state, action) => {
      state.series = action.payload;
    });

    builder.addCase(fetchSeries.rejected, (state, action) => {
      state.series = null;
    });
  },
});


export {
  fetchSeries,
};

export default seriesSlice.reducer;
