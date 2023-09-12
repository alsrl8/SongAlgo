import "./Body.css";
import React, { useEffect, useState } from "react";
import Menu from "./menu/Menu.jsx";
import Schedule from "./schedule/Schedule.jsx";
import { GetMenu } from "../../wailsjs/go/main/App.js";

function Body({
  setIsModalOpen,
  setSubmitHistories,
  setSelectedProblemTitle,
  setSelectedProblemDate,
}) {
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
        <Menu menu={menu} setSelectedMenuItem={setSelectedMenuItem} />
      ) : (
        <Schedule
          selectedMenuItem={selectedMenuItem}
          setSelectedMenuItem={setSelectedMenuItem}
          setIsModalOpen={setIsModalOpen}
          setSubmitHistories={setSubmitHistories}
          setSelectedProblemTitle={setSelectedProblemTitle}
          setSelectedProblemDate={setSelectedProblemDate}
        />
      )}
    </>
  );
}

export default Body;
