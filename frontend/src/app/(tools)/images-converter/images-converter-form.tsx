"use client";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { FormEvent, useState } from "react";

interface ImagesConverterRequest {
    targetConvert: string;
    files: File[];
}

export default function ImagesConverterForm() {
    const [form, setForm] = useState<ImagesConverterRequest>({
        targetConvert: "",
        files: [],
    });

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        const formData = new FormData();
        formData.append("targetConvert", form.targetConvert);
        if (form.files.length > 0)
            form.files.forEach((file) => formData.append("files", file));

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/images-converter`,
                {
                    method: "POST",
                    body: formData,
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
            a.download = "images_converted.zip";
            a.click();
            createAlert(200, message ?? "");
            a.remove();
        } catch (err) {
            console.error(err);
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            hideDialog();
        }
    };

    const targetConvertOptions = ["jpg", "png", "webp"];

    return (
        <form onSubmit={handleSubmit}>
            <div className="mb-3"></div>
            <div className="flex flex-col justify-center items-center gap-y-3">
                <Card className="w-full">
                    <CardContent className="flex flex-col justify-center items-center gap-y-3 sm:gap-x-3 sm:flex-row">
                        <div className="w-full sm:w-1/2">
                            <Label>Target Convert</Label>
                            <Select
                                value={form.targetConvert}
                                onValueChange={(value) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        targetConvert: value,
                                    }))
                                }
                                required
                                disabled={isLoading}
                            >
                                <SelectTrigger className="w-full">
                                    <SelectValue placeholder="Select a images target convert" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectGroup>
                                        {targetConvertOptions.map((tc, i) => (
                                            <SelectItem
                                                key={i}
                                                value={tc}
                                            >{`${tc.toUpperCase()}`}</SelectItem>
                                        ))}
                                    </SelectGroup>
                                </SelectContent>
                            </Select>
                        </div>
                        <div className="w-full sm:w-1/2">
                            <Label>Images</Label>
                            <Input
                                id="images"
                                type="file"
                                accept="image/*"
                                multiple
                                required
                                onChange={(e) =>
                                    setForm((prev) => ({
                                        ...prev,
                                        files: Array.from(e.target.files ?? []),
                                    }))
                                }
                                disabled={isLoading}
                            />
                        </div>
                    </CardContent>
                    <CardFooter className="flex flex-row justify-between items-center">
                        <Button
                            type="button"
                            onClick={() =>
                                setForm((prev) => ({ ...prev, files: [] }))
                            }
                            className="bg-white"
                        >
                            Clear
                        </Button>
                        <Button type="submit">Submit</Button>
                    </CardFooter>
                </Card>
            </div>
        </form>
    );
}
