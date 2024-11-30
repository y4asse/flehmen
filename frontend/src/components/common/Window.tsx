"use client";
import React, { useState } from "react";
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

  const [position, setPosition] = useState({ x, y, z });
  const [size, setSize] = useState({ width, height });
  const [isFullScreen, setIsFullScreen] = useState(false);

  // フルスクリーンの処理
  const handleFullScreen = () => {
    setSize({ width: window.innerWidth, height: window.innerHeight });
    setPosition({ x: 0, y: 0, z: 999 });
    setIsFullScreen(true);
  };

  // 初期サイズに戻す処理
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
        boxShadow: isActive ? "0px 0px 20px rgba(228, 0, 127)" : "0px 0px 10px rgba(128, 128, 128, 0.4)", // 影の変更
      }}
      onClick={handleClickWindow}
    >
    
    {isActive && (
  <div
    style={{
      position: "absolute",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundImage: "url('/images/hart.svg')",
      backgroundSize: "60%", // サイズを60%に
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      opacity: 0.4, // 透明度設定
      pointerEvents: "none", // クリックイベントを透過
      zIndex: -1,
    }}
  />
)}

{!isActive && (
  <div
    style={{
      position: "absolute",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      backgroundImage: "url('/images/hart.svg')",
      backgroundSize: "60%",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      opacity: 0.4,
      filter: "grayscale(100%)", // 白黒に変換
      pointerEvents: "none",
      zIndex: -1,
    }}
  />
)}


      {/* ヘッダー部分 */}
      <Flex
        className={`w-full h-10 rounded-sm py-2 ${isActive ? "bg-primary" : "bg-gray-500"}`}
        justify={"start"}
      >
        <Button
          className={`w-8 h-6 p-0 rounded-none ${
            isActive ? "bg-primary" : "bg-gray-500"
          }`}
          
          onClick={isFullScreen ? handleInitScreen : handleFullScreen}
        >
          <div className="w-4 h-4 border-[#D9D9D9] border-[3px] rounded-sm"></div>
        </Button>
        <p className="text-[#eee] ml-2">{title}</p>
      </Flex>

      {/* コンテンツ部分 */}
      <Flex className="w-full h-full">
        {children}
      </Flex>
    </Flex>
  );
};
