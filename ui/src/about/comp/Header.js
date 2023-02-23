import { Menus } from './CompList';

export default function Header(props) {

    return (
        <nav class="navbar" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
                <a class="navbar-item" href="https://rai-shatrughan.github.io">
                    <img src="sr.png" width="20" height="20" />
                </a>

                <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                </a>
            </div>

            <div id="navbarBasicExample" class="navbar-menu">
                <div class="navbar-start">

                    {Menus.map((menu) => (
                        <a class="navbar-item" onClick={props.onClicked}>
                            {menu}
                        </a>
                    ))}

                </div>
            </div>
        </nav>
    );
}