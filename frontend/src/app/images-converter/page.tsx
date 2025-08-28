import AlertMessage from "@/components/alert-message";
import ImagesConverterForm from "./images-converter-form";
export default function ImagesConverterPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">Images Converter</h1>
            <AlertMessage />
            <ImagesConverterForm />
        </>
    );
}
