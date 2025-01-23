import React from "react";
import BackGround from "@/components/layout/BackGround";
import { MobileBackGround } from "@/components/layout/MobileBackGround";

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <div>
      <div className="hidden md:block">
        <BackGround />
      </div>
      <div className="block md:hidden">
        <MobileBackGround />
      </div>
      {children}
    </div>
  );
}
