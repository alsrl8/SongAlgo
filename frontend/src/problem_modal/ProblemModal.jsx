import React from "react";
import "./ProblemModal.css";
import { Modal } from "antd";
import { UploadBjSourceToGithub } from "../../wailsjs/go/main/App.js";

const ProblemModal = ({
  isOpen,
  onClose,
  selectedProblemTitle,
  selectedProblemDate,
  submitHistories,
}) => {
  if (!isOpen) {
    return null;
  }

  let debounceTimeout = null;
  const handleClickCorrectHistory = (submission) => {
    if (debounceTimeout) return;
    showConfirm(submission);
    debounceTimeout = setTimeout(() => {
      clearTimeout(debounceTimeout);
      debounceTimeout = null;
    }, 500);
  };
  const showConfirm = (submission) => {
    Modal.confirm({
      title: "이 코드를 제출하시겠습니까?",
      content: "Github에 이미 해당 코드가 있다면 덮어쓰게 됩니다.",
      onOk() {
        UploadBjSourceToGithub(
          selectedProblemTitle,
          selectedProblemDate,
          submission,
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
