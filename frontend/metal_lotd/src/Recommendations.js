import React from "react";
import './Recommendations.css';

const Recommendations = (props) => {
    let links = props.links;

    return (
            <div className="recommendationsList">
                {links.map(item => (
                        <div className="metalItem" key={item.id}><a href={item.url} target={"_blank"}>{item.video_title}</a></div>
                ))}
            </div>
    )
};

export default Recommendations;
