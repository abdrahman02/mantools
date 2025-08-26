import TextCaseConverterForm from "./text-case-converter-form";
import AlertMessage from "@/components/alert-message";

export default function TextCaseConverterPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">
                Text Case Converter
            </h1>
            <AlertMessage />
            <TextCaseConverterForm />
        </>
    );
}
