import { Windows } from "@/components/common/Windows";
import React from "react";
import { SukipiInfo } from "./_components/SukipiInfo";
import { SukipiLikedAt } from "./_components/SukipiLikedAt";

type MBTI =
  | "ENFP"
  | "ENFJ"
  | "ENTP"
  | "ENTJ"
  | "ESFP"
  | "ESFJ"
  | "ESTP"
  | "ESTJ"
  | "INFP"
  | "INFJ"
  | "INTP"
  | "INTJ"
  | "ISFP"
  | "ISFJ"
  | "ISTP"
  | "ISTJ";

type Sukipi = {
  name: string;
  weight?: number;
  height?: number;
  mbti?: MBTI;
};

const Page = () => {
  const sukipiInfoWindow = {
    initSize: {
      width: 750,
      height: 500,
    },
    initPosition: {
      x: 650,
      y: 50,
      z: 2,
    },
  };

  const sukipiLikeAtWindow = {
    initSize: {
      width: 500,
      height: 400,
    },
    initPosition: {
      x: 300,
      y: 250,
      z: 1,
    },
  };

  const sukipi: Sukipi = {
    name: "早瀬",
    weight: 60,
    height: 170,
    mbti: "ENFP",
  };

  const windows = [
    {
      ...sukipiInfoWindow,
      title: `${sukipi.name}くんのこと`,
      children: <SukipiInfo />,
    },
    { ...sukipiLikeAtWindow, children: <SukipiLikedAt /> },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

export default Page;
