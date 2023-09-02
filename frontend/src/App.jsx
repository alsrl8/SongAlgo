import './App.css';
import songAlgoLogo from '../src/assets/images/song_algo_logo_white.png'
import wailsLogo from '../src/assets/images/logo-universal.png'
import bjLogo from '../src/assets/images/bj_logo.png'
import {GenerateCookieForBJ, GetMenu, } from '../wailsjs/go/main/App'
import {useEffect, useState} from "react";
import Schedule from "./schedule/Schedule";

function App() {
    const [bjCookie, setBjCookie] = useState("");
    const [isBJConnected, setIsBJConnected] = useState(false);
    const [menu, setMenu] = useState([]);
    const [selectedMenuItem, setSelectedMenuItem] = useState(null);

    // useEffect(() => {
    //     // Async function to handle async operations inside useEffect
    //     const fetchMenu = async () => {
    //         try {
    //             const fetchedMenu = await GetMenu();
    //             setMenu(fetchedMenu);
    //         } catch (error) {
    //             console.error('Failed to fetch menu:', error);
    //         }
    //     };
    //
    //     fetchMenu();
    // }, []);

    useEffect(() => {
        GetMenu().then((menu) => {
            setMenu(menu);
        });
    }, []);

    return (
        <div id="App">
            <div className='logos'>
                <img src={songAlgoLogo} id="logo" alt="logo"/>
                <img src={wailsLogo} id="logo" alt="logo"/>
            </div>
            <div>
                {isBJConnected ? <img src={bjLogo} className="cookieLogo activated" alt="logo"></img> :
                    <img src={bjLogo} className="cookieLogo deactivated" alt="logo" onClick={() => {
                        GenerateCookieForBJ().then((cookie) => {
                            cookie.value !== "" ? setIsBJConnected(true) : setIsBJConnected(false)
                            setBjCookie(cookie.value)
                        })
                    }}></img>}
            </div>
            {selectedMenuItem === null ? (<div>
                    {menu.map((item, index) => (
                        <div className="interactive-text" key={index} onClick={(e) => {
                            setSelectedMenuItem(item)
                        }}>{item}</div>
                    ))}
                </div>) :
                <Schedule selectedMenuItem={selectedMenuItem} setSelectedMenuItem={setSelectedMenuItem} isBJConnected={isBJConnected} bjCookie={bjCookie}/>
            }
        </div>
    )
}

export default App
