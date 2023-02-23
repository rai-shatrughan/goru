import { Menus } from './CompList';

export default function Header(props) {

    return (
        <div className="pure-menu pure-menu-horizontal">
            <ul className="pure-menu-list">
                {
                    Menus.map((menu) => (
                        <li className="pure-menu-item" key={menu}>
                            <a className="pure-menu-link"
                                href=""
                                key={menu}
                                onClick={props.onClicked}
                            >{menu}</a>
                        </li>
                    ))
                }
            </ul>
        </div>
    );
}