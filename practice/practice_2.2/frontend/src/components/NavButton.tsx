import { useMenuStore, MenuType, MenuStore } from "../store/MenuStore.ts"

interface ButtonProps {
    text: string
    menuType: MenuType
}


export default function NavButton(props: ButtonProps) {

    const chosenMenu: MenuType = useMenuStore((state: any) => state.menuType)
    const setMenuType: MenuStore["setMenuType"] = useMenuStore((state: any) => state.setMenuType)


    if (props.menuType == chosenMenu) {
        return (
            <button className="
            text-zinc-50
            underline
            hover:text-zinc-50
            transition-all transform duration-200
            ">
                {props.text}
            </button>
        )
    }
    return (
        <button className="
        text-zinc-300
        hover:text-zinc-50
        hover:underline
        transition-all transform duration-200
        "
        onClick={() => {setMenuType(props.menuType)}}
        >
            {props.text}
        </button>
    )
}