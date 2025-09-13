import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default async function DashboardPage() {
    const access = (await cookies()).get("access_token")?.value;
    if (!access) return redirect("/login");

    return <>Dashboard Page</>;
}
