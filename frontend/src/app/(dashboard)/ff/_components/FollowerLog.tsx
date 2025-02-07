import { Flex } from "@/components/ui/flex";
import { formatDateString } from "@/utils/date";
import React from "react";
import Image from "next/image";
import { Follower } from "../page";

type Props = {
  followerList: Follower[];
};

export const FollowerLog = (props: Props) => {
  const { followerList } = props;
  return (
    <Flex
      direction={"column"}
      justify={"start"}
      align={"start"}
      className="gap-4 ~mt-3/12 mb-3 w-full ~px-2/20"
    >
      {followerList.map((follower) => (
        <a
          href={`https://twitter.com/${follower.id}`}
          target="_blank"
          key={follower.id}
          className="hover:opacity-80 w-full"
        >
          <Flex className="gap-2" justify={"between"}>
            <Flex className="gap-2">
              <Image
                src={follower.icon}
                alt={follower.name}
                width={50}
                height={50}
                className="rounded-full"
              />
              <Flex direction="column" align={"start"}>
                <p className="text-[#ccc] text-sm">
                  {formatDateString(follower.followed_at)}
                </p>
                <Flex className="gap-4">
                  <p className="text-white truncate ~max-w-28/96e">
                    {follower.name}
                  </p>
                  <p className="text-[#ccc] truncate ~max-w-28/96">
                    {follower.bio}
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
