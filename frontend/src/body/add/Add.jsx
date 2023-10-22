import "./Add.css";
import React from "react";
import { Button, Form, Input } from "antd";
import { AddProblem } from "../../../wailsjs/go/main/App.js";

const formLabel = (label) => {
  return <div className="form-label">{label}</div>;
};

const Add = ({ setSelectedMenuItem }) => {
  const [form] = Form.useForm();
  const onFinish = () => {
    AddProblem(
      form.getFieldValue("UserID"),
      form.getFieldValue("Date"),
      form.getFieldValue("ProblemUrl1"),
      form.getFieldValue("ProblemUrl2"),
      form.getFieldValue("ProblemUrl3"),
    );
    form.setFieldValue("ProblemUrl1", "");
    form.setFieldValue("ProblemUrl2", "");
    form.setFieldValue("ProblemUrl3", "");
  };

  const onFinishFailed = () => {};
  const onProblemUrlChange = (changedUrl) => {};

  return (
    <>
      <Form form={form} onFinish={onFinish} onFinishFailed={onFinishFailed}>
        <Form.Item
          label={formLabel("사용자 이름")}
          name="UserID"
          rules={[
            {
              required: true,
              message: "사용자 ID를 입력하세요.",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label={formLabel("날짜")}
          name="Date"
          rules={[
            {
              required: true,
              message: "날짜를 입력하세요.",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label={formLabel("문제(1) URL")}
          name="ProblemUrl1"
          rules={[
            {
              required: true,
              message: "문제 URL을 입력하세요.",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label={formLabel("문제(2) URL")}
          name="ProblemUrl2"
          rules={[
            {
              required: true,
              message: "문제 URL을 입력하세요.",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label={formLabel("문제(3) URL")}
          name="ProblemUrl3"
          rules={[
            {
              required: true,
              message: "문제 URL을 입력하세요.",
            },
          ]}
        >
          <Input />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
      <button
        className="goBackButton"
        onClick={() => {
          setSelectedMenuItem(null);
        }}
      >
        Go Back
      </button>
    </>
  );
};

export default Add;
