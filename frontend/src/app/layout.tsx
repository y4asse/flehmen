import type { Metadata } from "next";
import { ClerkProvider, SignedIn, UserButton } from "@clerk/nextjs";
import "./globals.css";
import "./styles.css";
import { DotGothic16 } from "next/font/google";

const dotGothic16 = DotGothic16({
  weight: "400",
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "Flehmen",
  description: "恋愛最強ツール",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <ClerkProvider>
      <html lang="ja" className={dotGothic16.className}>
        <body>
        <header className="absolute z-10 top-0 right-0 flex justify-end items-center p-4 gap-4 h-16 text-foreground">
            <SignedIn>
              <UserButton />
            </SignedIn>
          </header>
          <main>{children}</main>
        </body>
      </html>
    </ClerkProvider>
  );
}
