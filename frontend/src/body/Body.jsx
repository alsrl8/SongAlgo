import "./Body.css";
import React, { useEffect, useState } from "react";
import Schedule from "../schedule/Schedule.jsx";
import { GetMenu } from "../../wailsjs/go/main/App.js";

function Body() {
  const [menu, setMenu] = useState([]);
  const [selectedMenuItem, setSelectedMenuItem] = useState(null);

  useEffect(() => {
    GetMenu().then((menu) => {
      setMenu(menu);
    });
  }, []);

  return (
    <>
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
    </>
  );
}

export default Body;
