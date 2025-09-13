export interface Auth {
    access: string | null;
    setAccess: (token: string | null) => void;
}
