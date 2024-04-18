import React from "react";
import { useState } from "react";
import { Layout, Divider } from "antd";
const { Content } = Layout;
import { Input, Button, Row, Col, Typography } from "antd";
import axios from "axios";
import { ArrowRightOutlined } from "@ant-design/icons";

function NetworkCalcsContent() {
  return (
    <>
      <div style={{ display: "flex", height: "100%", flexDirection: "row" }}>
        <div style={{ padding: "20px" }}>
          <div>
            <BasicBSCCalculations />
            <Divider />
          </div>
          <div>
            <BasicTrafficCalculations />
          </div>
        </div>
        <div style={{ padding: "20px" }}>
          <BasicNumTRXCalculations />
        </div>
      </div>
    </>
  );
}

function BasicBSCCalculations() {
  const [bhca, setBHCA] = useState("");
  const [sms, setSMS] = useState("");
  const [reg, setregistration] = useState("");
  const [callLoad, setCPCallLoad] = useState("");
  const [capacity, setCapacity] = useState("");

  const handleApiCall = async () => {
    try {
      const res = await axios.post(
        `http://localhost:5174/network/simple-bsc/${bhca}/${sms}/${reg}/${callLoad}`
      );
      setCapacity(res.data);
    } catch (error) {
      console.log(error.response);
    }
  };

  return (
    <>
      <Typography.Title level={4}>Simple BSC Capacity</Typography.Title>
      <Row gutter={[16, 16]}>
        <Col span={12}>
          <Typography.Title level={5}>BHCA / Sub.</Typography.Title>
          <div style={{ display: "flex",  alignItems: "center"  }}>
            <Input value={bhca} onChange={(e) => setBHCA(e.target.value)} />
            <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>bh</sub>
          </div>
        </Col>
        <Col span={12}>
          <Typography.Title level={5}>SMS</Typography.Title>
          <Input value={sms} onChange={(e) => setSMS(e.target.value)} />
        </Col>
        <Col span={12}>
          <Typography.Title level={5}>Registration</Typography.Title>
          <Input
            value={reg}
            onChange={(e) => setregistration(e.target.value)}
          />
        </Col>
        <Col span={12}>
          <Typography.Title level={5}>CP Call Load</Typography.Title>
          <Input
            value={callLoad}
            onChange={(e) => setCPCallLoad(e.target.value)}
          />
        </Col>
        <Col span={24}>
          <Typography.Title level={5}>capacity</Typography.Title>
          <div style={{ display: "flex", alignItems: "center" }}>
            <Button
              type="primary"
              shape="round"
              onClick={handleApiCall}
              style={{ borderRadius: "10px 0px 0px 10px" }}
              icon={<ArrowRightOutlined />}
            />
            <Input
              value={capacity}
              style={{ borderRadius: "0px 10px 10px 0px" }}
            />
          </div>
        </Col>
      </Row>
    </>
  );
}

function BasicTrafficCalculations() {
  const [bhca, setBHCA] = useState("");
  const [mht, setMHT] = useState("");
  const [traffic, setTraffic] = useState("");

  const handleApiCall = async () => {
    try {
      const res = await axios.post(
        `http://localhost:5174/network/basic-traffic/${bhca}/${mht}`
      );
      setTraffic(res.data);
    } catch (error) {
      console.log(error.response);
    }
  };
  return (
    <>
      <Typography.Title level={4}>Basic Traffic</Typography.Title>
      <Row gutter={[16, 16]}>
        <Col span={8}>
          <Typography.Title level={5}>BHCA / Sub.</Typography.Title>
          <Input value={bhca} onChange={(e) => setBHCA(e.target.value)} />
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>MHT</Typography.Title>
          <div style={{ display: "flex",  alignItems: "center"  }}>
            <Input value={mht} onChange={(e) => setMHT(e.target.value)} />
            <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>sec</sub>
          </div>
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>Traffic</Typography.Title>
          <div style={{ display: "flex", alignItems: "center" }}>
            <Button
              type="primary"
              shape="round"
              icon={<ArrowRightOutlined />}
              onClick={handleApiCall}
              style={{ borderRadius: "10px 0px 0px 10px" }}
            />
            <Input
              value={traffic}
              onChange={(e) => setregistration(e.target.value)}
              style={{ borderRadius: "0px 10px 10px 0px" }}
            />
            <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>mErl</sub>
          </div>
        </Col>
      </Row>
    </>
  );
}

function BasicNumTRXCalculations() {
  const [subs, setSubs] = useState("");
  const [subs_per_erl, setSubsPerErl] = useState("");
  const [trx_erl, setTrxErl] = useState("");

  const [trx, setTrx] = useState("");
  const [trx_cell, setTrxCell] = useState("");

  const [cells, setCells] = useState("");
  const handleApiCall_1 = async () => {
    try {
      const res = await axios.post(
        `http://localhost:5174/network/trx/${subs}/${subs_per_erl}/${trx_erl}`
      );
      setTrx(res.data);
    } catch (error) {
      console.log(error.response.data);
    }
  };

  const handleApiCall_2 = async () => {
    try {
      const trxString = trx.toString();
      const res = await axios.post(
        `http://localhost:5174/network/trx/${trxString}/${trx_cell}`
      );
      setCells(res.data);
    } catch (error) {
      console.log(error.response);
    }
  };

  return (
    <>
      <Typography.Title level={4}>No. of TRX </Typography.Title>
      <Row gutter={[16, 16]}>
        <Col span={8}>
          <Typography.Title level={5}>No. Subs.</Typography.Title>
          <Input value={subs} onChange={(e) => setSubs(e.target.value)} />
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>Erl/Subs.</Typography.Title>
          <Input
            value={subs_per_erl}
            onChange={(e) => setSubsPerErl(e.target.value)}
          />
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>Erl / TRX</Typography.Title>
          <Input value={trx_erl} onChange={(e) => setTrxErl(e.target.value)} />
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>No. Of TRX</Typography.Title>
          <div style={{ display: "flex", alignItems: "center" }}>
            <Button
              type="primary"
              shape="round"
              onClick={handleApiCall_1}
              style={{ borderRadius: "10px 0px 0px 10px" }}
              icon={<ArrowRightOutlined />}
            />
            <Input
              value={trx}
              onChange={(e) => setTrx(e.target.value)}
              style={{ borderRadius: "0px 10px 10px 0px" }}
            />
          </div>
        </Col>
        <Col span={8}>
          <Typography.Title level={5}>TRX / Cell</Typography.Title>
          <Input
            value={trx_cell}
            onChange={(e) => setTrxCell(e.target.value)}
          />
        </Col>
        <Col span={12}>
          <Typography.Title level={5}>No. Of Cells</Typography.Title>
          <div style={{ display: "flex", alignItems: "center" }}>
            <Button
              type="primary"
              shape="round"
              icon={<ArrowRightOutlined />}
              onClick={handleApiCall_2}
              style={{ borderRadius: "10px 0px 0px 10px" }}
            />
            <Input
              value={cells}
              onChange={(e) => setCells(e.target.value)}
              style={{ borderRadius: "0px 10px 10px 0px" }}
            />
          </div>
        </Col>
      </Row>
    </>
  );
}

export default NetworkCalcsContent;
