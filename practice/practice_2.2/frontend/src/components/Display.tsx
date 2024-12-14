import { useMazeStore } from "../store/MazeStore.ts"
import { useMenuStore, MenuType } from "../store/MenuStore.ts"
import Maze from "./Maze"

export default function Display() {
    const chosenMenu: MenuType = useMenuStore((state: any) => state.menuType)

    const maze = useMazeStore((state: any) => state.maze)
    const mazeGenSteps = useMazeStore((state: any) => state.mazeGenSteps)
    const mazeSolve = useMazeStore((state: any) => state.mazeSolve)
    const mazeSolveSteps = useMazeStore((state: any) => state.mazeSolveSteps)

    switch (chosenMenu) {
        case MenuType.Generate:
            return (
                <aside
                    className="
                        grow
                        border-l
                        flex flex-col
                        space-y-3
                        justify-around
                        items-center
                        ">
                    <Maze data={maze} dynData={undefined} dynamic={false}/>
                    <Maze data={mazeGenSteps} dynData={mazeGenSteps} dynamic={true}/>

                </aside>
            )
        case MenuType.Solve:
            return (
                <aside
                    className="
                        grow
                        border-l
                        flex flex-col
                        space-y-3
                        justify-around
                        items-center
                        ">
                    <Maze data={mazeSolve} dynData={undefined} dynamic={false}/>
                    <Maze data={mazeSolveSteps} dynData={mazeSolveSteps} dynamic={true}/>

                </aside>
            )
    }


}