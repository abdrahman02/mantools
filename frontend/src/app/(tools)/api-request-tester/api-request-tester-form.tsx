"use client";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
} from "@/components/ui/select";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Textarea } from "@/components/ui/textarea";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { FormEvent, useState } from "react";

type APIRequestTesterRequestMethod = "GET" | "POST" | "PUT" | "DELETE";

interface APIRequestTesterRequest {
    method: APIRequestTesterRequestMethod;
    url: string;
    headers: string;
    body: string;
}

export default function APIRequestTesterForm() {
    const initialForm: APIRequestTesterRequest = {
        method: "GET",
        url: "",
        headers: '{"Content-Type": "application/json"}',
        body: "",
    };

    const [form, setForm] = useState<APIRequestTesterRequest>(initialForm);
    const [result, setResult] = useState<string>("");

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/api-request-tester`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        ...form,
                        headers: JSON.parse(form.headers ?? "{}"),
                    }),
                },
            );

            const { status, content, message, error } = await res.json();
            if (status !== 200) {
                console.error(error);
                showDialog("error");
            }

            setResult(content?.data);
            createAlert(status, message);
        } catch (err) {
            console.error(err);
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            hideDialog();
        }
    };

    const methodAPIRequesTesterOptions: APIRequestTesterRequestMethod[] = [
        "GET",
        "POST",
        "PUT",
        "DELETE",
    ];

    return (
        <form onSubmit={handleSubmit}>
            <div className="mb-3 flex flex-row justify-center items-center">
                <Select
                    value={form.method}
                    onValueChange={(value) =>
                        setForm((prev) => ({
                            ...prev,
                            method: value as APIRequestTesterRequestMethod,
                        }))
                    }
                    disabled={isLoading}
                    required
                >
                    <SelectTrigger className="w-1/12 font-bold">
                        {form.method}
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            {methodAPIRequesTesterOptions.map((method, i) => (
                                <SelectItem key={i} value={method}>
                                    {method}
                                </SelectItem>
                            ))}
                        </SelectGroup>
                    </SelectContent>
                </Select>
                <Input
                    className="w-10/12"
                    onChange={(e) =>
                        setForm((prev) => ({ ...prev, url: e.target.value }))
                    }
                    placeholder="https://example.com/api/...."
                    disabled={isLoading}
                    required
                />
                <Button type="submit">Send</Button>
            </div>
            <div className="flex flex-row justify-center items-center gap-x-3">
                <Card className="w-full">
                    <CardContent>
                        <Tabs defaultValue="body">
                            <TabsList className="grid w-full grid-cols-2">
                                <TabsTrigger value="headers">
                                    Headers
                                </TabsTrigger>
                                <TabsTrigger value="body">Body</TabsTrigger>
                            </TabsList>
                            <TabsContent value="headers">
                                <Textarea
                                    name="headers"
                                    rows={12}
                                    value={form.headers}
                                    onChange={(e) =>
                                        setForm((prev) => ({
                                            ...prev,
                                            headers: e.target.value,
                                        }))
                                    }
                                    placeholder="type your headers here..."
                                    disabled={isLoading}
                                    required
                                />
                            </TabsContent>
                            <TabsContent value="body">
                                <Textarea
                                    name="body"
                                    rows={12}
                                    value={form.body}
                                    onChange={(e) =>
                                        setForm((prev) => ({
                                            ...prev,
                                            body: e.target.value,
                                        }))
                                    }
                                    placeholder="type your json here..."
                                    disabled={isLoading}
                                />
                            </TabsContent>
                        </Tabs>
                    </CardContent>
                </Card>
                <Card className="w-full">
                    <CardContent>
                        <Label>Output</Label>
                        <Textarea rows={14} value={result} readOnly />
                    </CardContent>
                </Card>
            </div>
        </form>
    );
}
