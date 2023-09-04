import "./App.css";
import songAlgoLogo from "../src/assets/images/song_algo_logo_white.png";
import wailsLogo from "../src/assets/images/logo-universal.png";
import bjLogo from "../src/assets/images/bj_logo.png";
import pgLogo from "../src/assets/images/programers_logo.png";
import { GenerateCookieForBj, GetMenu } from "../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import Schedule from "./schedule/Schedule";

function App() {
  const [isBjConnected, setIsBjConnected] = useState(false);
  const [isPgConnected, setIsPgConnected] = useState(false);
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
      <div className="logos">
        <img src={songAlgoLogo} id="logo" alt="logo" />
        <img src={wailsLogo} id="logo" alt="logo" />
      </div>
      <div>
        {isBjConnected ? (
          <img src={bjLogo} className="cookieLogo activated" alt="logo" />
        ) : (
          <img
            src={bjLogo}
            className="cookieLogo deactivated"
            alt="logo"
            onClick={() => {
              GenerateCookieForBj().then((cookies) => {
                cookies.length > 0
                  ? setIsBjConnected(true)
                  : setIsBjConnected(false);
              });
            }}
          />
        )}
        {isPgConnected ? (
          <img
            src={pgLogo}
            className="cookieLogo activated"
            alt="logo"
            onClick={() => {}}
          />
        ) : (
          <img
            src={pgLogo}
            className="cookieLogo deactivated"
            alt="logo"
            onClick={() => {}}
          />
        )}
      </div>
      {selectedMenuItem === null ? (
        <div>
          {menu.map((item, index) => (
            <div
              className="interactive-text"
              key={index}
              onClick={(e) => {
                setSelectedMenuItem(item);
              }}
            >
              {item}{" "}
            </div>
          ))}
        </div>
      ) : (
        <Schedule
          selectedMenuItem={selectedMenuItem}
          setSelectedMenuItem={setSelectedMenuItem}
          isBjConnected={isBjConnected}
        />
      )}
    </div>
  );
}

export default App;
