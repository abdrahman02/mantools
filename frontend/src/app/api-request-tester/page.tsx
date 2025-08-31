import APIRequestTesterForm from "./api-request-tester-form";
import AlertMessage from "@/components/alert-message";

export default function APIRequestTesterPage() {
    return (
        <>
            <h1 className="text-2xl text-center font-bold">Text Formatter</h1>
            <AlertMessage />
            <APIRequestTesterForm />
        </>
    );
}
