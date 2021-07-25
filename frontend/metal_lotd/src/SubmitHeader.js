import React from "react";
import './SubmitHeader.css';

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

export default SubmitHeader;
