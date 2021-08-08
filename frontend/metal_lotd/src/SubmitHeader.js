import React from "react";
import './SubmitHeader.css';
import mlotdLogo from './mlotd.png';

class SubmitHeader extends React.Component {
    constructor(props) {
        super(props);
        this.state = { video: '' };
    }

    handleChange = (event) => {
        this.setState({[event.target.name]: event.target.value});
    }

    handleSubmit = (event) => {
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

    render() {
        return (
            <div className="sub-wrapper">
                <div className="header_container">
                    <div className="sitelogo"><img src={mlotdLogo} width="125px" height="125px"></img></div>
                    <div className="centerMe">
                    <form onSubmit={this.handleSubmit}>
                    <label>
                    Submit a link: &nbsp;
                    <input className="submitText" type="text" value={this.state.value} name="video" onChange={this.handleChange} />
                    </label> &nbsp;
                    <input className="submit_button" type="submit" value="Submit" />
                    </form>
                    </div>
                </div>
            </div>
        );
    }
}

export default SubmitHeader;
