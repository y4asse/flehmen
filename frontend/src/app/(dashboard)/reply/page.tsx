"use client";
import { Windows } from "@/components/common/Windows";
import React from "react";
import ReplyList from "./_components/ReplyList";
import { Sukipi } from "../profile/page";
import Contents from "./_components/Contents";
import { MonthlyRep } from "./_components/MonthlyRep";

export type reply = {
  id: number;
  userId: string;
  name: string;
  count: number;
  intimacy: number;
  icon: string;
};

export type content = {
  id: number;
  userId: string;
  name: string;
  icon: string;
  value: string;
};

export type monthly = {
  date: string;
  count: number;
};

const Page = () => {
  const [userIndex, setUserIndex] = React.useState(1);
  const handleSetUserIndex = (index: number) => {
    setUserIndex(index);
  };
  const windows = [
    {
      ...monthlRep,
      title: `過去１ヶ月の${sukipi.name}くんのリプの絡み`,
      children: <MonthlyRep GraphInfo={GraphInfo} />,
    },
    {
      ...replyList,
      title: `${sukipi.name}くんとリプしてる人リスト`,
      children: (
        <ReplyList
          replyInfo={replyInfo}
          handleSetUserIndex={handleSetUserIndex}
        />
      ),
    },
    {
      ...contents,
      title: "内容",
      children: <Contents repContents={replyContent} userIndex={userIndex} />,
    },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

const replyInfo: reply[] = [
  {
    id: 1,
    icon: "/images/karin.png",
    count: 9,
    intimacy: 4,
    name: "かりんたん",
    userId: "dinsei_",
  },
  {
    id: 2,
    icon: "https://picsum.photos/200",
    count: 5,
    intimacy: 3,
    name: "ささandりょう",
    userId: "sasa",
  },
  {
    id: 3,
    icon: "https://picsum.photos/300",
    count: 3,
    intimacy: 2,
    name: "まえひろ",
    userId: "maehiro",
  },
  {
    id: 4,
    icon: "https://picsum.photos/400",
    count: 1,
    intimacy: 1,
    name: "たけし",
    userId: "takeshi",
  },
];

const replyContent: content[][] = [
  [
    {
      id: 1,
      icon: "/images/karin.png",
      name: "かりんたん",
      userId: "dinsei_",
      value: "課題どうしよ",
    },
    {
      id: 2,
      icon: "https://picsum.photos/500",
      name: "早瀬",
      userId: "yase",
      value: "明日までだよー",
    },
    {
      id: 3,
      icon: "/images/karin.png",
      name: "かりんたん",
      userId: "dinsei_",
      value: "ありがとう〜やった？",
    },
    {
      id: 4,
      icon: "https://picsum.photos/200",
      name: "早瀬",
      userId: "yase",
      value: "もち",
    },
  ],
  [
    {
      id: 3,
      icon: "https://picsum.photos/500",
      name: "早瀬",
      userId: "yase",
      value: "午後からバスケできる人おる？",
    },
    {
      id: 4,
      icon: "https://picsum.photos/200",
      name: "ささandりょう",
      userId: "yase",
      value: "行けるべ",
    },
  ],
];

const GraphInfo: monthly[] = [
  { date: "11/01", count: 2 },
  { date: "11/02", count: 0 },
  { date: "11/03", count: 5 },
  { date: "11/04", count: 1 },
  { date: "11/05", count: 4 },
  { date: "11/06", count: 0 },
  { date: "11/07", count: 3 },
  { date: "11/08", count: 2 },
  { date: "11/09", count: 0 },
  { date: "11/10", count: 1 },
  { date: "11/11", count: 4 },
  { date: "11/12", count: 0 },
  { date: "11/13", count: 3 },
  { date: "11/14", count: 2 },
  { date: "11/15", count: 0 },
  { date: "11/16", count: 1 },
  { date: "11/17", count: 4 },
  { date: "11/18", count: 0 },
  { date: "11/19", count: 3 },
  { date: "11/20", count: 2 },
  { date: "11/21", count: 0 },
  { date: "11/22", count: 1 },
  { date: "11/23", count: 4 },
  { date: "11/24", count: 0 },
  { date: "11/25", count: 3 },
  { date: "11/26", count: 2 },
  { date: "11/27", count: 0 },
  { date: "11/28", count: 1 },
  { date: "11/29", count: 4 },
  { date: "11/30", count: 0 },
];

const replyList = {
  initSize: { width: 600, height: 400 },
  initPosition: {
    x: 175,
    y: 300,
    z: 2,
  },
};

const contents = {
  initSize: { width: 600, height: 350 },
  initPosition: {
    x: 750,
    y: 400,
    z: 3,
  },
};

const monthlRep = {
  initSize: { width: 650, height: 500 },
  initPosition: {
    x: 730,
    y: 40,
    z: 1,
  },
};
const sukipi: Sukipi = {
  name: "早瀬",
};

export default Page;
