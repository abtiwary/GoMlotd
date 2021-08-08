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
    fetch('http://localhost:8088/api/v1/recommendations')
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
