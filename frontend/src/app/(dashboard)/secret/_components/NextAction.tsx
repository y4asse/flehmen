import { Button } from "@/components/ui/button";
import { Flex } from "@/components/ui/flex";
import React from "react";

type Props = {
  text?: string;
  onClickBack: () => void;
};
const NextAction = (props: Props) => {
  const { text, onClickBack } = props;

  return (
    <Flex direction={"column"} className="">
      <p className="text-white p-4">{text}</p>
      <Button onClick={onClickBack}>もどる</Button>
    </Flex>
  );
};

export default NextAction;
