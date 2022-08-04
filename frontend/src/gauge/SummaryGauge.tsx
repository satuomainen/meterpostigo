import React from 'react';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardMedia from '@material-ui/core/CardMedia';
import Gauge from 'react-svg-gauge';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

import { Summary } from '../dashboard/summarySlice';
import Button from '@material-ui/core/Button';

const gaugeStyles = {
  gaugeTitle: {
    fontSize: '1.15rem',
  },
  gaugeValue: {
    fontSize: '1.3rem',
  },
  minMaxValues: {
    fontSize: '0.75rem',
  },
};

function toGaugeValue(s: string): number | undefined {
  const numericValue = parseFloat(s);

  if (isNaN(numericValue)) {
    return undefined;
  }

  return Number(numericValue.toFixed(1));
}

function createValueFormatter(label: string) {
  if (label && label.length > 0) {
    return (value: number) => `${value} ${label}`;
  }

  return (value: number) => `${value}`;
}

interface ValueGaugeProps {
  dataSeriesSummary: Summary
}

export default function SummaryGauge({ dataSeriesSummary }: ValueGaugeProps) {

  return (
    <div>
      <Button href={`/dataseries/${dataSeriesSummary.dataSeriesId}`} className="summary-item">
        <Card className="summary-item--container">
          <CardActionArea className="summary-item--actions">
            <CardMedia className="gauge">
              <Gauge
                label={dataSeriesSummary.dataSeriesName}
                topLabelStyle={gaugeStyles.gaugeTitle}
                value={toGaugeValue(dataSeriesSummary.currentValue)}
                valueLabelStyle={gaugeStyles.gaugeValue}
                valueFormatter={createValueFormatter(dataSeriesSummary.dataSeriesLabel)}
                min={toGaugeValue(dataSeriesSummary.minValue)}
                max={toGaugeValue(dataSeriesSummary.maxValue)}
                minMaxLabelStyle={gaugeStyles.minMaxValues}
                width={250}
                height={250}
              />
            </CardMedia>
            <CardContent>
              <Typography>{dataSeriesSummary.dataSeriesName} ({dataSeriesSummary.dataSeriesLabel})</Typography>
              {dataSeriesSummary.dataSeriesDescription}
            </CardContent>
          </CardActionArea>
        </Card>
      </Button>
    </div>
  );
}
