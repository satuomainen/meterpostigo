import { configureStore } from '@reduxjs/toolkit';

import summarySlice, { Summary } from './dashboard/summarySlice';
import seriesSlice, { Series } from './series/seriesSlice';


export interface RootState {
  summary: {
    dataseries: Summary[] | null
  },
  series: {
    series: Series | null
  }
}

const store = configureStore({
  reducer: {
    summary: summarySlice,
    series: seriesSlice,
  },
});

export default store;
