import "./App.css";
import Header from "./header/Header.jsx";
import Body from "./body/Body.jsx";
import React from "react";

function App() {
  return (
    <div id="App">
      <div className="header-container">
        <Header />
      </div>
      <Body />
    </div>
  );
}

export default App;
