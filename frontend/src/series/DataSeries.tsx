import React, { useEffect, useMemo, useState } from 'react';
import { useParams } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { Link } from 'react-router-dom';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';

import './DataSeries.scss';

import { Summary } from '../dashboard/summarySlice';
import { RootState } from '../store';
import { fetchSeries, Series, SeriesState } from './seriesSlice';
import DataSeriesGraph from './DataSeriesGraph';

interface DataSeriesProps {
  summary: Summary
}

function getSeriesName(series: Series | null) {
  if (!series?.summary) {
    return 'Loading...';
  }
  return series.summary.dataSeriesName;
}

export default function DataSeries() {
  const { id } = useParams();
  const dispatch = useDispatch();
  const [ selectedTab, setSelectedTab ] = useState(0);
  const { series } = useSelector<RootState, SeriesState>(state => state.series);

  function onTabChanged(event: object, newTabIndex: number) {
    setSelectedTab(newTabIndex);
  }

  function downloadWanted() {
    // TODO(satuomainen): download from backend
    alert('Download not implemented yet');
  }

  useEffect(() => {
    if (!series) {
      dispatch(fetchSeries(id));
    }
  }, [ id, series, dispatch ]);

  const seriesName = useMemo(() => {
    return getSeriesName(series);
  }, [ series ]);

  const readingGraph = useMemo(() => {
    if (!series) {
      return null;
    }

    let readings = null;
    switch (selectedTab) {
      case 0: {
        readings = series.readings;
      }
    }
    return <DataSeriesGraph readings={readings} summary={series.summary}/>;
  }, [ series, selectedTab ]);

  return (
    <div className="dataseries">
      <div className="dataseries--header">
        <Typography variant="h1">{seriesName}</Typography>
      </div>
      <Container maxWidth="xl">
        <Grid container direction="column" spacing={2} alignItems="stretch">
          <Grid item>
            <Link to="/">Dashboard</Link> › {seriesName}
          </Grid>
          <Grid container justify="space-between">
            <Tabs
              value={selectedTab}
              onChange={onTabChanged}
              indicatorColor="primary"
              textColor="primary"
              aria-label="Select between time series and daily averages"
            >
              <Tab value={0} label="Time series"/>
              <Tab value={1} label="Daily averages"/>
            </Tabs>
            <Button color="primary" variant="contained" onClick={downloadWanted}>Download as CSV</Button>
          </Grid>
          <Grid item className="dataseries--graph-container">
            {readingGraph}
          </Grid>
        </Grid>
      </Container>
    </div>
  );
}
