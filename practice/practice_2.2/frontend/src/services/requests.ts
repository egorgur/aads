/**
 *  Requests to generate and solve maze
 */
import { useMazeStore } from "../store/MazeStore";
import axios from "axios";

const URL = "http://localhost:8000"

interface GenerationParameters {
    rows: number,
    cols: number,
    console: boolean,
    text: boolean,
    image: boolean,
    gif: boolean,
}

interface SolveParameters {
    start_x: number,
    start_y: number,
    end_x: number,
    end_y: number,
}

export function requestGeneration(data: GenerationParameters) {
    axios({
        method: "post",
        url: URL + "/generate/",
        data: data
    })
        .then(function (response) {
            console.log(response)
            useMazeStore.setState({ maze: response.data.maze })
            useMazeStore.setState({ mazeGenSteps: response.data.generation_steps })
        })
        .catch(function (error) {
            console.log(error);
        });
}

export function requestSolve(data: SolveParameters) {
    axios({
        method: "post",
        url: URL + "/solve/",
        data: data
    })
        .then(function (response) {
            console.log(response);
            useMazeStore.setState({ mazeSolve: response.data.solve })
            useMazeStore.setState({ mazeSolveSteps: response.data.solving_steps })
        })
        .catch(function (error) {
            console.log(error);
        });
}
