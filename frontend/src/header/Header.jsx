import "./Header.css";
import songAlgoLogo from "./../assets/images/song_algo_logo_white.png";
import wailsLogo from "./../assets/images/logo-universal.png";
import bjLogo from "./../assets/images/bj_logo.png";
import pgLogo from "./../assets/images/programers_logo.png";
import React, { useState } from "react";
import {
  IsBjLoggedIn,
  CloseSeleniumBrowser,
  IsPgLoggedIn,
  NavigateToPgLoginPage,
  NavigateToBjLoginPage,
} from "../../wailsjs/go/main/App.js";
import { Modal } from "antd";

const showWarningAlreadyLoggedIn = (domain) => {
  Modal.warning({
    title: "이미 로그인 되어 있습니다.",
    content: (
      <div>
        {domain} 사이트에 <br />
        이미 로그인 되어 있습니다.
      </div>
    ),
  });
};
const handleBjLogin = async (setIsLoading, setLoadingText) => {
  setIsLoading(true);
  setLoadingText("백준 로그인 상태를 확인하고 있습니다.");
  await CloseSeleniumBrowser().then(() => {
    IsBjLoggedIn("https://www.acmicpc.net/problem/1000").then((result) => {
      setIsLoading(false);
      if (result === true) {
        showWarningAlreadyLoggedIn("백준");
        return;
      }
      CloseSeleniumBrowser().then(() => {
        NavigateToBjLoginPage();
      });
    });
  });
};
const handlePgLogin = async (setIsLoading, setLoadingText) => {
  setIsLoading(true);
  setLoadingText("프로그래머스 로그인 상태를 확인하고 있습니다.");
  await CloseSeleniumBrowser().then(() => {
    IsPgLoggedIn(
      "https://school.programmers.co.kr/learn/courses/30/lessons/1829",
    ).then((result) => {
      setIsLoading(false);
      if (result === true) {
        showWarningAlreadyLoggedIn("Programmers");
        return;
      }
      CloseSeleniumBrowser().then(() => {
        NavigateToPgLoginPage();
      });
    });
  });
};

function Header({
  userId,
  setIsLoading,
  setLoadingText,
  setIsSettingModalOpen,
}) {
  return (
    <>
      <div className="header">
        <div className="headerLeft">
          {userId !== "" ? "User Name: " + userId : "NO USER"}
        </div>
        <div className="headerRight">
          <img
            src={bjLogo}
            id="bjLogo"
            className="siteLogo"
            alt="logo"
            onClick={() => handleBjLogin(setIsLoading, setLoadingText)}
          />
          <img
            src={pgLogo}
            id="pgLogo"
            className="siteLogo"
            alt="logo"
            onClick={() => handlePgLogin(setIsLoading, setLoadingText)}
          />
          <img
            src={songAlgoLogo}
            id="songAlgoLogo"
            className="logo"
            alt="logo"
            onClick={() => {
              setIsSettingModalOpen(true);
            }}
          />
        </div>
      </div>
    </>
  );
}

export default Header;
