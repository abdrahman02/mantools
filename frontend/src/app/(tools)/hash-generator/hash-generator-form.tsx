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
import { Slider } from "@/components/ui/slider";
import { Textarea } from "@/components/ui/textarea";
import { useAlertMessage } from "@/hooks/use-alert-message";
import { useGlobalDialog } from "@/hooks/use-global-dialog";
import { FormEvent, useState } from "react";

interface HashGeneratorRequest {
    input: string;
    algorithm: string;
    costFactor: number;
}

export default function HashGeneratorForm() {
    const initValForm: HashGeneratorRequest = {
        input: "",
        algorithm: "",
        costFactor: 10,
    };
    const [form, setForm] = useState<HashGeneratorRequest>(initValForm);
    const [result, setResult] = useState("");

    const { isLoading, showDialog, hideDialog } = useGlobalDialog();
    const { createAlert } = useAlertMessage();

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        showDialog();

        try {
            const res = await fetch(
                `${process.env.NEXT_PUBLIC_API_BASE_URL}/hash-generator`,
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
            hideDialog();
        }
    };

    const algorithmOptions = ["md5", "sha1", "sha256", "bcrypt"];

    return (
        <form onSubmit={handleSubmit}>
            <div className="mb-3">
                <Label>Algorithm</Label>
                <Select
                    value={form.algorithm}
                    onValueChange={(value) =>
                        setForm((prev) => ({ ...prev, algorithm: value }))
                    }
                    disabled={isLoading}
                    required
                >
                    <SelectTrigger className="w-full">
                        <SelectValue placeholder="Select a algorithm" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                            {algorithmOptions.map((opt, i) => (
                                <SelectItem key={i} value={opt}>
                                    {String(opt.toUpperCase())}
                                </SelectItem>
                            ))}
                        </SelectGroup>
                    </SelectContent>
                </Select>
            </div>
            <div className="flex flex-col justify-center items-center gap-y-3 sm:gap-x-3 sm:flex-row">
                <Card className="w-full">
                    <CardContent>
                        <div>
                            <Label>Input</Label>
                            <Textarea
                                name="input"
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
                        </div>
                        {form.algorithm === "bcrypt" && (
                            <div>
                                <Label>Cost Factor: {form.costFactor}</Label>
                                <Slider
                                    min={1}
                                    max={20}
                                    step={1}
                                    value={[form.costFactor]}
                                    onValueChange={(val) =>
                                        setForm((prev) => ({
                                            ...prev,
                                            costFactor: val[0],
                                        }))
                                    }
                                />
                            </div>
                        )}
                    </CardContent>
                    <CardFooter className="flex flex-row justify-between items-center">
                        <Button
                            type="button"
                            onClick={() => setForm(initValForm)}
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
                        <Textarea rows={17} value={result} readOnly />
                    </CardContent>
                </Card>
            </div>
        </form>
    );
}
