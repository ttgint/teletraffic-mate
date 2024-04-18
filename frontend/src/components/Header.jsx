import React from "react";
import { useState } from "react";

import { Layout, Menu } from "antd";
const { Header } = Layout;

const tabs = [
  { key: 1, label: "Traffic" },
  { key: 2, label: "Network Calculations" },
  { key: 3, label: "RF" },
  { key: 4, label: "Trunking" },
  { key: 5, label: "Multi-Link Blocking" },
  { key: 6, label: "Units" },
];

const headerStyle = {
  display: "flex",
  padding: 0,
};
function MyHeader({ onSelectTab }) {
  const [selectedTab, setSelectedTab] = useState("1");
  const handleMenuClick = ({ key }) => {
    onSelectTab(key);
  };
  return (
    <Header style={headerStyle}>
      <>
        <img src="../ttg.gif" alt="Logo" />

        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["1"]}
          selectedKeys={[selectedTab]}
          onClick={handleMenuClick}
          items={tabs}
          style={{
            flex: 1,
            minWidth: 0,
            fontSize: "10px",
          }}
        >
          {tabs.map((tab) => (
            <Menu.Item key={tab.key}>{tab.label}</Menu.Item>
          ))}
        </Menu>
      </>
    </Header>
  );
}

export default MyHeader;
