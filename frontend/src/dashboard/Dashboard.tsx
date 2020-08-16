import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';

import './Dashboard.scss';

import { fetchSummaries, SummaryState } from './summarySlice';
import { RootState } from '../store';
import SummaryGauge from '../gauge/SummaryGauge';

export default function Dashboard() {
  const dispatch = useDispatch();
  const { dataseries } = useSelector<RootState, SummaryState>(state => state.summary);

  useEffect(() => {
    if (!dataseries || !Array.isArray(dataseries)) {
      dispatch(fetchSummaries());
    }
  }, [ dataseries, dispatch ]);

  function createSummaryItems() {
    if (!dataseries) {
      return undefined;
    }

    return dataseries.map(ds => (<SummaryGauge key={ds.id} dataSeriesSummary={ds}/>));
  }

  return (
    <div className="dashboard">
      <div className="dashboard--header">
        <Typography variant="h1">Project title goes here</Typography>
      </div>
      <Container maxWidth="xl">

        <Grid container spacing={2} alignItems="stretch">
          {createSummaryItems()}
        </Grid>
      </Container>
    </div>
  );
}
