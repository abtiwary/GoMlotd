import React from 'react';
import './App.css';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = { video: '' };
  }

  handleChange = (event) => {
    this.setState({[event.target.name]: event.target.value});
  }

  handleSubmit = (event) => {
    //alert('A form was submitted: ' + this.state.video);
    
    fetch('http://localhost:8088/api/v1/recommendation', {
        method: 'POST',
        // We convert the React state to JSON and send it as the POST body
        body: JSON.stringify(this.state)
      }).then(function(response) {
        console.log(response)
        return response.json();
      });

    event.preventDefault();
  }

  componentDidMount() {
    alert("fetch here")
  }

  render() {
    return (
      <div className="wrapper">
      <form onSubmit={this.handleSubmit}>
      <label>
          Video:
          <input type="text" value={this.state.value} name="video" onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    </div>
    );
  }
}

export default App;
