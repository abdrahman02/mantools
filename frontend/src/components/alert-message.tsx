"use client";

import { CheckCircle2Icon, XCircle } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { Button } from "./ui/button";
import { useAlertMessage } from "@/hooks/use-alert-message";

export default function AlertMessage() {
    const { status, message, dismissAlert } = useAlertMessage();
    if (!status || !message) return;

    const successTitle = "Success! Your request has been completed";
    const errorTitle = "Error! Your request has been canceled";

    return (
        <Alert variant={status === 200 ? "default" : "destructive"}>
            {status === 200 ? <CheckCircle2Icon /> : <XCircle />}
            <AlertTitle>
                {status === 200 ? successTitle : errorTitle}
            </AlertTitle>
            <AlertDescription>{message}</AlertDescription>
            <Button
                onClick={() => dismissAlert()}
                size="sm"
                variant="noShadow"
                className="absolute top-2.5 right-3 h-6 bg-secondary-background text-foreground"
            >
                Close
            </Button>
        </Alert>
    );
}
