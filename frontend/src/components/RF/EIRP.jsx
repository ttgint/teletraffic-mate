import React, { useState } from "react";
import axios from "axios";
import { Input, Button, Row, Col, Typography } from "antd";
import { ArrowRightOutlined } from "@ant-design/icons";

const EIRP = () => {
    const [eirp, setEIRP] = useState("");
    const [power, setPower] = useState("");
    const [combiner, setCombiner] = useState("");
    const [feeder, setFeeder] = useState("");
    const [antenna, setAtntenna] = useState("");
  
    const handleApiCall = async (
      endpoint,
      param_1,
      param_2,
      param_3,
      param_4,
      setStateFunction
    ) => {
      try {
        const res = await axios.post(
          `http://localhost:5174/rf/${endpoint}/${param_1}/${param_2}/${param_3}/${param_4}`
        );
        setStateFunction(res.data);
      } catch (error) {
        console.log(error.response);
      }
    };
  
    const handleEIRPCall      = async () => handleApiCall("eirp", power, combiner, feeder, antenna, setEIRP);
    const handlePowerCall     = async () => handleApiCall("power", eirp, combiner, feeder, antenna, setPower);
    const handleCombinerCall  = async () => handleApiCall("combiner", eirp, power, feeder, antenna, setCombiner);
    const handleFeederCall    = async () => handleApiCall("feeder", eirp, power, combiner, antenna, setFeeder);
    const handleAntennaCall   = async () => handleApiCall("antenna", eirp, power, combiner, feeder, setAtntenna);
  
    return (
      <div style={{marginLeft: "50px"}}>
        <Row className="gutter-row" span={8}>
          <Col>
            <Typography.Title level={5}>EIRP</Typography.Title>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Button
                type="primary"
                shape="round"
                onClick={handleEIRPCall}
                icon={<ArrowRightOutlined />}
                style={{ borderRadius: "10px 0px 0px 10px" }}
              />
              <Input
                value={eirp}
                onChange={(e) => setEIRP(e.target.value)}
                style={{ borderRadius: "0px 10px 10px 0px" }}
              />
              <div>
                <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
                  [dBm]
                </sub>
              </div>
            </div>
          </Col>
        </Row>
        <Row className="gutter-row" span={8}>
          <Col>
            <Typography.Title level={5}>Transmitted Power in BTS</Typography.Title>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Button
                type="primary"
                shape="round"
                onClick={handlePowerCall}
                icon={<ArrowRightOutlined />}
                style={{ borderRadius: "10px 0px 0px 10px" }}
              />
              <Input
                value={power}
                onChange={(e) => setPower(e.target.value)}
                style={{ borderRadius: "0px 10px 10px 0px" }}
              />
              <div>
                <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
                  [dBm]
                </sub>
              </div>
            </div>
          </Col>
        </Row>
        <Row className="gutter-row" span={8}>
          <Col>
            <Typography.Title level={5}>Combiner Loss in BTS</Typography.Title>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Button
                type="primary"
                shape="round"
                onClick={handleCombinerCall}
                icon={<ArrowRightOutlined />}
                style={{ borderRadius: "10px 0px 0px 10px" }}
              />
              <Input
                value={combiner}
                onChange={(e) => setCombiner(e.target.value)}
                style={{ borderRadius: "0px 10px 10px 0px" }}
              />
              <div>
                <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
                  [dBm]
                </sub>
              </div>
            </div>
          </Col>
        </Row>
        <Row className="gutter-row" span={8}>
          <Col>
            <Typography.Title level={5}>Feeder Loss in BTS</Typography.Title>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Button
                type="primary"
                shape="round"
                onClick={handleFeederCall}
                icon={<ArrowRightOutlined />}
                style={{ borderRadius: "10px 0px 0px 10px" }}
              />
              <Input
                value={feeder}
                onChange={(e) => setFeeder(e.target.value)}
                style={{ borderRadius: "0px 10px 10px 0px" }}
              />
              <div>
                <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
                  [dBm]
                </sub>
              </div>
            </div>
          </Col>
        </Row>
        <Row className="gutter-row" span={8}>
          <Col>
            <Typography.Title level={5}>Antenna Gain in BTS</Typography.Title>
            <div style={{ display: "flex", alignItems: "center" }}>
              <Button
                type="primary"
                shape="round"
                onClick={handleAntennaCall}
                icon={<ArrowRightOutlined />}
                style={{ borderRadius: "10px 0px 0px 10px" }}
              />
              <Input
                value={antenna}
                onChange={(e) => setAtntenna(e.target.value)}
                style={{ borderRadius: "0px 10px 10px 0px" }}
              />
              <div>
                <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
                  [dBm]
                </sub>
              </div>
            </div>
          </Col>
        </Row>
      </div>
    );
  };
  
  export default EIRP;
