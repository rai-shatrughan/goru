import React from 'react';
import Simple from './simple';

export default function Header({ onClicked }) {

    return (
        <div className="pure-menu pure-menu-horizontal">
            <ul className="pure-menu-list">
                <Simple
                    name="Home"
                    onShow={onClicked} />
                <Simple
                    name="Skills"
                    onShow={onClicked} />
                <Simple
                    name="TechStack"
                    onShow={onClicked} />
                <Simple
                    name="Domains"
                    onShow={onClicked} />
            </ul>
        </div>

    );
}






