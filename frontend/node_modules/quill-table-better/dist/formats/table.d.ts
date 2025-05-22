import type { BlockBlot, ContainerBlot, LinkedList } from 'parchment';
import type { Props } from '../types';
import TableHeader from './header';
import { ListContainer } from './list';
declare const Block: typeof BlockBlot;
declare const Container: typeof ContainerBlot;
declare class TableCellBlock extends Block {
    static blotName: string;
    static className: string;
    static tagName: string;
    next: this | null;
    parent: TableCell;
    static create(value: string): HTMLElement;
    format(name: string, value: string | Props): import("parchment").Blot;
    formats(): {
        [key: string]: any;
    };
    getCellFormats(parent: TableCell): Props;
    wrapTableCell(parent: TableCell): void;
}
declare class TableCell extends Container {
    static blotName: string;
    static tagName: string;
    children: LinkedList<TableCellBlock | TableHeader | ListContainer>;
    next: this | null;
    parent: TableRow;
    prev: this | null;
    checkMerge(): boolean;
    static create(value: Props): HTMLElement;
    static formats(domNode: Element): Props;
    formats(): {
        [x: number]: any;
    };
    static getEmptyRowspan(domNode: Element): number;
    static hasColgroup(domNode: Element): boolean;
    html(): string;
    row(): TableRow;
    rowOffset(): number;
    setChildrenId(cellId: string): void;
    table(): TableContainer;
    optimize(context?: unknown): void;
}
declare class TableRow extends Container {
    static blotName: string;
    static tagName: string;
    children: LinkedList<TableCell>;
    next: this | null;
    parent: TableBody;
    prev: this | null;
    checkMerge(): boolean;
    rowOffset(): number;
}
declare class TableBody extends Container {
    static blotName: string;
    static tagName: string;
    children: LinkedList<TableRow>;
    next: this | null;
    parent: TableContainer;
}
declare class TableTemporary extends Block {
    static blotName: string;
    static className: string;
    static tagName: string;
    next: this | null;
    static create(value: Props): HTMLElement;
    static formats(domNode: Element): Props;
    formats(): {
        [x: number]: any;
    };
    optimize(...args: unknown[]): void;
}
declare class TableCol extends Block {
    static blotName: string;
    static tagName: string;
    next: this | null;
    static create(value: Props): HTMLElement;
    static formats(domNode: Element): Props;
    formats(): {
        [x: number]: any;
    };
    html(): string;
}
declare class TableColgroup extends Container {
    static blotName: string;
    static tagName: string;
    children: LinkedList<TableCol>;
    next: this | null;
}
declare class TableContainer extends Container {
    static blotName: string;
    static defaultClassName: string;
    static tagName: string;
    children: LinkedList<TableBody | TableTemporary | TableColgroup>;
    colgroup(): TableColgroup;
    deleteColumn(changeTds: [Element, number][], delTds: Element[], deleteTable: () => void, cols?: Element[]): void;
    deleteRow(rows: TableRow[], deleteTable: () => void): void;
    deleteTable(): void;
    findChild(blotName: string): TableBody | TableTemporary | TableColgroup;
    getCopyTable(): string;
    getCorrectRow(prev: TableRow, maxColumns: number): TableRow;
    getInsertRow(prev: TableRow, ref: TableRow | null, offset: number): TableRow;
    getMaxColumns(children: LinkedList<TableCell>): number;
    insertColumn(position: number, isLast: boolean, w: number, offset: number): void;
    insertCol(colgroup: TableColgroup, ref: TableCol | null): void;
    insertColumnCell(row: TableRow | null, id: string, ref: TableCell | null): TableCell;
    insertRow(index: number, offset: number): void;
    insertTableCell(colspan: number, formats: Props, row: TableRow): void;
    optimize(context: unknown): void;
    setCellColspan(cell: TableCell, offset: number): void;
    setCellRowspan(parentElement: Element): void;
    private setClassName;
    private setColumnCells;
    tbody(): TableBody;
    temporary(): TableTemporary;
}
declare function cellId(): string;
declare function tableId(): string;
export { cellId, TableCellBlock, TableCell, TableRow, TableBody, TableTemporary, TableContainer, tableId, TableCol, TableColgroup };
