import React, { useEffect, useState } from "react";
import {
  GetSchedule,
  NavigateToBjProblemWithCookie,
  IsPgLoggedIn,
  IsSubmittedCodeCorrect,
  UploadPgSourceToGithub,
  GetPgSourceData,
  GetGithubRepositoryPgSource,
  UploadBjSourceToGithub,
} from "../../../wailsjs/go/main/App.js";
import "./Schedule.css";
import cdLogo from "../../assets/images/code_logo.png";
import { Modal } from "antd";

function Schedule({
  selectedMenuItem,
  setSelectedMenuItem,
  setIsModalOpen,
  setIsLoading,
  setLoadingText,
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

  const showWarningPgLogin = () => {
    Modal.warning({
      title: "코드를 제출할 수 없습니다.",
      content: (
        <div>
          브라우저가 Programmers에
          <br />
          로그인되어 있지 않습니다.
        </div>
      ),
    });
  };

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

  const showConfirmOverwriteCode = (
    problemTitle,
    problemDate,
    githubId,
    code,
    extension,
    sha,
  ) => {
    Modal.confirm({
      title: "이 코드를 제출하시겠습니까?",
      content:
        "Github에 이미 해당 코드가 있습니다. Ok 버튼을 누르면 코드를 덮어쓰게 됩니다.",
      onOk() {
        UploadPgSourceToGithub(
          problemTitle,
          problemDate,
          githubId,
          code,
          extension,
          sha,
        );
      },
      onCancel() {},
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
          "",
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
                      setIsLoading(true);
                      setLoadingText("백준 제출 이력을 읽어오고 있습니다.");
                      await NavigateToBjProblemWithCookie(problem.url).then(
                        (_submitHistories) => {
                          setSubmitHistories(_submitHistories);
                          setSelectedProblemTitle(problem.name);
                          setSelectedProblemDate(item.date);
                          setIsModalOpen(true);
                          setIsLoading(false);
                        },
                      );
                    } else if (problem.platform === "programmers") {
                      setIsLoading(true);
                      setLoadingText(
                        "프로그래머스 제출 이력을 읽어오고 있습니다.",
                      );
                      await IsPgLoggedIn(problem.url).then((result) => {
                        if (result === false) {
                          setIsLoading(false);
                          showWarningPgLogin();
                          return;
                        }
                        IsSubmittedCodeCorrect(problem.url).then((result) => {
                          if (result === false) {
                            setIsLoading(false);
                            showWarningPgCode();
                            return;
                          }
                          setLoadingText(
                            "프로그래머스 코드가 Github에 업로드 됐는지 확인하고 있습니다.",
                          );
                          GetPgSourceData(problem.url).then((pgSourceData) => {
                            GetGithubRepositoryPgSource(
                              problem.name,
                              item.date,
                              "alsrl8",
                              pgSourceData.extension,
                            ).then((fileResponse) => {
                              setIsLoading(false);
                              if (fileResponse.statusCode === "302") {
                                showConfirmOverwriteCode(
                                  problem.name,
                                  item.date,
                                  "alsrl8",
                                  pgSourceData.code,
                                  pgSourceData.extension,
                                  fileResponse.file.sha,
                                );
                              } else {
                                showConfirmSubmitPgCode(
                                  problem.name,
                                  item.date,
                                  "alsrl8",
                                  pgSourceData.code,
                                  pgSourceData.extension,
                                );
                              }
                            });
                          });
                        });
                      });
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
