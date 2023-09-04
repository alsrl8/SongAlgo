import React, { useEffect, useState } from "react";
import { GetSchedule, OpenBjWithCookie } from "../../wailsjs/go/main/App";
import "./Schedule.css";

function Schedule({ selectedMenuItem, setSelectedMenuItem }) {
  const [scheduleList, setScheduleList] = useState([]);

  useEffect(() => {
    GetSchedule().then((_scheduleList) => {
      setScheduleList(_scheduleList.list);
    });
  }, []);

  return (
    <div className="scheduleContainer">
      <h2>{selectedMenuItem}</h2>
      {scheduleList.map((item, index) => (
        <div key={index} className="scheduleCard">
          <span className="date">{item.date}</span>
          {item.problems.map((problem, pi) => (
            <div
              onClick={() => {
                window.open(problem.url, "_blank");
              }}
            >
              <div
                className={`problem ${
                  pi + 1 === item.problems.length ? "problemLast" : ""
                }`}
              >
                <span className="problemDetail title">
                  {pi + 1}. {problem.name}
                </span>
                <span className="problemDetail platform">
                  {problem.algorithmType}
                </span>
                <span className="problemDetail platform">
                  {problem.platform} - {problem.difficulty}
                </span>
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
