import React from "react";
import "./ProblemModal.css";

const ProblemModal = ({ isOpen, onClose, submitHistories }) => {
  if (!isOpen) {
    return null;
  }

  if (submitHistories.length > 0) {
    console.log(submitHistories);
  }

  return (
    <div className="modal-background">
      <div className="modal" onClick={(e) => e.stopPropagation()}>
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
