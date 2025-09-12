import { AuthContext } from "@/contexts/auth-context";
import { Auth } from "@/types/auth";
import { useContext } from "react";

export const useAuth = (): Auth => {
    const context = useContext(AuthContext);
    if (context === undefined) throw new Error("AuthContext is undefined");
    return context;
};
