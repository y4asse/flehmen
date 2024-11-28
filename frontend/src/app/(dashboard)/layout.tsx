import React from 'react';
import  BackGround   from '@/components/layout/BackGround';

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return <div>
    <BackGround />
    {children}
  </div>;
}

