"use client";

import { SidebarProvider } from "@/components/ui/sidebar";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

export default function SidebarWrapper({
    children,
}: {
    children: React.ReactNode;
}) {
    const pathname = usePathname();
    const [open, setOpen] = useState(pathname !== "/");

    useEffect(() => {
        if (pathname === "/") {
            setOpen(false);
        }
    }, [pathname]);

    return (
        <SidebarProvider open={open} onOpenChange={setOpen}>
            {children}
        </SidebarProvider>
    );
}
