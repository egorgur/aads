import { create } from "zustand"



export const useMazeStore = create((set) => ({
    maze: [],
    setMazeData: (data: any) => set({ maze: data }),
    mazeGenSteps: [],
    setGenSteps: (data: any) => set({ mazeGenSteps: data }),
    mazeSolve: [],
    setMazeSolve: (data: any) => set({ mazeSolve: data }),
    mazeSolveSteps: [],
    setSolveSteps: (data: any) => set({ mazeSolveSteps: data }),
}))