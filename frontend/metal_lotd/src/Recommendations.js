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
                {this.state.recommendations.map(item => (
                        <div className="metalItem" key={item.id}><a href={item.url} target={"_blank"}>{item.video_title}</a></div>
                ))}
            </div>
        );
    }
}

export default Recommendations;