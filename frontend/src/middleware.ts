import { NextRequest, NextResponse } from "next/server";

export const config = {
    matcher: ["/dashboard"],
};

export default async function middleware(req: NextRequest) {
    const refresh = req.cookies.get("refresh_token")?.value;
    const access = req.cookies.get("access_token")?.value;
    if (!refresh && !access)
        return NextResponse.redirect(new URL("/login", req.url));

    if (access) return NextResponse.next();

    if (refresh) {
        try {
            const res = await fetch(
                `${process.env.API_BASE_URL}/auth/refresh`,
                {
                    method: "POST",
                    headers: {
                        cookie: `refresh_token=${refresh}`,
                    },
                },
            );

            if (!res.ok)
                return NextResponse.redirect(new URL("/login", req.url));

            const setCookie = res.headers.get("set-cookie");
            const response = NextResponse.next();

            if (setCookie) response.headers.set("set-cookie", setCookie);
            return response;
        } catch (err) {
            console.error("Middleware refresh error:", err);
            return NextResponse.redirect(new URL("/login", req.url));
        }
    }

    return NextResponse.next();
}
