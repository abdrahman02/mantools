"use client";
import { Button } from "@/components/ui/button";
import { CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useAuth } from "@/hooks/use-auth";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { useRouter } from "next/navigation";
import { ChangeEvent, FormEvent, useState } from "react";

interface LoginRequest {
    email: string;
    password: string;
}

export default function LoginForm() {
    const initValForm = {
        email: "",
        password: "",
    };
    const [form, setForm] = useState<LoginRequest>(initValForm);

    const router = useRouter();
    const { setAccess } = useAuth();
    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { id, value } = e.target;
        setForm((prev) => ({ ...prev, [id]: value }));
    };

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/auth/login`,
                {
                    method: "POST",
                    body: JSON.stringify(form),
                    credentials: "include",
                },
            );

            const { status, content, message } = await res.json();
            createAlert(status, message);
            if (!res.ok) {
                showDialog("error");
                return;
            }
            setAccess(content?.data ? content?.data.access : null);
            router.push("/dashboard");
        } catch (err) {
            console.error({ err });
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            hideDialog();
        }
    };

    return (
        <CardContent>
            <form onSubmit={handleSubmit}>
                <div className="flex flex-col gap-6">
                    <div className="grid gap-2">
                        <Label htmlFor="email">Email</Label>
                        <Input
                            id="email"
                            type="email"
                            placeholder="m@example.com"
                            onChange={handleChange}
                            value={form.email}
                            disabled={isLoading}
                            required
                        />
                    </div>
                    <div className="grid gap-2">
                        <Label htmlFor="password">Password</Label>
                        <Input
                            id="password"
                            type="password"
                            placeholder="********"
                            onChange={handleChange}
                            value={form.password}
                            disabled={isLoading}
                            required
                        />
                    </div>
                    <Button
                        type="submit"
                        className="w-full font-semibold"
                        disabled={isLoading}
                    >
                        {isLoading ? "Loading" : "Login"}
                    </Button>
                </div>
            </form>
        </CardContent>
    );
}
