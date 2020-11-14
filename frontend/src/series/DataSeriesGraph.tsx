import React from 'react';
import moment from 'moment';
import { CartesianGrid, Label, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';

import { Reading, Series } from './seriesSlice';

const TIMESTAMP_FORMAT = 'YYYY-MM-DD HH:mm:ss';

interface Datapoint {
  x: string,
  y: number,
  unit: string
}

function convertReading(reading: Reading, unit: string): Datapoint {
  return {
    x: moment(reading.createdAt).format(TIMESTAMP_FORMAT),
    y: parseFloat(reading.value),
    unit,
  };
}

function createChartData(readings: Reading[] | null, unit: string) {
  return readings?.map(r => convertReading(r, unit));
}

export default function DataSeriesGraph(props: Series) {
  const { summary, readings } = props;

  if (!readings || !summary) {
    return null;
  }

  const data = createChartData(readings, summary?.dataSeriesLabel);

  return (
    <ResponsiveContainer>
      <LineChart
        data={data}
        width={500}
        height={350}
        margin={{ left: 20, bottom: 20 }}
      >
        <CartesianGrid strokeDasharray="3 3"/>
        <XAxis dataKey="x"
               type="category"
        >
          <Label position="bottom"
                 content={() => 'Reading time'}
          />
        </XAxis>
        <YAxis type="number"
               domain={[ 'dataMin', 'dataMax' ]}
        >
          <Label angle={-90}
                 content={() => `${summary.dataSeriesName} (${summary.dataSeriesLabel})`}
                 position="left"
                 dy={-75}
                 dx={-5}
          />
        </YAxis>
        <Tooltip formatter={val => [ val, `${summary.dataSeriesName} (${summary.dataSeriesLabel})` ]}/>
        <Line type="monotone" dataKey="y" stroke="#8884d8"/>
      </LineChart>
    </ResponsiveContainer>
  );
}
