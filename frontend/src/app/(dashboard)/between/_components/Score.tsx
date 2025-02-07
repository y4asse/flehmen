import React from "react";
import Image from "next/image";
import { Flex } from "@/components/ui/flex";

type ScoreProps = {
  total?: number;
};

export const Score = (props: ScoreProps) => {
  if (!props.total) {
    return null;
  }
  const { total } = props;
  const brainCount = total / 16;
  return (
    <Flex direction={"column"} className="gap-2">
      <p className="text-white">今の親密度は...</p>
      <p className="text-white text-4xl">{total}</p>
      <Flex className="gap-6">
        <ScoreHeart isLeft={brainCount >= 0} isRight={brainCount >= 2} />
        <ScoreHeart isLeft={brainCount >= 3} isRight={brainCount >= 4} />
        <ScoreHeart isLeft={brainCount >= 5} isRight={brainCount >= 6} />
      </Flex>
    </Flex>
  );
};

type ScoreHeartProps = {
  isLeft: boolean;
  isRight: boolean;
};

const ScoreHeart = (props: ScoreHeartProps) => {
  const { isLeft, isRight } = props;
  return (
    <Flex>
      <Image
        src={
          isLeft
            ? "/images/score_brain_left_active.svg"
            : "/images/score_brain_left.svg"
        }
        width={30}
        height={40}
        className="object-contain"
        alt="brain"
      />
      <Image
        src={
          isRight
            ? "/images/score_brain_right_active.svg"
            : "/images/score_brain_right.svg"
        }
        width={30}
        height={40}
        className="object-contain"
        alt="brain"
      />
    </Flex>
  );
};
