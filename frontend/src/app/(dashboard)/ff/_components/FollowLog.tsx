import React from "react";
import { Follow } from "../page";
import { Flex } from "@/components/ui/flex";
import Image from "next/image";
import { formatDateString } from "@/utils/date";

type Props = {
  followList: Follow[];
};

export const FollowLog = (props: Props) => {
  const { followList } = props;
  return (
    <Flex
      direction={"column"}
      justify={"start"}
      align={"start"}
      className="gap-4 ~mt-3/12 mb-3 w-full ~px-4/20"
    >
      {followList.map((follow) => (
        <a
          href={`https://twitter.com/${follow.id}`}
          target="_blank"
          key={follow.id}
          className="hover:opacity-80 w-full"
        >
          <Flex className="gap-2" justify={"between"}>
            <Flex className="gap-2">
              <Image
                src={follow.icon}
                alt={follow.name}
                width={50}
                height={50}
                className="rounded-full"
              />
              <Flex direction="column" align={"start"}>
                <p className="text-[#ccc] text-sm">
                  {formatDateString(follow.followed_at)}
                </p>
                <Flex className="~gap-2/4">
                  <p className="text-white truncate ~max-w-28/96">
                    {follow.name}
                  </p>
                  <p className="text-[#ccc] truncate ~max-w-28/96">
                    {follow.bio}
                  </p>
                </Flex>
              </Flex>
            </Flex>
            <p className="bg-primary rounded-md p-2">とぶ</p>
          </Flex>
        </a>
      ))}
    </Flex>
  );
};
