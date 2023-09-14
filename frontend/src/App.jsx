import "./App.css";
import Header from "./header/Header.jsx";
import Body from "./body/Body.jsx";
import React, { useState } from "react";
import ProblemModal from "./problem_modal/ProblemModal.jsx";
import { Modal } from "antd";

function App() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [submitHistories, setSubmitHistories] = useState([]);
  const [selectedProblemTitle, setSelectedProblemTitle] = useState("");
  const [selectedProblemDate, setSelectedProblemDate] = useState("");
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
        setSelectedProblemTitle={setSelectedProblemTitle}
        setSelectedProblemDate={setSelectedProblemDate}
      />
      <ProblemModal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        selectedProblemTitle={selectedProblemTitle}
        selectedProblemDate={selectedProblemDate}
        submitHistories={submitHistories}
      />
    </div>
  );
}

export default App;
