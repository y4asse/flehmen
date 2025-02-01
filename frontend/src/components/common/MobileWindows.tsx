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
      {windows.map((mobile, index) => (
        <div
          key={index}
          className="flex flex-col gap-4"
          style={{
            width: mobile.initSize.width,
            height: mobile.initSize.height,
            position: "absolute",
            left: mobile.initPosition.x,
            top: mobile.initPosition.y,
            zIndex: mobile.initPosition.z,
          }}
        >
          {mobile.children}
        </div>
      ))}
    </div>
  );
};
