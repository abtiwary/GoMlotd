import {useState, useEffect} from 'react';
import React from 'react';
import './App.css';

import SubmitHeader from "./SubmitHeader";
import Recommendations from "./Recommendations";

const App = () => {
  const [metalLinks, setMetalLinks] = useState(null);

  const forceUpdate = () => {
    setMetalLinks([]);
    window.location.reload();
    return;
  };

  useEffect(() => {
    console.log("useeffect triggered");
    fetch('http://14.203.254.111:8666/api/v1/recommendations', {
      crossDomain: true,
      method: 'GET',
      mode: 'cors',
      headers: {},
    })
      .then(res => res.json())
      .then(
        (result) => {
          console.log(result);
          setMetalLinks(result);
        },
        (error) => {
          console.log(error)
        });
  }, []);

  return (
    <div>
      <SubmitHeader refresher={forceUpdate} />
      { metalLinks && <Recommendations links={metalLinks} />}
    </div>
  );

};

export default App;
