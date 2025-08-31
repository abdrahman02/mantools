import JWTDecoderForm from "./jwt-decoder-form";
import AlertMessage from "@/components/alert-message";

export default function JWTDecoderPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">JWT Decoder</h1>
            <AlertMessage />
            <JWTDecoderForm />
        </>
    );
}
