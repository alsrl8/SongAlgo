import "./App.css";
import songAlgoLogo from "../src/assets/images/song_algo_logo_white.png";
import wailsLogo from "../src/assets/images/logo-universal.png";
import bjLogo from "../src/assets/images/bj_logo.png";
import pgLogo from "../src/assets/images/programers_logo.png";
import { GetMenu } from "../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import Schedule from "./schedule/Schedule";

function App() {
  const [menu, setMenu] = useState([]);
  const [selectedMenuItem, setSelectedMenuItem] = useState(null);

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
        <img src={bjLogo} className="cookieLogo" alt="logo" />
        <img src={pgLogo} className="cookieLogo" alt="logo" />
      </div>
      {selectedMenuItem === null ? (
        <div>
          {menu.map((item, index) => (
            <div
              className="interactive-text"
              key={index}
              onClick={() => {
                setSelectedMenuItem(item);
              }}
            >
              {item}
            </div>
          ))}
        </div>
      ) : (
        <Schedule
          selectedMenuItem={selectedMenuItem}
          setSelectedMenuItem={setSelectedMenuItem}
        />
      )}
    </div>
  );
}

export default App;
