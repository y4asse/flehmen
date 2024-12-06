import { Windows } from "@/components/common/Windows";
import React from "react";
import ReplyList from "./_components/ReplyList";
import { Sukipi } from "../profile/page";
import Contents from "./_components/Contents";
import { MonthlyRep } from "./_components/MonthlyRep";

export type reply = {
  id: string;
  count: number;
  intimacy: number;
};

export type content = {
  id: string;
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
    id: "dinsei_",
    count: 9,
    intimacy: 4,
  },
  {
    id: "kipeo22",
    count: 1,
    intimacy: 1,
  },
  {
    id: "dinsei_",
    count: 9,
    intimacy: 4,
  },
  {
    id: "kipeo22",
    count: 1,
    intimacy: 1,
  },
  {
    id: "dinsei_",
    count: 9,
    intimacy: 4,
  },
  {
    id: "kipeo22",
    count: 1,
    intimacy: 1,
  },
];

const replyContent: content[] = [
  { id: "dinsei_", content: "n競技プログラミングごみ" },
  { id: "kipeo22", content: "おもろいバイトしたい" },
  { id: "adeam02", content: "気持ちめちゃくちゃわかるで" },
  {
    id: "7Rqebz",
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
  initSize: { width: 800, height: 400 },
  initPosition: {
    x: 600,
    y: 40,
    z: 1,
  },
};
const sukipi: Sukipi = {
  name: "早瀬",
};

export default Page;
