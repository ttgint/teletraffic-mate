import React from "react";

import { Layout } from "antd";
const { Footer } = Layout;

function MyFooter() {
  return (
    <Footer
      style={{
        textAlign: "center",
      }}
    >
      Teletraffic Mate ©{new Date().getFullYear()} Created by TTG Internatinal LTD
    </Footer>
  );
}

export default MyFooter;
