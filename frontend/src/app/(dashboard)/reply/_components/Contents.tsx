import React from "react";
import { Flex } from "@/components/ui/flex";
import Image from "next/image";
import { content } from "./ReplyPage";

type Props = {
  repContents: content[][];
  userIndex: number;
};

const Contents = ({ repContents, userIndex }: Props) => {
  const contents = repContents[userIndex];
  return (
    <Flex
      direction={"column"}
      className="gap-8 h-full py-8 mb-2 w-full ~pl-8/20"
      justify={"start"}
      align={"start"}
    >
      {contents.map((content) => (
        <a
          href={`https://twitter.com/${content.userId}`}
          target="_blank"
          key={content.id}
        >
          <Flex
            className="gap-3"
            direction={"column"}
            justify={"start"}
            align={"start"}
          >
            <Flex className="gap-2 w-full" justify={"start"}>
              <Image
                src={content.icon}
                alt={content.name}
                width={50}
                height={50}
                className="rounded-full"
              />
              <p className="text-white">{content.name}</p>
              <p className="text-[#aaa]">@{content.userId}</p>
            </Flex>
            <p className="text-white">{content.value}</p>
          </Flex>
        </a>
      ))}
    </Flex>
  );
};

export default Contents;
