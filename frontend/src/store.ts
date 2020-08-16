import { configureStore } from '@reduxjs/toolkit';

import summarySlice, { Summary } from './dashboard/summarySlice';

export interface RootState {
  summary: {
    dataseries: Summary[] | null
  }
}

const store = configureStore({
  reducer: {
    summary: summarySlice,
  },
});

export default store;
