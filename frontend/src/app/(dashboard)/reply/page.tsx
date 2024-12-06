import { Windows } from "@/components/common/Windows";
import { title } from "process";
import React from "react";
import ReplyList from "./_components/ReplyList";
import { Sukipi } from "../profile/page";
import Contents from "./_components/Contents";
import { MonthlyRep } from "./_components/MonthlyRep";

export type reply = {
  id: number;
  count: number;
  intimacy: number;
};

export type content = {
  id: number;
  content: string;
};

export type monthly = {
  date: string;
  count: number;
};

const Page = () => {
  const windows = [
    {
      ...monthlRep,
      title: `過去１ヶ月の${sukipi.name}くんのリプの絡み`,
      children: <MonthlyRep GraphInfo={GraphInfo} />,
    },
    {
      ...replyList,
      title: `${sukipi.name}くんとリプしてる人リスト`,
      children: <ReplyList replyInfo={replyInfo} />,
    },
    {
      ...contents,
      title: "内容",
      children: <Contents repContents={replyContent} />,
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
    count: 9,
    intimacy: 4,
  },
  {
    id: 2,
    count: 1,
    intimacy: 1,
  },
  {
    id: 1,
    count: 9,
    intimacy: 4,
  },
  {
    id: 2,
    count: 1,
    intimacy: 1,
  },
  {
    id: 1,
    count: 9,
    intimacy: 4,
  },
  {
    id: 2,
    count: 1,
    intimacy: 1,
  },
];

const replyContent: content[] = [
  { id: 1, content: "おはよう" },
  { id: 2, content: "こんにちは" },
  { id: 3, content: "こんばんは" },
  {
    id: 4,
    content:
      "やばい課題おわってないよ〜！誰かこの問題おしえてくれ〜😭教えてくだれらご飯奢る！",
  },
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
    x: 150,
    y: 300,
    z: 2,
  },
};

const contents = {
  initSize: { width: 600, height: 400 },
  initPosition: {
    x: 650,
    y: 400,
    z: 3,
  },
};

const monthlRep = {
  initSize: { width: 800, height: 600 },
  initPosition: {
    x: 800,
    y: 40,
    z: 1,
  },
};
const sukipi: Sukipi = {
  name: "早瀬",
};

export default Page;
