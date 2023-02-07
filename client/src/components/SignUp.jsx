import React, { useState } from "react";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/esm/Button";
import { useMutation } from "react-query";
import { API } from "../config/api";
import { Alert } from "bootstrap";
import { useNavigate } from "react-router-dom";

export default function SignUp(props) {
  let navigate = useNavigate();
  const [message, setMessage] = useState(null);
  const [userSignUp, setUserSignUp] = useState({
    fullname: "",
    username: "",
    email: "",
    password: "",
    phone: "",
    listAs: "",
    gender: "",
    address: "",
  });

  const redirectSignin = (e) => {
    props.onHide();
    props.openSignin();
  };

  const handleOnChange = (e) => {
    setUserSignUp({
      ...userSignUp,
      [e.target.name]: e.target.value,
    });
  };

  const handleOnSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();
      const response = await API.post("/register", userSignUp);
      console.log("ini register: ", response);

      const alert = (
        <Alert variant="success" className="py-1">
          Register Success
        </Alert>
      );
      setMessage(alert);
      console.log("register : ", response);

      // Handling response here
    } catch (error) {
      const alert = (
        <Alert variant="danger" className="py-1">
          Failed
        </Alert>
      );
      setMessage(alert);
      console.log(error);
    }
  });

  return (
    <Modal {...props} aria-labelledby="contained-modal-title-vcenter" centered>
      <Modal.Body
        style={{
          height: "600px",
          overflow: "hidden",
        }}
      >
        <h1 className="fw-bold text-center my-5">Sign Up</h1>
        
        <Form
          onSubmit={(e) => handleOnSubmit.mutate(e)}
          style={{
            height: "420px",
            overflow: "auto",
          }}
        >
          <Form.Group className="mb-3" controlId="fullname">
            <Form.Label className="fw-bold">Fullname</Form.Label>
            <Form.Control
              type="text"
              placeholder=""
              name="fullname"
              value={userSignUp.fullname}
              onChange={handleOnChange}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="username">
            <Form.Label className="fw-bold">Username</Form.Label>
            <Form.Control
              type="text"
              placeholder=""
              name="username"
              value={userSignUp.username}
              onChange={handleOnChange}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="email">
            <Form.Label className="fw-bold">Email</Form.Label>
            <Form.Control
              type="email"
              placeholder="Enter email"
              name="email"
              value={userSignUp.email}
              onChange={handleOnChange}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="phone">
            <Form.Label className="fw-bold">phone</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter phone number"
              name="phone"
              value={userSignUp.phone}
              onChange={handleOnChange}
            />
          </Form.Group>
          <Form.Group className="mb-3" controlId="password">
            <Form.Label className="fw-bold">Password</Form.Label>
            <Form.Control
              type="password"
              placeholder="Password"
              name="password"
              value={userSignUp.password}
              onChange={handleOnChange}
            />
          </Form.Group>
          <Form.Group className="mb-3 " controlId="listAs">
            <Form.Label className="fw-bold">List As</Form.Label>
            <Form.Select
              name="listAs"
              aria-label="Default select example"
              value={userSignUp.listAs}
              onChange={handleOnChange}
            >
              <option>~Select~</option>
              <option value="Tenant">Tenant</option>
              <option value="Owner">Owner</option>
            </Form.Select>
          </Form.Group>
          <Form.Group className="mb-3" controlId="gender">
            <Form.Label className="fw-bold">Gender</Form.Label>
            <Form.Select
              name="gender"
              aria-label="Default select example"
              value={userSignUp.gender}
              onChange={handleOnChange}
            >
              <option>~Select~</option>
              <option value="Male">Male</option>
              <option value="Female">Female</option>
            </Form.Select>
          </Form.Group>
          <Form.Group className="mb-3 " controlId="alamat">
            <Form.Label className="fw-bold">Address</Form.Label>
            <Form.Control
              className="rs"
              as="textarea"
              name="address"
              style={{ height: "100px" }}
              value={userSignUp.address}
              onChange={handleOnChange}
            />
          </Form.Group>

          <Button
            onClick={(e) => redirectSignin(e)}
            className="w-100"
            variant="primary"
            type="submit"
          >
            Submit
          </Button>
        </Form>
      </Modal.Body>
    </Modal>
  );
}
