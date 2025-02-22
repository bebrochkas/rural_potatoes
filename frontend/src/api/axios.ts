import axios from "axios";

const baseURL = "http://localhost:3000/api";

const getAuthorizationHeader = () =>
    `Bearer ${localStorage.getItem("jwtToken")}`;

export const axiosInst = axios.create({
    baseURL,
    headers: { Authorization: getAuthorizationHeader() },
});
