import { Windows } from "@/components/common/Windows";
import React from "react";
import { MonthlyGraph } from "./_components/MonthlyGraph";
import { WeeklyGraph } from "./_components/WeeklyGraph";
import { Habit } from "./_components/Habit";

export type DateTweetCount = {
  date: string;
  count: number;
};

export type TweetDayCount = {
  day: string;
  count: number;
};

export type WeeklyTweetCount = {
  day: "げつ" | "にち" | "どー" | "きん" | "もく" | "すい" | "かー";
  count: number;
};

const Page = () => {
  const windows = [
    {
      ...timeGraphWindow,
      title: "いつひまだったの",
      children: <MonthlyGraph monthlyTweetCounts={monthlyTweetCounts} />,
    },
    {
      ...weekGraphWindow,
      title: "なんようにひまなんだろ",
      children: (
        <WeeklyGraph
          weeklyAllTweetCounts={weeklyAllTweetCounts}
          nameKey="day"
          dataKey="count"
        />
      ),
    },
    {
      ...habitWindow,
      title: "生息時間",
      children: <Habit busy_color_index_list={busy_color_index_list} />,
    },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

const monthlyTweetCounts: DateTweetCount[] = [
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

const weeklyAllTweetCounts: WeeklyTweetCount[] = [
  { day: "げつ", count: 400 },
  { day: "にち", count: 239 },
  { day: "どー", count: 189 },
  { day: "きん", count: 278 },
  { day: "もく", count: 200 },
  { day: "すい", count: 300 },
  { day: "かー", count: 300 },
];

const busy_color_index_list = [
  [0, 1, 2, 3, 4, 5, 0],
  [1, 2, 3, 4, 5, 0, 1],
  [2, 3, 4, 5, 0, 1, 2],
  [3, 4, 5, 0, 1, 2, 3],
  [4, 5, 0, 1, 2, 3, 4],

  [5, 0, 1, 2, 3, 4, 5],
  [0, 1, 2, 3, 4, 5, 0],
  [1, 2, 3, 4, 5, 0, 1],
  [2, 3, 4, 5, 0, 1, 2],
  [3, 4, 5, 0, 1, 2, 3],

  [4, 5, 0, 1, 2, 3, 4],
  [5, 0, 1, 2, 3, 4, 5],
  [0, 1, 2, 3, 4, 5, 0],
  [1, 2, 3, 4, 5, 0, 1],
  [2, 3, 4, 5, 0, 1, 2],

  [3, 4, 5, 0, 1, 2, 3],
  [4, 5, 0, 1, 2, 3, 4],
  [5, 0, 1, 2, 3, 4, 5],
  [0, 1, 2, 3, 4, 5, 0],
  [1, 2, 3, 4, 5, 0, 1],

  [2, 3, 4, 5, 0, 1, 2],
  [3, 4, 5, 0, 1, 2, 3],
  [4, 5, 0, 1, 2, 3, 4],
  [5, 0, 1, 2, 3, 4, 5],
  [4, 5, 0, 1, 2, 3, 4],
];

const timeGraphWindow = {
  initSize: {
    width: 800,
    height: 530,
  },
  initPosition: {
    x: 170,
    y: 80,
    z: 1,
  },
};

const weekGraphWindow = {
  initSize: {
    width: 430,
    height: 330,
  },
  initPosition: {
    x: 950,
    y: 50,
    z: 2,
  },
};
const habitWindow = {
  initSize: {
    width: 460,
    height: 330,
  },
  initPosition: {
    x: 900,
    y: 350,
    z: 3,
  },
};

export default Page;
