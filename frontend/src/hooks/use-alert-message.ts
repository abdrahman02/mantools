"use client";
import { AlertMessageContext } from "@/contexts/alert-message-context";
import { AlertMessage } from "@/types/alert-message";
import { useContext } from "react";

export function useAlertMessage(): AlertMessage {
    const context = useContext(AlertMessageContext);
    if (context === undefined) {
        throw new Error("AlertMessageContext is undefined");
    }
    return context;
}
