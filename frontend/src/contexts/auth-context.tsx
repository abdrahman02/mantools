"use client";
import { Auth } from "@/types/auth";
import { createContext, ReactNode, useState } from "react";

export const AuthContext = createContext<Auth | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [access, setAccess] = useState<string | null>(null);

    return (
        <AuthContext.Provider value={{ access, setAccess }}>
            {children}
        </AuthContext.Provider>
    );
}
