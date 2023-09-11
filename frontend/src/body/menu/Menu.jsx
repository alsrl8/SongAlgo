import "./Menu.css";
import React from "react";

function Menu({ menu, setSelectedMenuItem }) {
  return (
    <>
      {menu.map((item, index) => (
        <div
          className="interactive-text"
          key={"menu" + index}
          onClick={() => {
            switch (index) {
              case 0:
                setSelectedMenuItem(item);
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
