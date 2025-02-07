"use client";
import { Flex } from "@/components/ui/flex";
import Image from "next/image";

type WindowProps = {
  title?: string;
  initSize: { width: number | string; height: number | string };
  children: React.ReactNode;
};

export const MobileWindow = (props: WindowProps) => {
  const { children, initSize } = props;
  const { width, height } = initSize;

  const handleJudgeType = (value: number | string) => {
    if (typeof value === "number") {
      return `${value}px`;
    } else {
      return value;
    }
  };

  const judgeWidth = handleJudgeType(width);
  const judgeHeight = handleJudgeType(height);

  return (
    <Flex
      direction={"column"}
      style={{
        width: judgeWidth,
        maxWidth: "93vw",
        height: judgeHeight,
        backgroundColor: "#000",
        border: "0.7px solid #e4007f",
        borderRadius: "10px",
        margin: "10px",
        padding: "20px 0",
        position: "relative",
      }}
    >
      {/* コンテンツ部分 */}
      <Flex
        className="w-full h-full overflow-y-scroll max-h-full overflow-x-hidden relative z-10"
        style={{
          scrollbarWidth: "none", // Firefox用
          msOverflowStyle: "none", // IE用
        }}
      >
        {children}
      </Flex>
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
          opacity: 0.3,
          filter: "grayscale(100%)", // 白黒に変換
          pointerEvents: "none",
          zIndex: 1,
        }}
      />
    </Flex>
  );
};
