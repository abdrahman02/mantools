"use client";
import { pageView } from "@/lib/gtag";
import { usePathname } from "next/navigation";
import { useEffect } from "react";

export default function AnalyticsTracker() {
    const pathname = usePathname();

    useEffect(() => {
        if (pathname) pageView(pathname);
    }, [pathname]);

    return null;
}
