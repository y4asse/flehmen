import { Flex } from "@/components/ui/flex";
import { formatDateString } from "@/utils/date";
import React from "react";

type LogProps = {
  tweets: Tweet[];
};

type Tweet = {
  text: string;
  createdAt: string;
};

export const Log = (props: LogProps) => {
  const { tweets } = props;
  return (
    <Flex direction={"column"}>
      {tweets.map((tweet, i) => (
        <Flex key={i}>
          <p className="text-primary mr-2">
            {formatDateString(tweet.createdAt)}
          </p>
          <p className="text-primary">{tweet.text}</p>
        </Flex>
      ))}
    </Flex>
  );
};
