import React, { useState } from "react";
import "./setting.css";
import { Input } from "antd";

const forbiddenUsernames = ["", "main", "schedule", "base_user"];

const Setting = ({ isSettingModalOpen, onClose, setUserName }) => {
  if (!isSettingModalOpen) return null;

  const [inputUserName, setInputUserName] = useState("");
  const [error, setError] = useState("");

  const handleOkButtonClick = () => {
    if (!inputUserName) {
      setError("Username cannot be empty.");
      return;
    }
    if (forbiddenUsernames.includes(inputUserName.toLowerCase())) {
      setError("This username is not allowed.");
      return;
    }
    setUserName(inputUserName);
    onClose();
  };

  const handleUserNameInputChange = (e) => {
    setInputUserName(e.target.value);
    if (error) setError("");
  };

  return (
    <div className="setting">
      <div className="background">
        <div className="modal">
          <button className="ok-button" onClick={handleOkButtonClick}>
            <h3>OK</h3>
          </button>
          <button className="close-button" onClick={onClose}>
            <h3>Close</h3>
          </button>
          <div className="setting-header">
            <h3>Setting</h3>
          </div>
          <div className="setting-body">
            Github Repository
            <Input value="https://github.com/alsrl8/SongAlgo" disabled />
            User Name
            <Input
              className={error ? "setting-input-error" : ""}
              placeholder="사용자 이름"
              onChange={handleUserNameInputChange}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Setting;
