import React, { useState } from "react";
import "./ProblemModal.css";
import { Modal } from "antd";
import {
  GetGithubRepositoryBjSource,
  UploadBjSourceToGithub,
} from "../../wailsjs/go/main/App.js";

const ProblemModal = ({
  isOpen,
  onClose,
  userName,
  selectedProblemTitle,
  selectedProblemDate,
  submitHistories,
}) => {
  const [source, setSource] = useState(null);
  const [statusCode, setStatusCode] = useState("");

  if (!isOpen) {
    return null;
  }

  let debounceTimeout = null;
  const handleClickCorrectHistory = (submission) => {
    if (debounceTimeout) return;
    GetGithubRepositoryBjSource(
      selectedProblemTitle,
      selectedProblemDate,
      submission.ID,
      submission.Language,
    ).then(async (result) => {
      if (result.statusCode === "302") {
        showConfirmOverwriteCode(submission, result.file.sha);
      } else {
        showConfirmWriteCode(submission);
      }
      setSource(result.file); // TODO 필요한지 검토
      setStatusCode(result.statusCode); // TODO 필요한지 검토
    });

    debounceTimeout = setTimeout(() => {
      clearTimeout(debounceTimeout);
      debounceTimeout = null;
    }, 500);
  };
  const showConfirmOverwriteCode = (submission, sha) => {
    Modal.confirm({
      title: "이 코드를 제출하시겠습니까?",
      content:
        "Github에 이미 해당 코드가 있습니다. Ok 버튼을 누르면 코드를 덮어쓰게 됩니다.",
      onOk() {
        UploadBjSourceToGithub(
          selectedProblemTitle,
          selectedProblemDate,
          submission,
          sha,
          userName,
        );
      },
      onCancel() {},
    });
  };

  const showConfirmWriteCode = (submission) => {
    Modal.confirm({
      title: "이 코드를 제출하시겠습니까?",
      content: "",
      onOk() {
        UploadBjSourceToGithub(
          selectedProblemTitle,
          selectedProblemDate,
          submission,
          "",
          userName,
        );
      },
      onCancel() {},
    });
  };

  return (
    <div className="problem-modal-background">
      <div className="problem-modal" onClick={(e) => e.stopPropagation()}>
        <table className="submissionTable">
          <thead>
            <tr>
              <th>제출 번호</th>
              <th>결과</th>
              <th>메모리</th>
              <th>시간</th>
              <th>언어</th>
              <th>코드 길이</th>
              <th>제출한 시간</th>
            </tr>
          </thead>
          <tbody>
            {submitHistories.map((submission) => (
              <tr
                key={submission.SubmissionNumber}
                className={`submissionHistory ${
                  submission.Time ? "correct" : "wrong"
                }`}
                onClick={() => handleClickCorrectHistory(submission)}
              >
                <td>{submission.SubmissionNumber}</td>
                <td>{submission.Result}</td>
                <td>
                  {submission.Memory ? (
                    <>
                      {submission.Memory}
                      <span style={{ color: "orangered" }}> KB</span>
                    </>
                  ) : (
                    "N/A"
                  )}
                </td>
                <td>
                  {submission.Time ? (
                    <>
                      {submission.Time}
                      <span style={{ color: "orangered" }}> ms</span>
                    </>
                  ) : (
                    "N/A"
                  )}
                </td>
                <td>{submission.Language}</td>
                <td>
                  {submission.CodeLength ? (
                    <>
                      {submission.CodeLength}
                      <span style={{ color: "orangered" }}> B</span>
                    </>
                  ) : (
                    "N/A"
                  )}
                </td>
                <td>{submission.SubmissionTime}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <button className="close-button" onClick={onClose}>
          <h3>Close</h3>
        </button>
      </div>
    </div>
  );
};

export default ProblemModal;
