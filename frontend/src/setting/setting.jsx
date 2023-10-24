import React, { useState } from "react";
import "./setting.css";
import { Input } from "antd";

const Setting = ({ isSettingModalOpen, onClose, setUserId }) => {
  if (!isSettingModalOpen) return null;

  const [userName, setUserName] = useState("");

  const handleOkButtonClick = () => {
    setUserId(userName);
    onClose();
  };

  const handleUserNameInputChange = (e) => {
    setUserName(e.target.value);
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
