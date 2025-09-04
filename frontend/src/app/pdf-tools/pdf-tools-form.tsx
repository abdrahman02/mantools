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

interface PDFToolsRequest {
    pdfAction: string;
    files: File[];
    rangePage: number[];
}

export default function PDFToolsForm() {
    const initFormVal: PDFToolsRequest = {
        pdfAction: "",
        files: [],
        rangePage: [],
    };
    const [form, setForm] = useState<PDFToolsRequest>(initFormVal);

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        if (form.pdfAction === "" || form.files.length < 1) {
            showDialog("error");
            createAlert(400, "Fil the required fields");
            setTimeout(() => hideDialog(), 1000);
            return;
        }

        const formData = new FormData();
        formData.append("pdfAction", form.pdfAction);
        form.files.forEach((file) => formData.append("files", file));
        if (form.pdfAction === "split_pdf") {
            formData.append("rangePage", form.rangePage.join("-"));
        }

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/pdf-tools`,
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

            const message = res.headers.get("X-Message") ?? "";
            const blob = await res.blob();

            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            const outFilename: Record<string, string> = {
                merge_pdf: "pdf_merged.pdf",
                split_pdf: "pdf_splited.zip",
                compress_pdf: "pdf_compressed.zip",
            };

            a.href = url;
            a.download = outFilename[form.pdfAction];
            a.click();
            a.remove();

            createAlert(200, message ?? "");
        } catch (err) {
            console.error(err);
            showDialog("error");
            createAlert(500, "Something went wrong");
        } finally {
            setTimeout(() => hideDialog(), 1000);
        }
    };

    const pdfActionOptions = ["Merge PDF", "Split PDF", "Compress PDF"];

    return (
        <form onSubmit={handleSubmit}>
            <div className="mb-3"></div>
            <div className="flex flex-col justify-center items-center gap-y-3">
                <Card className="w-full">
                    <CardContent className="flex flex-col justify-center items-center gap-x-3">
                        <div className="w-full flex flex-row justify-center items-center gap-x-3">
                            <div className="w-1/2">
                                <Label>PDF Action</Label>
                                <Select
                                    value={form.pdfAction}
                                    onValueChange={(value) =>
                                        setForm((prev) => ({
                                            ...prev,
                                            pdfAction: value,
                                        }))
                                    }
                                    disabled={isLoading}
                                    required
                                >
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder="Select a pdf action" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectGroup>
                                            {pdfActionOptions
                                                .reverse()
                                                .map((act, i) => (
                                                    <SelectItem
                                                        key={i}
                                                        value={act
                                                            .toLowerCase()
                                                            .split(" ")
                                                            .join("_")}
                                                    >
                                                        {act}
                                                    </SelectItem>
                                                ))}
                                        </SelectGroup>
                                    </SelectContent>
                                </Select>
                            </div>
                            <div className="w-1/2">
                                <Label>PDF File(s)</Label>
                                <Input
                                    id="pdf_file"
                                    type="file"
                                    accept="application/pdf"
                                    multiple={
                                        form.pdfAction === "merge_pdf" ||
                                        form.pdfAction === "compress_pdf"
                                    }
                                    onChange={(e) =>
                                        setForm((prev) => ({
                                            ...prev,
                                            files: Array.from(
                                                e.target.files ?? [],
                                            ),
                                        }))
                                    }
                                    disabled={isLoading}
                                    required
                                />
                            </div>
                        </div>
                        {form.pdfAction === "split_pdf" && (
                            <div className="w-full flex flex-row gap-x-3">
                                <div className="w-1/2">
                                    <Label>Started page</Label>
                                    <Input
                                        type="number"
                                        onChange={(e) =>
                                            setForm((prev) => {
                                                const newRange = [
                                                    ...prev.rangePage,
                                                ];
                                                newRange[0] = Number(
                                                    e.target.value,
                                                );
                                                return {
                                                    ...prev,
                                                    rangePage: newRange,
                                                };
                                            })
                                        }
                                        placeholder="Started page..."
                                        disabled={isLoading}
                                        required={
                                            form.pdfAction === "split_pdf"
                                        }
                                    />
                                </div>
                                <div className="w-1/2">
                                    <Label>Ending page</Label>
                                    <Input
                                        type="number"
                                        onChange={(e) =>
                                            setForm((prev) => {
                                                const newRange = [
                                                    ...prev.rangePage,
                                                ];
                                                newRange[1] = Number(
                                                    e.target.value,
                                                );
                                                return {
                                                    ...prev,
                                                    rangePage: newRange,
                                                };
                                            })
                                        }
                                        placeholder="Ending page..."
                                        disabled={isLoading}
                                        required={
                                            form.pdfAction === "split_pdf"
                                        }
                                    />
                                </div>
                            </div>
                        )}
                    </CardContent>
                    <CardFooter className="flex flex-row justify-between items-center">
                        <Button
                            type="button"
                            onClick={() => setForm(initFormVal)}
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
