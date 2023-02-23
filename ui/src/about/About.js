import React, { useState } from "react";
import "./style/index.css";
import "bulma/css/bulma.css";
import "purecss/build/pure-min.css"
import "purecss/build/grids-responsive-min.css";
import {
    Header,
    CompMap
} from "./comp/CompList";

export default function About() {
    const [show, setShow] = useState("Home");
    const [theme, setTheme] = useState("dark");
    const VisibleComponent = CompMap[show];

    function headerClicked(e) {
        e.preventDefault();
        setShow(e.target.innerText);
    };

    return (
        <div>
            <Header onClicked={headerClicked} />

            <div id="div-center">
                <VisibleComponent />
            </div>
        </div>
    );
}

