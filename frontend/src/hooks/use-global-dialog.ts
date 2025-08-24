import { GlobalDialogContext } from "@/contexts/global-dialog-context";
import { GlobalDialog } from "@/types/global-dialog";
import { useContext } from "react";

export const useGlobalDialog = (): GlobalDialog => {
    const context = useContext(GlobalDialogContext);
    if (context === undefined) {
        throw new Error("Global dialog context is undefined");
    }

    return context;
};
