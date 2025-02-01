import React from "react";
import Image from "next/image";
import { Flex } from "../ui/flex";
import Link from "next/link";

export const MobileBackGround = () => {
  return (
    <Flex className="bg-black h-[100vh] w-[100vw] absolute top-0 left-0">
      <Image
        src="/images/cat_bg.svg"
        alt="background"
        width={300}
        height={300}
        className="z-0"
        style={{ opacity: 0.3 }}
      />
      <div className="fixed bottom-2  flex justify-center w-[95%] px-2 py-2 rounded-sm">
        <Link href="/">
          <Image src="/images/logo.svg" alt="icon" width={200} height={200} />
        </Link>
      </div>
    </Flex>
  );
};
