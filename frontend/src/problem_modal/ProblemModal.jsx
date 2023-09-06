import React from "react";
import "./ProblemModal.css";

const ProblemModal = ({ isOpen, onClose, children }) => {
  if (!isOpen) {
    return null;
  }

  return (
    <div className="modal-background">
      <div className="modal" onClick={(e) => e.stopPropagation()}>
        {children}
        <button onClick={onClose}>Close</button>
      </div>
    </div>
  );
};

export default ProblemModal;
