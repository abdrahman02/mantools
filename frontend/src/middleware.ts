import { NextRequest, NextResponse } from "next/server";

export const config = {
    matcher: ["/dashboard:path*", "/login"],
};

export default async function middleware(req: NextRequest) {
    const { pathname } = req.nextUrl;
    const loginPath = "/login";
    const refresh = req.cookies.get("refresh_token")?.value;
    const access = req.cookies.get("access_token")?.value;

    // If user has been authenticated and try to access /login -> redirect to /dashboard
    if (pathname === loginPath && access)
        return NextResponse.redirect(new URL("/dashboard", req.url));

    // If user is not authenticated and try to access /login -> continue the request
    if (pathname === loginPath && !access && !refresh)
        return NextResponse.next();

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

            if (!res.ok) {
                return NextResponse.redirect(new URL("/login", req.url));
            }

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
