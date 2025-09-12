import {
    Card,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import LoginForm from "./login-form";
import AlertMessage from "@/components/alert-message";

export default function LoginPage() {
    return (
        <div className="w-full h-container flex flex-col justify-center items-center">
            <div className="w-full max-w-sm mb-2">
                <AlertMessage />
            </div>
            <Card className="w-full max-w-sm">
                <CardHeader>
                    <CardTitle>Login to your account</CardTitle>
                    <CardDescription>
                        Enter your email below to login to your account
                    </CardDescription>
                </CardHeader>
                <LoginForm />
                <CardFooter className="text-sm text-gray-600 font-thin flex-col gap-2">
                    &copy; 2025 M. Abdul Rahman. All Rights Reserved.
                </CardFooter>
            </Card>
        </div>
    );
}
