import React from 'react';
import { CartesianGrid, Label, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';

import { Summary } from '../dashboard/summarySlice';

interface SeriesData {
  summary: Summary,
  readings: Datapoint[]
}

export interface Datapoint {
  x: string,
  y: number,
  unit: string
}

export default function DataSeriesGraph(props: SeriesData) {
  const { summary, readings } = props;

  if (!readings || !summary) {
    return null;
  }

  return (
    <ResponsiveContainer>
      <LineChart
        data={readings}
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
