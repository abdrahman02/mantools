import AlertMessage from "@/components/alert-message";
import PDFToolsForm from "./pdf-tools-form";
export default function PDFToolsPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">PDF Tools</h1>
            <AlertMessage />
            <PDFToolsForm />
        </>
    );
}
