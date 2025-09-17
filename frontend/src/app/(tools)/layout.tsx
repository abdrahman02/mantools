import { AppSidebar } from "@/components/app-sidebar";
import { SidebarInset, SidebarTrigger } from "@/components/ui/sidebar";
import { cookies } from "next/headers";
import { ReactNode } from "react";
import LogoutButton from "./logout-button";

export default async function ToolsLayout({
    children,
}: Readonly<{ children: ReactNode }>) {
    const access = (await cookies()).get("access_token")?.value;

    return (
        <>
            <AppSidebar accessToken={access} />
            <SidebarInset>
                <header className="flex h-16 shrink-0 justify-between me-4 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
                    <div className="flex items-center gap-2 px-4">
                        <SidebarTrigger className="-ml-1" />
                    </div>

                    <LogoutButton />
                </header>
                <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
                    {children}
                </div>
            </SidebarInset>
        </>
    );
}
