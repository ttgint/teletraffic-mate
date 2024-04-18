import React, { useState } from "react";
import {
  Layout,
  theme,
  Input,
  Typography,
  Col,
  Row,
  Radio,
  Button,
} from "antd";
import { ArrowRightOutlined } from "@ant-design/icons";
import axios from "axios";
const { Content } = Layout;


/**
 * Main Component
 * <Traffic Content>
 *    <ModelBasedCalculations /> ==> Calculations that are specific to Models (ErlangC, ErlangB, Poisson)
 *    <CalculationFields /> ==> The rest of the calculations
 * <Traffic Content/>
 */
function TrafficContent() {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const ContentStyle = {
    background: colorBgContainer,
    minHeight: 350,
    padding: 24,
    borderRadius: borderRadiusLG,
  };

  return (
    <Content>
      <div style={ContentStyle}>
        <ModelBasedCalculations />
        <CalculationFields />
      </div>
    </Content>
  );
}

const ModelBasedCalculations = () => {
  const ModelsSelectionStyle = {
    padding: "0",
    margin: "0 0 -10px 0",
    textAlign: "center",
  };
  const ModelBasedSectionStyles = {
    background: "#eee",
    padding: "30px 10px"
  }
  const [selectedValue, setSelectedValue] = useState("erlangB");
  const [num_of_channels, setNumOfChannels] = useState("");
  const [gos, setGos] = useState("");
  const [traffic, setTraffic] = useState("");

  const handleRadioChange = (e) => {
    setSelectedValue(e.target.value);
  };

  const handleApiCall = async (
    erlangType,
    endpoint,
    param_1,
    param_2,
    setStateFunction
  ) => {
    try {
      const res = await axios.post(
        `http://localhost:5174/traffic/${erlangType}/${endpoint}/${param_1}/${param_2}`
      );
      setStateFunction(res.data);
    } catch (error) {
      console.error(error);
    }
  };

  const handleTraffic = async () => handleApiCall(selectedValue, "traffic", gos, num_of_channels, setTraffic);
  const handleGoS = async () => handleApiCall(selectedValue, "gos", traffic, num_of_channels, setGos);
  const handleNumOfChannels = async () => handleApiCall(selectedValue, "num-of-channels", gos, traffic, setNumOfChannels);

  return (
    <>
      <div style={ModelBasedSectionStyles}>
        <Row gutter={[24, 24]}>
          <Col className="gutter-row" span={8}>
            <div style={ModelsSelectionStyle}>
        
              <Radio.Group onChange={handleRadioChange} value={selectedValue}>
                <Radio value="erlangB">ErlangB</Radio>
              </Radio.Group>
            </div>
          </Col>
          <Col className="gutter-row" span={8}>
            <div style={ModelsSelectionStyle}>
              <Radio.Group onChange={handleRadioChange} value={selectedValue}>
                <Radio value="erlangC">ErlangC</Radio>
              </Radio.Group>
            </div>
          </Col>
          <Col className="gutter-row" span={8}>
            <div style={ModelsSelectionStyle}>
              <Radio.Group onChange={handleRadioChange} value={selectedValue}>
                <Radio value="poisson">Poisson</Radio>
              </Radio.Group>
            </div>
          </Col>
        </Row>
        <CalculationFieldRow
          MeasurementName={["Number of Channels", "Traffic", "GoS"]}
          HandleFunctions={[handleNumOfChannels, handleTraffic, handleGoS]}
          States={[num_of_channels, traffic, gos]}
          SetStates={[setNumOfChannels, setTraffic, setGos]}
          units={["#", "Erl", "%"]}
          />
      </div>
    </>
  );
};

