export type GlobalDialogType = "loading" | "error";

export interface GlobalDialog {
    isLoading: boolean;
    showDialog: (type?: GlobalDialogType) => void;
    hideDialog: () => void;
}
