import React from "react";
import { MobileWindow } from "./MobileWindow";
export type MobileProps = {
  windows: Mobile[];
};

type Mobile = {
  title?: string;
  initSize: { width: number; height: number };
  initPosition: { x: number; y: number; z: number };
  children: React.ReactNode;
};
export const MobileWindows = (props: MobileProps) => {
  const { windows } = props;

  return (
    <div
      style={{
        position: "absolute",
        top: "0",
        left: "0",
        width: "100vw",
        maxHeight: "100vh",
        // height: "100vh",
        display: "flex",
        flexDirection: "column",
        // justifyContent: "center",
        alignItems: "center",
        marginTop: "50px",
      }}
    >
      {windows.map((mobile, index) => (
        // <div
        //   key={index}
        //   // className="flex flex-col gap-4"
        //   style={{
        //     width: mobile.initSize.width,
        //     height: mobile.initSize.height,
        //     // position: "absolute",
        //     // left: mobile.initPosition.x,
        //     // top: mobile.initPosition.y,
        //     // zIndex: mobile.initPosition.z,
        //   }}
        // >
        //   {mobile.children}
        // </div>
        <MobileWindow
          key={index}
          initSize={mobile.initSize}
          title={mobile.title}
        >
          {mobile.children}
        </MobileWindow>
      ))}
    </div>
  );
};
