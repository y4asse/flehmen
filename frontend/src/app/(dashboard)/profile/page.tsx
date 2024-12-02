import { Windows } from "@/components/common/Windows";
import React from "react";
import dayjs from "dayjs";
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

export type Sukipi = {
  name: string;
  weight?: number;
  height?: number;
  mbti?: MBTI;
  birthday?: string;
  hobby: string;
  shoeSize: number;
  famiry: string;
  nearyStation: string;
};

type SukipiLikedAt = {
  likedAt: string;
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
    weight: 0,
    height: 170,
    mbti: "INTP",
    birthday: "2003-02-18",
    hobby: "セクハラ",
    shoeSize: 27,
    famiry: "",
    nearyStation: "新宿",
  };

  const sukipiLikedAt: SukipiLikedAt = {
    likedAt: dayjs("2010-04-01").toISOString(), // 1月始まり
  };

  const windows = [
    {
      ...sukipiInfoWindow,
      title: `${sukipi.name}くんのこと`,
      children: <SukipiInfo sukipi={sukipi} />,
    },
    {
      ...sukipiLikeAtWindow,
      title: "好きになって何日？",
      children: <SukipiLikedAt likedAt={sukipiLikedAt.likedAt} />,
    },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

export default Page;