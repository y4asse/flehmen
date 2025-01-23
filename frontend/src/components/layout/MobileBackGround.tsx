import React from "react";
import Image from "next/image";
import { Flex } from "../ui/flex";
import Link from "next/link";

const bottomNavItems = [
  { name: "ホーム", href: "/", icon: "/images/file.svg" },
  { name: "プロフ", href: "/profile", icon: "/images/file.svg" },
  { name: "FF", href: "/ff", icon: "/images/file.svg" },
  { name: "リプ", href: "/reply", icon: "/images/file.svg" },
  { name: "ひみつ", href: "/secret", icon: "/images/file.svg" },
];

export const MobileBackGround = () => {
  return (
    <Flex className="bg-black h-[100vh] w-[100vw] absolute top-0 left-0">
      <div className="fixed bottom-2  flex justify-between w-[95%] bg-zinc-800 px-2 py-2 rounded-sm">
        {bottomNavItems.map((item) => (
          <Link
            href={item.href}
            key={item.name}
            className="flex flex-col items-center gap-2"
          >
            <Image src={item.icon} alt="icon" width={35} height={35} />
            <span className="text-white">{item.name}</span>
          </Link>
        ))}
      </div>
    </Flex>
  );
};
