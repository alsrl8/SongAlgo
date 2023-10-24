import "./App.css";
import Header from "./header/Header.jsx";
import Body from "./body/Body.jsx";
import React, { useState } from "react";
import ProblemModal from "./problem_modal/ProblemModal.jsx";
import Loading from "./loading/loading.jsx";
import Setting from "./setting/setting.jsx";

function App() {
  const [userId, setUserId] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [loadingText, setLoadingText] = useState("");
  const [isProblemModalOpen, setIsProblemModalOpen] = useState(false);
  const [submitHistories, setSubmitHistories] = useState([]);
  const [selectedProblemTitle, setSelectedProblemTitle] = useState("");
  const [selectedProblemDate, setSelectedProblemDate] = useState("");
  const [isSettingModalOpen, setIsSettingModalOpen] = useState(false);

  const handleCloseProblemModal = () => {
    setIsProblemModalOpen(false);
  };

  const handleCloseSettingModal = () => {
    setIsSettingModalOpen(false);
  };

  return (
    <div id="App">
      <Loading isLoading={isLoading} loadingText={loadingText} />
      <div className="header-container">
        <Header
          userId={userId}
          setIsLoading={setIsLoading}
          setLoadingText={setLoadingText}
          setIsSettingModalOpen={setIsSettingModalOpen}
        />
      </div>
      <Body
        setIsModalOpen={setIsProblemModalOpen}
        setIsLoading={setIsLoading}
        setLoadingText={setLoadingText}
        setSubmitHistories={setSubmitHistories}
        setSelectedProblemTitle={setSelectedProblemTitle}
        setSelectedProblemDate={setSelectedProblemDate}
      />
      <ProblemModal
        isOpen={isProblemModalOpen}
        onClose={handleCloseProblemModal}
        selectedProblemTitle={selectedProblemTitle}
        selectedProblemDate={selectedProblemDate}
        submitHistories={submitHistories}
      />
      <Setting
        isSettingModalOpen={isSettingModalOpen}
        onClose={handleCloseSettingModal}
        setUserId={setUserId}
      />
    </div>
  );
}

export default App;
