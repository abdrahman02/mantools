import HashGeneratorForm from "./hash-generator-form";
import AlertMessage from "@/components/alert-message";

export default function HashGeneratorPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">Hash Generator</h1>
            <AlertMessage />
            <HashGeneratorForm />
        </>
    );
}
