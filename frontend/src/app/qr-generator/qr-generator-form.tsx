"use client";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { FormEvent, useState } from "react";

export default function QRGeneratorForm() {
    const [qrContent, setQrContent] = useState("");

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/qr-generator`,
                {
                    method: "POST",
                    body: JSON.stringify({ qrContent }),
                },
            );

            if (!res.ok) {
                const { status, message, error } = await res.json();
                console.error(error);
                showDialog("error");
                createAlert(status, message);
                return;
            }

            const message = res.headers.get("X-Message");

            const blob = await res.blob();
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.href = url;
            a.download = "qr_generated.png";
            a.click();

            createAlert(200, message ?? "");
        } catch (err) {
            console.error(err);
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            setTimeout(() => {
                hideDialog();
            }, 1000);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div className="mb-3"></div>
            <div className="flex flex-col justify-center items-center gap-y-3">
                <Card className="w-full">
                    <CardContent className="flex flex-col justify-center items-start gap-y-2">
                        <Label>Qr Content</Label>
                        <Textarea
                            name="input"
                            rows={12}
                            value={qrContent}
                            onChange={(e) => setQrContent(e.target.value)}
                            disabled={isLoading}
                        />
                    </CardContent>
                    <CardFooter className="flex flex-row justify-between items-center">
                        <Button
                            type="button"
                            onClick={() => setQrContent("")}
                            className="bg-white"
                        >
                            Clear
                        </Button>
                        <Button type="submit" disabled={isLoading}>
                            Submit
                        </Button>
                    </CardFooter>
                </Card>
            </div>
        </form>
    );
}
