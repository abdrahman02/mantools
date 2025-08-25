"use client";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { FormEvent, useState } from "react";

interface TextFormatterRequest {
    input: string;
    format: string;
}

export default function TextFormatterForm() {
    const [form, setForm] = useState<TextFormatterRequest>({
        input: "",
        format: "",
    });
    const [result, setResult] = useState("");

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/text-formatter`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(form),
                },
            );

            const { status, content, message } = await res.json();
            if (status !== 200) {
                showDialog("error");
            }

            setResult(content?.data);
            createAlert(status, message);
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
            <div className="mb-3">
                <Label>Format</Label>
                <Select
                    value={form.format}
                    onValueChange={(value) =>
                        setForm((prev) => ({ ...prev, format: value }))
                    }
                    disabled={isLoading}
                >
                    <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select a fruit" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            <SelectItem value="json">JSON</SelectItem>
                            <SelectItem value="xml">XML</SelectItem>
                            <SelectItem value="sql" disabled>
                                SQL
                            </SelectItem>
                            <SelectItem value="markdown" disabled>
                                Markdown
                            </SelectItem>
                        </SelectGroup>
                    </SelectContent>
                </Select>
            </div>
            <div className="flex flex-row justify-center items-center gap-x-3">
                <Card className="w-full">
                    <CardContent>
                        <Label>Input</Label>
                        <Textarea
                            name="output"
                            rows={12}
                            value={form.input}
                            onChange={(e) =>
                                setForm((prev) => ({
                                    ...prev,
                                    input: e.target.value,
                                }))
                            }
                            disabled={isLoading}
                        />
                    </CardContent>
                    <CardFooter className="flex flex-row justify-between items-center">
                        <Button
                            type="button"
                            onClick={() =>
                                setForm((prev) => ({ ...prev, input: "" }))
                            }
                            className="bg-white"
                        >
                            Clear
                        </Button>
                        <Button type="submit">Submit</Button>
                    </CardFooter>
                </Card>
                <Card className="w-full">
                    <CardContent>
                        <Label>Output</Label>
                        <Textarea rows={15} value={result} readOnly />
                    </CardContent>
                </Card>
            </div>
        </form>
    );
}
