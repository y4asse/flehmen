import React from "react";
import { reply } from "../page";
import { Flex } from "@/components/ui/flex";
import Image from "next/image";

type rep = {
  replyInfo?: reply[];
  handleSetUserIndex: (index: number) => void;
};

const ReplyList = ({ replyInfo = [], handleSetUserIndex }: rep) => {
  return (
    <Flex
      direction={"column"}
      className="gap-8 h-full py-8 mb-2"
      justify={"start"}
      align={"start"}
    >
      {replyInfo.map((reply, index) => (
        <Flex
          className="gap-3"
          key={`${reply.id}-${index}`}
          direction={"column"}
          onClick={() => handleSetUserIndex(index)}
        >
          <Flex className="gap-2">
            <p className="text-white">□ </p>
            <Image
              src={reply.icon}
              alt={reply.name}
              width={50}
              height={50}
              className="rounded-full"
            />
            <p className="text-white">{reply.name}</p>
            <p className="text-[#aaa]">@{reply.userId}</p>
          </Flex>
          <Flex
            direction={"column"}
            justify={"start"}
            align={"start"}
            className="ml-8 gap-2"
          >
            <p className="text-white">過去1週間でリプした回数: {reply.count}</p>
            <p className="text-white">親密度: {reply.intimacy}</p>
          </Flex>
        </Flex>
      ))}
    </Flex>
  );
};

export default ReplyList;
