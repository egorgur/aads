import { create } from "zustand"

export enum MenuType {
    Generate,
    Solve
}

export interface MenuStore {
    menuType: MenuType,
    setMenuType: (type: MenuType) => {}
    rows: number,
    cols: number
    console: boolean,
    text: boolean,
    image: boolean,
    gif: boolean,
    startX: number,
    startY: number,
    endX: number,
    endY: number,
}

export const useMenuStore = create((set) => ({
    menuType: MenuType.Generate,
    setMenuType: (type: MenuType) => set({ menuType: type }),
    rows: 0,
    increaseRows: () => set((state: any) => ({ rows: state.rows + 1 })),
    decreaseRows: () => set((state: any) => ({ rows: state.rows - 1 })),
    cols: 0,
    increaseCols: () => set((state: any) => ({ cols: state.cols + 1 })),
    decreaseCols: () => set((state: any) => ({ cols: state.cols - 1 })),
    console: false,
    setConsole: (flag: boolean) => set({ console: flag }),
    text: false,
    setText: (flag: boolean) => set({ text: flag }),
    image: false,
    setImage: (flag: boolean) => set({ image: flag }),
    gif: false,
    setGif: (flag: boolean) => set({ gif: flag }),
    startX: 0,
    increaseStartX: () => set((state: any) => ({ startX: state.startX + 1 })),
    decreaseStartX: () => set((state: any) => ({ startX: state.startX - 1 })),
    startY: 0,
    increaseStartY: () => set((state: any) => ({ startY: state.startY + 1 })),
    decreaseStartY: () => set((state: any) => ({ startY: state.startY - 1 })),
    endX: 0,
    increaseEndX: () => set((state: any) => ({ endX: state.endX + 1 })),
    decreaseEndX: () => set((state: any) => ({ endX: state.endX - 1 })),
    endY: 0,
    increaseEndY: () => set((state: any) => ({ endY: state.endY + 1 })),
    decreaseEndY: () => set((state: any) => ({ endY: state.endY - 1 })),
}))