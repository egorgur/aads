import NumberInput from "./NumberInput.tsx"
import CheckBox from "./CheckBox.tsx"
import { requestGeneration, requestSolve } from "../services/requests.ts"
import { useMenuStore, MenuType, MenuStore } from "../store/MenuStore.ts"

export default function Menu() {
    const chosenMenu: MenuType = useMenuStore((state: any) => state.menuType)

    const rows: number = useMenuStore((state: any) => state.rows)
    const increaseRows = useMenuStore((state: any) => state.increaseRows)
    const decreaseRows = useMenuStore((state: any) => state.decreaseRows)

    const cols: number = useMenuStore((state: any) => state.cols)
    const increaseCols = useMenuStore((state: any) => state.increaseCols)
    const decreaseCols = useMenuStore((state: any) => state.decreaseCols)

    const consoleFlag: boolean = useMenuStore((state: any) => state.console)
    const setConsole = useMenuStore((state: any) => state.setConsole)

    const textFlag: boolean = useMenuStore((state: any) => state.text)
    const setText = useMenuStore((state: any) => state.setText)

    const imageFlag: boolean = useMenuStore((state: any) => state.image)
    const setImage = useMenuStore((state: any) => state.setImage)

    const gifFlag: boolean = useMenuStore((state: any) => state.gif)
    const setGif = useMenuStore((state: any) => state.setGif)


    const startX: number = useMenuStore((state: any) => state.startX)
    const increaseStartX = useMenuStore((state: any) => state.increaseStartX)
    const decreaseStartX = useMenuStore((state: any) => state.decreaseStartX)

    const startY: number = useMenuStore((state: any) => state.startY)
    const increaseStartY = useMenuStore((state: any) => state.increaseStartY)
    const decreaseStartY = useMenuStore((state: any) => state.decreaseStartY)

    const endX: number = useMenuStore((state: any) => state.endX)
    const increaseEndX = useMenuStore((state: any) => state.increaseEndX)
    const decreaseEndX = useMenuStore((state: any) => state.decreaseEndX)

    const endY: number = useMenuStore((state: any) => state.endY)
    const increaseEndY = useMenuStore((state: any) => state.increaseEndY)
    const decreaseEndY = useMenuStore((state: any) => state.decreaseEndY)


    switch (chosenMenu) {
        case MenuType.Generate:
            return (
                <menu className="
                w-0 grow min-w-40 max-w-56
                flex flex-col px-5
                ">

                    <NumberInput
                        label="Rows"
                        id="rows"
                        value={rows}
                        incrementFunction={increaseRows}
                        decrementFunction={decreaseRows}
                    />

                    <NumberInput
                        label="Cols"
                        id="cols"
                        value={cols}
                        incrementFunction={increaseCols}
                        decrementFunction={decreaseCols}
                    />

                    <CheckBox
                        id="console"
                        label="Console"
                        value={consoleFlag}
                        setFunction={setConsole}
                    />

                    <CheckBox
                        id="text"
                        label="Text"
                        value={textFlag}
                        setFunction={setText}
                    />
                    <CheckBox
                        id="image"
                        label="Image"
                        value={imageFlag}
                        setFunction={setImage}
                    />
                    <CheckBox
                        id="gif"
                        label="Gif"
                        value={gifFlag}
                        setFunction={setGif}
                    />

                    <button className="
                    mt-5
                    text-xl
                    flex flex-row
                    align-left
                    text-zinc-200
                    hover:text-zinc-50
                    hover:underline
                    transition-all transform duration-200
                    "
                        onClick={() => {
                            requestGeneration({
                                rows: rows,
                                cols: cols,
                                console: consoleFlag,
                                text: textFlag,
                                image: imageFlag,
                                gif: gifFlag,
                            })
                        }}
                    >Generate</button>

                </menu>
            )
        case MenuType.Solve:
            return (
                <menu className="
                w-0 grow min-w-40 max-w-56
                flex flex-col px-5

                ">
                    <NumberInput
                        label="Start X"
                        id="startX"
                        value={startX}
                        incrementFunction={increaseStartX}
                        decrementFunction={decreaseStartX}
                    />

                    <NumberInput
                        label="Start Y"
                        id="startY"
                        value={startY}
                        incrementFunction={increaseStartY}
                        decrementFunction={decreaseStartY}
                    />

                    <NumberInput
                        label="End X"
                        id="endX"
                        value={endX}
                        incrementFunction={increaseEndX}
                        decrementFunction={decreaseEndX}
                    />

                    <NumberInput
                        label="End Y"
                        id="endY"
                        value={endY}
                        incrementFunction={increaseEndY}
                        decrementFunction={decreaseEndY}
                    />

                    <button className="
                    mt-5
                    text-xl
                    flex flex-row
                    align-left
                    text-zinc-200
                    hover:text-zinc-50
                    hover:underline
                    transition-all transform duration-200
                    "
                    onClick={() => {
                        requestSolve({
                            start_x:startX,
                            start_y:startY,
                            end_x: endX,
                            end_y: endY,
                        })
                    }}
                    >Solve</button>

                </menu>
            )
    }
}