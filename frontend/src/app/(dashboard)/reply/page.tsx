import { Windows } from "@/components/common/Windows";
import { title } from "process";
import React from "react";
import ReplyList from "./_components/ReplyList";
import { Sukipi } from "../profile/page";
import Contents from "./_components/Contents";

export type reply = {
  id: number;
  count: number;
  intimacy: number;
};

export type content = {
  id: number;
  content: string;
};

const Page = () => {
  const windows = [
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
    // {
    //   ...contents,
    //   title: `内容`,
    //   children: <ReplyList replyContent={replyContent} />,
    // },
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

const sukipi: Sukipi = {
  name: "早瀬",
};

export default Page;
