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
    <Flex direction={"column"} className="~gap-2/6">
      <p className="text-white ~text-base/2xl">今の親密度は...</p>
      <p className="text-white ~text-4xl/7xl">{total}</p>
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
        width={40}
        height={40}
        className="object-contain ~w-8/16 ~h-10/20"
        alt="brain"
      />
      <Image
        src={
          isRight
            ? "/images/score_brain_right_active.svg"
            : "/images/score_brain_right.svg"
        }
        width={40}
        height={40}
        className="object-contain ~w-8/16 ~h-10/20"
        alt="brain"
      />
    </Flex>
  );
};
