import AlertMessage from "@/components/alert-message";
import ImagesCompressorForm from "./images-compressor-form";
export default function ImagesCompressorPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">
                Images Compressor
            </h1>
            <AlertMessage />
            <ImagesCompressorForm />
        </>
    );
}
