import axios from "axios";

const API_URL = "http://localhost:9090/containers";

export const getPingResults = async () => {
    const response = await axios.get(API_URL);
    return response.data;
};
