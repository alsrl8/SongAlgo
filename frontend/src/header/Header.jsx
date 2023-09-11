import "./Header.css";
import songAlgoLogo from "./../assets/images/song_algo_logo_white.png";
import wailsLogo from "./../assets/images/logo-universal.png";
import bjLogo from "./../assets/images/bj_logo.png";
import pgLogo from "./../assets/images/programers_logo.png";
import React from "react";

function Header() {
  return (
    <>
      <div className="header">
        <img src={bjLogo} id="bjLogo" className="siteLogo" alt="logo" />
        <img src={pgLogo} id="pgLogo" className="siteLogo" alt="logo" />
        <img src={songAlgoLogo} id="songAlgoLogo" className="logo" alt="logo" />
        <img src={wailsLogo} id="wailsLogo" className="logo" alt="logo" />
      </div>
    </>
  );
}

export default Header;
