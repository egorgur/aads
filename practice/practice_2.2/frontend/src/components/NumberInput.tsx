
interface NumberInputProps {
    id: string,
    label: string,
    value: number,
    incrementFunction: () => ({}),
    decrementFunction: () => ({}),
}


export default function NumberInput(props: NumberInputProps) {
    return (
        <div className="
        w-full
        flex flex-row
        pt-5
        ">
            <span>{props.label}:</span>
            <button className="shrink-0 w-5 ml-3" onClick={() => { props.decrementFunction() }}>-</button>
            <span className="mx-2">{props.value}</span>
            <button className="shrink-0 w-5" onClick={() => { console.log(props.value);props.incrementFunction() }} >+</button>

        </div>
    )
}