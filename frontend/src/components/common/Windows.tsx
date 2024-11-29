"use client";
import React, { useState } from "react";
import { Window } from "@/components/common/Window";

export type WindowsProps = {
  windows: Window[];
};

type Window = {
  title?: string;
  initSize: { width: number; height: number };
  initPosition: { x: number; y: number; z: number };
  children: React.ReactNode;
};

export const Windows = (props: WindowsProps) => {
  const { windows } = props;
  const [activeWindowId, setActiveWindowId] = useState(0);
  const handleSetActive = (id: number) => {
    setActiveWindowId(id);
  };
  return (
    <>
      {windows.map((window, i) => (
        <Window
          key={i}
          initSize={window.initSize}
          initPosition={window.initPosition}
          title={window.title}
          isActive={activeWindowId === i}
          handleClickWindow={() => handleSetActive(i)}
        >
          {window.children}
        </Window>
      ))}
    </>
  );
};
