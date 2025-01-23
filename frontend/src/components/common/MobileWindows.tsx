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

  return (
    <div className="flex flex-col gap-4">
      {windows.map((window, index) => (
        <div
          key={index}
          className="flex flex-col gap-4"
          style={{
            width: window.initSize.width,
            height: window.initSize.height,
            position: "absolute",
            left: window.initPosition.x,
            top: window.initPosition.y,
            zIndex: window.initPosition.z,
          }}
        >
          {window.children}
        </div>
      ))}
    </div>
  );
};
