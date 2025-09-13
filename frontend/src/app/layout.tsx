import type { Metadata } from "next";
import { Space_Grotesk } from "next/font/google";
import "./globals.css";
import { SidebarProvider } from "@/components/ui/sidebar";
import { GlobalDialogProvider } from "@/contexts/global-dialog-context";
import { AlertMessageProvider } from "@/contexts/alert-message-context";
import { AuthProvider } from "@/contexts/auth-context";

const spaceGrotesk = Space_Grotesk({
  variable: "--font-space-grotesk",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: {
    template: "Mantools | %s", // %s will be replaced by the page-specific title
    default: "Mantools", // Fallback title
  },
  description:
    "Mantools is a versatile toolbox website that provides a variety of online tools such as formatters, converters, generators, and other productivity utilities for free and easy to use.",
  icons: {
    icon: "/thumbs-up.ico", // Path relative to the public directory
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${spaceGrotesk.variable} antialiased`}>
        <AuthProvider>
          <AlertMessageProvider>
            <GlobalDialogProvider>
              <SidebarProvider>{children}</SidebarProvider>
            </GlobalDialogProvider>
          </AlertMessageProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
