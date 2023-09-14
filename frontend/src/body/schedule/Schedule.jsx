import React, { useEffect, useState } from "react";
import {
  GetSchedule,
  NavigateToBjProblemWithCookie,
  IsSubmittedCodeCorrect,
  UploadPgSourceToGithub,
  GetPgSourceData,
} from "../../../wailsjs/go/main/App.js";
import "./Schedule.css";
import cdLogo from "../../assets/images/code_logo.png";
import { Modal } from "antd";

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

  const showWarningPgCode = () => {
    Modal.warning({
      title: "코드를 제출할 수 없습니다.",
      content: (
        <div>
          이 코드는 오답 판정을 받았기 때문에
          <br />
          Github에 올릴 수 없습니다.
        </div>
      ),
    });
  };

  const showConfirmSubmitPgCode = (
    problemTitle,
    problemDate,
    githubId,
    code,
    extension,
  ) => {
    Modal.confirm({
      title: "코드를 제출하시겠습니까?",
      content: "",
      onOk() {
        UploadPgSourceToGithub(
          problemTitle,
          problemDate,
          githubId,
          code,
          extension,
        );
      },
      onCancel() {},
    });
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
                          setSubmitHistories(_submitHistories);
                          setSelectedProblemTitle(problem.name);
                          setSelectedProblemDate(item.date);
                          setIsModalOpen(true);
                        },
                      );
                    } else if (problem.platform === "programmers") {
                      await IsSubmittedCodeCorrect(problem.url).then(
                        (result) => {
                          if (result === false) {
                            showWarningPgCode();
                            return;
                          }
                          // TODO Github에 이미 제출된 코드 이력이 있는지 조사
                          // TODO 프로그래머스 code를 받아오도록 수정
                          GetPgSourceData(problem.url).then((pgSourceData) => {
                            showConfirmSubmitPgCode(
                              problem.name,
                              item.date,
                              "alsrl8",
                              pgSourceData.code,
                              pgSourceData.extension,
                            );
                          });
                        },
                      );
                    }
                  }}
                />
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
