import { useEffect } from "react";
import { SunIcon, MoonIcon } from "@primer/octicons-react";

export function setTheme(themeName) {
    localStorage.setItem('theme', themeName);
    document.documentElement.className = themeName;
}

export function getTheme() {
    if (Number.isNaN(localStorage.getItem('theme'))) {
        return "dark"
    } else {
        return localStorage.getItem('theme');
    }
}

export function toggleTheme() {
    if (getTheme() === 'dark') {
        setTheme('light');
    } else {
        setTheme('dark');
    }
}

export default function Theme(props) {

    useEffect(() => {
        setTheme(props.theme);
    });

    return (
        <button
            className="theme-button"
            onClick={() => {
                toggleTheme();
                props.onClicked();
            }}>

            {props.theme === 'dark' ?
                <SunIcon size={20} fill="#FFFFFF" /> :
                <MoonIcon size={20} fill="#000000" />
            }
        </button>
    );
}
