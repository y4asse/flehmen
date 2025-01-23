import React from "react";
export type WindowsProps = {
  windows: Window[];
};

type Window = {
  title?: string;
  initSize: { width: number; height: number };
  initPosition: { x: number; y: number; z: number };
  children: React.ReactNode;
};
export const MobileWindows = (props: WindowsProps) => {
  const { windows } = props;

  return <div>MobileWindows</div>;
};