function CalculationFields() {
  const [num_of_subs, setNumOfSubs] = useState("");
  const [traffic_per_subs, setTraffPerSub] = useState("");
  const [traffic_2, setTraffic2] = useState("");

  const [traffic_3, setTraffic3] = useState("");
  const [bhca, setBHCA] = useState("");
  const [mht, setMHT] = useState("");

  const [subs_vlr, setSubsVLR] = useState("");
  const [traffic_per_subs_2, setTraffPerSub2] = useState("");
  const [switch_traffic, setSwitchTraffic] = useState("");

  const handleApiCall = async (
    endpoint,
    param_1,
    param_2,
    setStateFunction
  ) => {
    try {
      const res = await axios.post(
        `http://localhost:5174/traffic/${endpoint}/${param_1}/${param_2}`
      );
      setStateFunction(res.data);
    } catch (error) {
      console.error(error);
    }
  };

  /*==========================================================*/
  /*========================Second Row========================*/
  /*==========================================================*/

  const handleNumOfSubs = async () => handleApiCall("num-of-subs", traffic_per_subs, traffic_2, setNumOfSubs);
  const handleTrafficPerSubs = async () => handleApiCall("traffic-per-subs", num_of_subs, traffic_2, setTraffPerSub);
  const handleTraffic2 = async () => handleApiCall("traffic-2", num_of_subs, traffic_per_subs, setTraffic2);

  /*==========================================================*/
  /*========================Third Row========================*/
  /*==========================================================*/

  const handleTraffic3 = async () => handleApiCall("traffic-3", bhca, mht, setTraffic3);
  const handleBHCA = async () => handleApiCall("bhca", traffic_3, mht, setBHCA);
  const handleMHT = async () => handleApiCall("mht", bhca, traffic_3, setMHT);

  /*==========================================================*/
  /*========================Fourth Row========================*/
  /*==========================================================*/

  const handleSubsVLR = async () => handleApiCall("subs-vlr", traffic_per_subs_2, switch_traffic, setSubsVLR);
  const handleTrafficSub = async () => handleApiCall("traffic-per-subs-2", subs_vlr, switch_traffic, setTraffPerSub2);
  const handleSwitchTraffic = async () => handleApiCall("switch-traffic", traffic_per_subs_2, subs_vlr, setSwitchTraffic);

  return (
    <>
      <CalculationFieldRow
        MeasurementName={["Number of Subs", "Traffic/Subs", "Traffic"]}
        HandleFunctions={[
          handleNumOfSubs,
          handleTrafficPerSubs,
          handleTraffic2,
        ]}
        States={[num_of_subs, traffic_per_subs, traffic_2]}
        SetStates={[setNumOfSubs, setTraffPerSub, setTraffic2]}
        units={["#", "mErl/sub", "Erl"]}
      />
      <CalculationFieldRow
        MeasurementName={["Traffic", "BHCA", "MHT"]}
        HandleFunctions={[handleTraffic3, handleBHCA, handleMHT]}
        States={[traffic_3, bhca, mht]}
        SetStates={[setTraffic3, setBHCA, setMHT]}
        units={["Erl", "#", "Sec"]}
      />
      <CalculationFieldRow
        MeasurementName={["Subs. VLR", "Traffic/Sub.", "Switch Traffic"]}
        HandleFunctions={[handleSubsVLR, handleTrafficSub, handleSwitchTraffic]}
        States={[subs_vlr, traffic_per_subs_2, switch_traffic]}
        SetStates={[setSubsVLR, setTraffPerSub2, setSwitchTraffic]}
        units={["#", "mErl/sub", "Erl"]}
      />
    </>
  );
}

function CalculationFieldRow({
  MeasurementName,
  HandleFunctions,
  States,
  SetStates,
  units,
}) {
  return (
    <Row gutter={[24, 24]}>
      <CalculationField
        MeasurementName={MeasurementName[0]}
        handleFunction={HandleFunctions[0]}
        state={States[0]}
        setStateFunction={SetStates[0]}
        units={units[0]}
      />
      <CalculationField
        MeasurementName={MeasurementName[1]}
        handleFunction={HandleFunctions[1]}
        state={States[1]}
        setStateFunction={SetStates[1]}
        units={units[1]}
      />
      <CalculationField
        MeasurementName={MeasurementName[2]}
        handleFunction={HandleFunctions[2]}
        state={States[2]}
        setStateFunction={SetStates[2]}
        units={units[2]}
      />
    </Row>
  );
}

function CalculationField({
  MeasurementName,
  handleFunction,
  state,
  setStateFunction,
  units,
}) {
  return (
    <Col className="gutter-row" span={8}>
      <Typography.Title level={5}>{MeasurementName}</Typography.Title>
      <div style={{ display: "flex", alignItems: "center" }}>
        {" "}
        <Button
          type="primary"
          shape="round"
          icon={<ArrowRightOutlined />}
          onClick={handleFunction}
          style={{ borderRadius: "10px 0px 0px 10px" }}
        />
        <Input
          placeholder={MeasurementName}
          value={state}
          onChange={(e) => setStateFunction(e.target.value)}
          style={{ borderRadius: "0px 10px 10px 0px" }}
        />
        <div>
          {units && (
            <sub style={{ marginLeft: "5px", fontStyle: "italic" }}>
              [{units}]
            </sub>
          )}
        </div>
      </div>
    </Col>
  );
}

export default TrafficContent;
