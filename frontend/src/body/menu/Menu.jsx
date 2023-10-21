import "./Menu.css";
import React from "react";
import { CloseProgram } from "../../../wailsjs/go/main/App.js";

function Menu({ menu, setSelectedMenuItem }) {
  return (
    <>
      {menu.map((item, index) => (
        <div
          className="interactive-text"
          key={"menu" + index}
          onClick={() => {
            switch (index) {
              case 0: // Problem list
                setSelectedMenuItem(item);
                break;
              case 1: // Close program
                CloseProgram();
                break;
            }
          }}
        >
          {item}
        </div>
      ))}
    </>
  );
}

export default Menu;
