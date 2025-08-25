export interface AlertMessage {
    status: number;
    message: string;
    createAlert: (status: number, message: string) => void;
    dismissAlert: () => void;
}
