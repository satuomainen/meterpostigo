import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import CssBaseline from '@material-ui/core/CssBaseline';

import './App.css';

import Dashboard from './dashboard/Dashboard';
import DataSeries from './series/DataSeries';

function App() {
  return (
    <React.Fragment>
      <CssBaseline/>
      <BrowserRouter>
        <Switch>
          <Route path="/dataseries/:id">
            <DataSeries/>
          </Route>
          <Route path="/">
            <Dashboard/>
          </Route>
        </Switch>
      </BrowserRouter>
    </React.Fragment>
  );
}

export default App;
