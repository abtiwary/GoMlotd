import React from 'react';
import './App.css';

import SubmitHeader from "./SubmitHeader";
import Recommendations from "./Recommendations";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { };
  }

  render() {
    return (
      <div>
        <SubmitHeader />
        <Recommendations />
      </div>
    );
  }
}

export default App;
