import "./Body.css";
import React, { useEffect, useState } from "react";
import Menu from "./menu/Menu.jsx";
import Schedule from "./schedule/Schedule.jsx";
import Add from "./add/Add.jsx";
import { GetMenu } from "../../wailsjs/go/main/App.js";

function Body({
  setIsModalOpen,
  setIsLoading,
  setLoadingText,
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
      ) : selectedMenuItem === 0 ? (
        <Schedule
          selectedMenuItem={selectedMenuItem}
          setSelectedMenuItem={setSelectedMenuItem}
          setIsModalOpen={setIsModalOpen}
          setIsLoading={setIsLoading}
          setLoadingText={setLoadingText}
          setSubmitHistories={setSubmitHistories}
          setSelectedProblemTitle={setSelectedProblemTitle}
          setSelectedProblemDate={setSelectedProblemDate}
        />
      ) : (
        <Add setSelectedMenuItem={setSelectedMenuItem} />
      )}
    </>
  );
}

export default Body;
