import React, { useState } from "react";
import axios from "axios";
import { Layout, Menu } from "antd";
import { Input, Button, Row, Col, Typography } from "antd";
import { ArrowRightOutlined } from "@ant-design/icons";

import EIRP from "./EIRP";
import PathLoss from "./Path_Loss";
const { Content } = Layout;
const tabs = [
  { key: 1, label: "EIRP" },
  { key: 2, label: "Path Loss Free Space" },
  { key: 3, label: "Path Balance" },
  { key: 4, label: "Fersnel Zone" },
];

function RF() {
  const [selectedTab, setSelectedTab] = useState("1");
  const handleMenuClick = ({ key }) => {
    setSelectedTab(key);
  };

  const subNavbar = () => {
    switch (selectedTab) {
      case "1":
        return <EIRP />;
      case "2":
        return <PathLoss />;
      case "3":
        return <PathBalance />;
      case "4":
        return <FersnelZone />;
      default:
        return null;
    }
  };

  return (
    <>
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
      />
      {subNavbar()}
    </>
  );
}

const PathBalance = () => {
  return <>PathBalance</>;
};
const FersnelZone = () => {
  return <>FersnelZone</>;
};
export default RF;
