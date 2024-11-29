"use client";
import React from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "../ui/button";

type WindowProps = {
  title?: string;
  initSize: { width: number; height: number };
  initPosition: { x: number; y: number; z: number };
  children: React.ReactNode;
  handleClickWindow: () => void;
  isActive: boolean;
};
export const Window = (props: WindowProps) => {
  const {
    title,
    children,
    initPosition,
    initSize,
    isActive,
    handleClickWindow,
  } = props;
  const { width, height } = initSize;
  const { x, y, z } = initPosition;

  const [position, setPosition] = React.useState({ x, y, z });
  const [size, setSize] = React.useState({ width, height });
  const [isFullScreen, setIsFullScreen] = React.useState(false);
  const handleFullScreen = () => {
    setSize({ width: window.innerWidth, height: window.innerHeight });
    setPosition({ x: 0, y: 0, z: 999 });
    setIsFullScreen(true);
  };

  const handleInitScreen = () => {
    setSize(initSize);
    setPosition(initPosition);
    setIsFullScreen(false);
  };

  return (
    <Flex
      direction={"column"}
      style={{
        width: `${size.width}px`,
        height: `${size.height}px`,
        position: "absolute",
        left: `${position.x}px`,
        top: `${position.y}px`,
        backgroundColor: "#000",
        borderRadius: "5px",
        zIndex: isActive ? 999 : position.z,
      }}
      onClick={handleClickWindow}
    >
      <Flex
        className="w-full h-10 bg-primary rounded-sm py-2"
        justify={"start"}
      >
        <Button
          className="w-8 h-6 p-0 rounded-none"
          onClick={isFullScreen ? handleInitScreen : handleFullScreen}
        >
          <div className="w-4 h-4  border-[#D9D9D9] border-[3px] rounded-sm"></div>
        </Button>
        <p className="text-[#eee]">{title}</p>
      </Flex>
      <Flex className="w-full h-full">{children}</Flex>
    </Flex>
  );
};
