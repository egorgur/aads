
interface CheckBoxProps {
    id: string,
    label: string,
    value: boolean,
    setFunction: (argument: boolean) => ({}),
}


export default function CheckBox(props: CheckBoxProps) {
    return (
        <div className="
        w-full
        flex flex-row
        pt-5
        ">
            <span>{props.label}:</span>
            <span className="grow"></span>
            <input type="checkbox" checked={props.value}
            className="
            ml-3 mr-3
            "
            onChange={() => {
                if (props.value) {
                    props.setFunction(false)
                } else {
                    props.setFunction(true)
                }
                console.log(props.value);
            }}
            />

        </div>
    )
}