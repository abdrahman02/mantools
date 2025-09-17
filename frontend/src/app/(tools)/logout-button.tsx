"use client";
import { Button } from "@/components/ui/button";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useAuth } from "@/hooks/use-auth";
import { useGlobalDialog } from "@/hooks/use-global-dialog";

export default function LogoutButton() {
    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();
    const { access, setAccess } = useAuth();

    if (!access) return null;

    const handleLogout = async () => {
        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/auth/logout`,
                {
                    method: "POST",
                    headers: { Authorization: `Bearer ${access}` },
                    credentials: "include",
                },
            );

            const { status, message } = await res.json();
            if (status !== 200) showDialog("error");

            createAlert(status, message);
            console.log({ status, message });
            if (status === 200) return setAccess(null);
        } catch (err) {
            console.error(err);
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            hideDialog();
        }
    };

    return (
        <Button
            onClick={handleLogout}
            variant="reverse"
            className="h-7"
            disabled={isLoading}
        >
            Logout
        </Button>
    );
}
