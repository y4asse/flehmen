"use client";
import { Flex } from "@/components/ui/flex";

type WindowProps = {
  title?: string;
  initSize: { width: number; height: number };
  children: React.ReactNode;
};

export const MobileWindow = (props: WindowProps) => {
  const { children, initSize } = props;
  const { width, height } = initSize;

  console.log(initSize);

  return (
    <Flex
      direction={"column"}
      style={{
        width: `${width}px`,
        maxWidth: "93vw",
        height: `${height}px`,
        backgroundColor: "#000",
        border: "0.7px solid #e4007f",
        borderRadius: "10px",
        margin: "10px",
        padding: "20px 0",
      }}
    >
      {/* コンテンツ部分 */}
      <Flex
        className="w-full h-full overflow-y-scroll max-h-full overflow-x-hidden"
        style={{
          scrollbarWidth: "none", // Firefox用
          msOverflowStyle: "none", // IE用
        }}
      >
        {children}
      </Flex>
    </Flex>
  );
};
