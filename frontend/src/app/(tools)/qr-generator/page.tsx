import AlertMessage from "@/components/alert-message";
import QRGeneratorForm from "./qr-generator-form";
export default function QRGeneratorPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">QR Generator</h1>
            <AlertMessage />
            <QRGeneratorForm />
        </>
    );
}
