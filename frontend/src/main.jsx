import React from "react";
import { useState } from "react";
import { createRoot } from "react-dom/client";
import { Layout } from "antd";
import TrafficContent  from "./components/Traffic";
import NetworkCalcsContent from "./components/Network_calculations";
import RF from "./components/RF/RF"
import MyFooter from "./components/Footer";
import MyHeader from "./components/Header";

function App() {
  const [selectedTab, setSelectedTab] = useState("2");

  const renderContent = () => {
    switch (selectedTab) {
      case "1":
        return <TrafficContent />;
      case "2":
        return <NetworkCalcsContent />;
      case "3":
        return <RF />;
      // case "4":
      //   return <TrunkingContent />;
      // case "5":
      //   return <MultiLinkBlockingContent />;
      // case "6":
      //   return <UnitsContent />;
      default:
        return null;
    }
  };

  return (
    <Layout>
      <MyHeader onSelectTab={(key) => setSelectedTab(key)}  />
      {renderContent()}
      <MyFooter />
    </Layout>
  );
}
const container = document.getElementById("root");

const root = createRoot(container);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
