import React from "react";
import './Recommendations.css';

class Recommendations extends React.Component {
    constructor(props) {
        super(props);
        this.state = { recommendations: [] };
    }

    componentDidMount() {
        fetch('http://localhost:8088/api/v1/recommendations')
            .then(res => res.json())
            .then(
                (result) => {
                    console.log(result);
                    this.setState({
                        recommendations: result
                    });
                },
                (error) => {
                    console.log(error)
                });
    }

    render() {
        return (
            <div className="recommendationsList">
                <ul>
                    {this.state.recommendations.map(item => (
                        <li className="listItem" key={item.id}><a href={item.url} target={"_blank"}>{item.video_title}</a></li>
                    ))}
                </ul>
            </div>
        );
    }
}

export default Recommendations;