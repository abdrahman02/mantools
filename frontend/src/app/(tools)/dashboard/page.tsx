import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import VisitorChart from "./visitor-chart";
import PopularPageChart from "./popular-page-chart";
import MetricCard from "./metric-card";

export default async function DashboardPage() {
    const access = (await cookies()).get("access_token")?.value;

    const res = await fetch(`${process.env.API_BASE_URL}/dashboard/chart`, {
        headers: { Authorization: `Bearer ${access}` },
        cache: "no-store",
    });
    if (!res.ok) {
        return redirect("/error");
    }

    const { content } = await res.json();
    return (
        <div className="flex flex-col gap-y-3 sm:gap-x-3 sm:flex-row">
            <div className="w-full sm:w-4/6 flex flex-col gap-y-2">
                <VisitorChart data={content?.visitors} />
                <PopularPageChart data={content?.popularPage} />
            </div>
            <div className="w-full sm:w-2/6">
                <MetricCard
                    title="Visitor by Country"
                    description="Showing visitors in the last 30 days"
                    data={content.usersByCountry}
                />
            </div>
        </div>
    );
}
