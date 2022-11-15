import axios, { AxiosError, AxiosResponse } from "axios";
import { Moment } from "moment";
import { Matches } from "../models/matches.interface";

const instance = axios.create({
    // baseURL: process.env.REACT_APP_MATCH_DISPLAY_URL,
    baseURL: "https://match-display.herokuapp.com",
    timeout: 15000,
    headers: {
        "Content-type": "application/json"
      }
});

const responseBody = (response: AxiosResponse) => response.data;

const requests = {
    get: (url: string, params: { date: string }) => instance.get(url, {params})
                                                                .then(responseBody)
                                                                .catch((err: Error | AxiosError) => {
                                                                    console.log("error in api call");
                                                                    throw err
                                                                }),
};

export const Match = {
    getMatches: (date: Moment): Promise<Matches[]> => {
        const dateStr = date.format("YYYY-MM-DD");
        // console.log("date string in getMatches func", dateStr);
        const params = {
            date: dateStr
        }
        return requests.get('/matches', params)
    },
};
