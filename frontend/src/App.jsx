import "./App.css";
import Header from "./header/Header.jsx";
import Body from "./body/Body.jsx";
import React, { useState } from "react";
import ProblemModal from "./problem_modal/ProblemModal.jsx";

function App() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [submitHistories, setSubmitHistories] = useState([]);
  const handleCloseModal = () => {
    setIsModalOpen(false);
  };
  return (
    <div id="App">
      <div className="header-container">
        <Header />
      </div>
      <Body
        setIsModalOpen={setIsModalOpen}
        setSubmitHistories={setSubmitHistories}
      />
      <ProblemModal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        submitHistories={submitHistories}
      />
    </div>
  );
}

export default App;
