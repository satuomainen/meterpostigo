import React, { useEffect, useMemo, useState } from 'react';
import { useParams } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import moment from 'moment';
import Papa from 'papaparse';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { Link } from 'react-router-dom';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';

import './DataSeries.scss';

import { RootState } from '../store';
import { fetchSeries, Reading, ReadingAverage, Series, SeriesState } from './seriesSlice';
import DataSeriesGraph, { Datapoint } from './DataSeriesGraph';

interface DownloadableRecord {
  "Reading time": string,
  "Value": string
}

const TIMESTAMP_FORMAT = 'YYYY-MM-DD HH:mm:ss';
const DATE_FORMAT = 'YYYY-MM-DD';

function downloadContent(content: string, seriesName: string) {
  const element = document.createElement("a");
  const file = new Blob([content], {type: 'text/csv'});

  element.href = URL.createObjectURL(file);
  element.download = `${seriesName}.csv`;
  document.body.appendChild(element);

  element.click();
}

function mapReadingToDownloadableRecord(reading: Reading): DownloadableRecord {
  return {
    "Reading time": reading.createdAt,
    "Value": reading.value
  };
}

function mapReadingAverageToDownloadableRecord(average: ReadingAverage): DownloadableRecord {
  return {
    "Reading time": average.date,
    "Value": average.value
  };
}

function getSeriesName(series: Series | null) {
  if (!series?.summary) {
    return 'Loading...';
  }
  return series.summary.dataSeriesName;
}

function convertReading(reading: Reading, unit: string): Datapoint {
  return {
    x: moment(reading.createdAt).format(TIMESTAMP_FORMAT),
    y: Number(parseFloat(reading.value).toFixed(1)),
    unit,
  };
}

function mapReadingsToDatapoints(readings: Reading[] | null, unit: string): Datapoint[] {
  if (!readings) {
    return []
  }

  return readings?.map(r => convertReading(r, unit));
}

function convertAverage(average: ReadingAverage, unit: string): Datapoint {
  return {
    x: moment(average.date).format(DATE_FORMAT),
    y: Number(parseFloat(average.value).toFixed(1)),
    unit
  };
}

function mapAveragesToDatapoints(averages: ReadingAverage[], unit: string) {
  if (!averages) {
    return [];
  }

  return averages?.map(a => convertAverage(a, unit));
}

export default function DataSeries() {
  const { id } = useParams();
  const dispatch = useDispatch();
  const [ selectedTab, setSelectedTab ] = useState(0);
  const [ isDownloadAvailable, setIsDownloadAvailable ] = useState(false);
  const { series } = useSelector<RootState, SeriesState>(state => state.series);

  function onTabChanged(event: object, newTabIndex: number) {
    setSelectedTab(newTabIndex);
  }

  function downloadWanted() {
    if (selectedTab === 0) {
      if (series?.readings) {
        const csv = Papa.unparse(series.readings.map(mapReadingToDownloadableRecord));
        downloadContent(csv, getSeriesName(series));
      }
    } else if (selectedTab === 1) {
      if (series?.averages) {
        const csv = Papa.unparse(series.averages.map(mapReadingAverageToDownloadableRecord))
        downloadContent(csv, getSeriesName(series));
      }
    }
  }

  useEffect(() => {
    if (!series) {
      dispatch(fetchSeries(id));
    } else {
      setIsDownloadAvailable(true);
    }
  }, [ id, series, dispatch ]);

  const seriesName = useMemo(() => {
    return getSeriesName(series);
  }, [ series ]);

  const readingGraph = useMemo(() => {
    if (!series?.summary) {
      return null;
    }

    const unit = series?.summary?.dataSeriesLabel || '';
    let readings: Datapoint[] = [];
    if (selectedTab === 0) {
        if (!series?.readings) {
          return null;
        }
        readings = mapReadingsToDatapoints(series.readings, unit);
    } else if (selectedTab === 1) {
        if (!series?.averages) {
          return null;
        }
        readings = mapAveragesToDatapoints(series.averages, unit);
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
            <Button
              color="primary"
              variant="contained"
              onClick={downloadWanted}
              disabled={!isDownloadAvailable}
            >
              Download as CSV
            </Button>
          </Grid>
          <Grid item className="dataseries--graph-container">
            {readingGraph}
          </Grid>
        </Grid>
      </Container>
    </div>
  );
}
