import React, { useEffect, useState } from "react";
import {
  GetSchedule,
  NavigateToBjProblemWithCookie,
  IsChromeRunning,
} from "../../../wailsjs/go/main/App.js";
import "./Schedule.css";
import cdLogo from "../../assets/images/code_logo.png";
import githubLogo from "../../assets/images/github-logo.png";

function Schedule({
  selectedMenuItem,
  setSelectedMenuItem,
  setIsModalOpen,
  setSubmitHistories,
  setSelectedProblemTitle,
  setSelectedProblemDate,
}) {
  const [scheduleList, setScheduleList] = useState([]);

  useEffect(() => {
    GetSchedule().then((_scheduleList) => {
      setScheduleList(_scheduleList.list);
    });
  }, []);

  const convertDateToYYMMDD = (dateString) => {
    const parts = dateString.split("-");
    return parts[0].substring(2) + parts[1] + parts[2];
  };

  return (
    <div className="scheduleContainer">
      <h2>{selectedMenuItem}</h2>
      {scheduleList.map((item, index) => (
        <div key={"schedule" + index} className="scheduleCard">
          <span className="date">{item.date}</span>
          {item.problems.map((problem, pi) => (
            <div className="problem" key={"schedule" + index + "problem" + pi}>
              <div
                className={`problemContents ${
                  pi + 1 === item.problems.length ? "last" : ""
                }`}
                onClick={async () => {
                  window.open(problem.url, "_blank");
                }}
              >
                <div className="problemContent">
                  <span className="problemContent title">
                    {pi + 1}. {problem.name}
                  </span>
                  <span className="problemContent platform">
                    {problem.algorithmType}
                  </span>
                  <span className="problemContent platform">
                    {problem.platform} - {problem.difficulty}
                  </span>
                </div>
              </div>
              <div className="problemFeatures">
                <img
                  src={cdLogo}
                  alt="logo"
                  className="logo"
                  onClick={async () => {
                    if (problem.platform === "baekjoon") {
                      await NavigateToBjProblemWithCookie(problem.url).then(
                        (_submitHistories) => {
                          if (_submitHistories.length === 0) {
                            return;
                          }
                          setSubmitHistories(_submitHistories);
                          setSelectedProblemTitle(problem.name);
                          setSelectedProblemDate(
                            convertDateToYYMMDD(item.date),
                          );
                          setIsModalOpen(true);
                        },
                      );
                    }
                  }}
                />
                <img src={githubLogo} alt="logo" className="logo" />
              </div>
            </div>
          ))}
        </div>
      ))}
      <button
        className="goBackButton"
        onClick={() => {
          setSelectedMenuItem(null);
        }}
      >
        Go Back
      </button>
    </div>
  );
}

export default Schedule;
