"use client";
import { AlertMessage } from "@/types/alert-message";
import { createContext, ReactNode, useState } from "react";

export const AlertMessageContext = createContext<AlertMessage | undefined>(
    undefined,
);

export const AlertMessageProvider = ({ children }: { children: ReactNode }) => {
    const [status, setStatus] = useState<number>(0);
    const [message, setMessage] = useState<string>("");
    const createAlert = (status: number, message: string) => {
        setStatus(status);
        setMessage(message);
    };
    const dismissAlert = () => {
        setStatus(0);
        setMessage("");
    };
    return (
        <AlertMessageContext.Provider
            value={{ status, message, createAlert, dismissAlert }}
        >
            {children}
        </AlertMessageContext.Provider>
    );
};
