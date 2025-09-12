import TextFormatterForm from "./text-formatter-form";
import AlertMessage from "@/components/alert-message";

export default function TextFormatterPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">Text Formatter</h1>
            <AlertMessage />
            <TextFormatterForm />
        </>
    );
}
