import {useState} from "react";
import './SubmitHeader.css';
import mlotdLogo from './mlotd.png';

const SubmitHeader = (props) => {
    const [submitState, setSubmitState] = useState({video: ''});

    let actionFn = props.action;
    let stateUpdater = props.stateupdater;

    const handleSubmit = (event) => {
        fetch('http://localhost:8088/api/v1/recommendation', {
            method: 'POST',
            // We convert the React state to JSON and send it as the POST body
            body: JSON.stringify(submitState)
        }).then(function(response) {
            console.log(response)
            actionFn(stateUpdater);
            return response.text();
        });

        event.preventDefault();
    }

    const handleChange = (event) => {
        setSubmitState({[event.target.name]: event.target.value});
    }

    return (
        <div className="sub-wrapper">
            <div className="header_container">
                <div className="sitelogo"><img src={mlotdLogo} width="125px" height="125px"></img></div>
                <div className="centerMe">
                <form onSubmit={handleSubmit}>
                <label>
                Submit a link: &nbsp;
                <input className="submitText" type="text" name="video" onChange={handleChange} />
                </label> &nbsp;
                <input className="submit_button" type="submit" value="Submit" />
                </form>
                </div>
            </div>
        </div>
    );
};

export default SubmitHeader;
