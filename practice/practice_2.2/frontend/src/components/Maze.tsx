
import { useState, useEffect } from "react"
interface MazeProps {
    data: Array<Array<number>>
    dynData: Array<Array<Array<number>>> | undefined
    dynamic: boolean
}

interface BlockProps {
    solid: boolean
    path: boolean
}

interface LineProps {
    data: Array<number>
}

function Block(props: BlockProps) {
    let style = "w-3 h-3"
    if (!props.solid) {
        style = style + " bg-zinc-50"
    }
    if (props.path) {
        style = style + " bg-red-900"
    }
    return (
        <div
            className={style}>

        </div>
    )
}

function Row(props: LineProps) {

    const elements = [];

    for (let element of props.data) {
        if (element == 1) {
            elements.push(Block({ solid: true, path: false }))
        
        } else if (element == -100) {
            elements.push(Block({ solid: false, path: true }))
        }
        else {
            elements.push(Block({ solid: false, path: false }))
        }

    }

    return (
        <div className="flex flex-row">{elements}</div>
    )

}

export default function Maze(props: MazeProps) {

    let rows = []

    if (!props.dynamic) {
        for (let row of props.data) {
            rows.push(Row({ data: row }))
        }

        return (
            <div className="
           
            
            ">
                {rows}
            </div>
        )
    } else {
        const len = props.data.length
        console.log(len);
        if (len == 0) {
            return(
                <div className="
           
            
            ">
                
            </div>
            )
        }
        const [intervalCount, setIntervalCount] = useState(0);

        useEffect(() => {

            // wait 5 s before cause a re-render
            setTimeout(() => {
                if (intervalCount == len - 1) {
                    setIntervalCount(0);
                } else {
                    setIntervalCount(count => count + 1);
                }
            }, 1000);

        }, [intervalCount]);

        if (props.dynData !== undefined) {
            const mazeData: Array<Array<number>> = props.dynData[intervalCount]
            console.log(intervalCount,mazeData);
            
            

            for (let row of mazeData) {
                rows.push(Row({ data: row }))
            }

            return (
                <div className="">
                    {rows}
                </div>
            )
        }
    }
}

