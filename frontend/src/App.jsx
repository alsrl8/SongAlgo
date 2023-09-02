import './App.css';
import logo from '../src/assets/images/song_algo_logo_white.png'
import {GetMenu} from '../wailsjs/go/main/App'
import {useEffect, useState} from "react";
import Schedule from "./schedule/Schedule";

function App() {
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
            <img src={logo} id="logo" alt="logo"/>
            {selectedMenuItem === null ? (<div>
                    {menu.map((item, index) => (
                        <div className="interactive-text" key={index} onClick={(e) => {
                            setSelectedMenuItem(item)
                        }}>{item}</div>
                    ))}
                </div>) :
                <Schedule selectedMenuItem={selectedMenuItem} setSelectedMenuItem={setSelectedMenuItem}/>
            }
        </div>
    )
}

export default App
