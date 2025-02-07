"use client";
import { MobileWindows } from "@/components/common/MobileWindows";
import React from "react";
import { Habit } from "../(home)/_components/Habit";
import { MonthlyGraph } from "../(home)/_components/MonthlyGraph";
import { WeeklyGraph } from "../(home)/_components/WeeklyGraph";
import { DateTweetCount, WeeklyTweetCount } from "../(home)/page";

const Page = () => {
  const windows = [
    {
      ...monthlyGraphWindow,
      title: "いつひまだったの",
      children: (
        <MonthlyGraph monthlyTweetCounts={monthlyTweetCounts} isMobile />
      ),
    },
    {
      ...habitWindow,
      title: "生息時間",
      children: (
        <Habit busy_color_index_list={busy_color_index_list} isMobile />
      ),
    },
    {
      ...weekGraphWindow,
      title: "なんようにひまなんだろ",
      children: (
        <WeeklyGraph
          weeklyAllTweetCounts={weeklyAllTweetCounts}
          nameKey="day"
          dataKey="count"
          isMobile
        />
      ),
    },
  ];
  return (
    <div>
      <div className="hidden md:block">
        <p className="text-white absolute z-20 top-[50vh] left-[40vw] text-2xl">
          スマホで見てね！
        </p>
      </div>
      <div className="block md:hidden">
        <MobileWindows windows={windows} />
      </div>
    </div>
  );
};

const monthlyTweetCounts: DateTweetCount[] = [
  { date: "11/01", count: 2 },
  { date: "11/03", count: 5 },
  { date: "11/05", count: 4 },
  { date: "11/07", count: 3 },
  { date: "11/09", count: 0 },
  { date: "11/11", count: 4 },
  { date: "11/13", count: 3 },
  { date: "11/15", count: 0 },
  { date: "11/17", count: 4 },
  { date: "11/19", count: 3 },
  { date: "11/21", count: 0 },
  { date: "11/23", count: 4 },
  { date: "11/25", count: 3 },
  { date: "11/27", count: 0 },
  { date: "11/29", count: 4 },
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
  [2, 2, 3, 3, 4, 0, 1],
  [3, 5, 4, 5, 0, 3, 4],
  [4, 5, 5, 1, 1, 2, 3],
  [1, 2, 0, 2, 1, 0, 0],

  [2, 5, 1, 5, 2, 3, 5],
  [0, 2, 2, 2, 3, 2, 0],
  [1, 4, 3, 4, 5, 4, 4],
  [1, 1, 1, 1, 2, 3, 2],
  [3, 4, 4, 5, 5, 1, 2],

  [4, 5, 0, 1, 2, 3, 4],
  [3, 3, 0, 1, 2, 1, 1],
  [3, 4, 0, 2, 3, 5, 5],
  [3, 5, 0, 2, 3, 3, 4],
  [1, 0, 0, 2, 3, 3, 1],

  [4, 1, 5, 0, 0, 1, 2],
  [0, 2, 0, 1, 1, 1, 1],
  [0, 2, 5, 3, 3, 4, 1],
  [0, 4, 0, 3, 4, 5, 0],
  [0, 5, 1, 4, 5, 0, 1],

  [2, 3, 4, 2, 0, 1, 5],
  [1, 1, 1, 0, 1, 3, 4],
  [5, 5, 0, 1, 1, 2, 3],
  [5, 1, 1, 3, 3, 1, 5],
  [1, 2, 0, 1, 0, 3, 0],
];

const monthlyGraphWindow = {
  initSize: {
    width: 800,
    height: "calc( (100vh - 170px ) / 3 )",
  },
  initPosition: {
    x: 140,
    y: 80,
    z: 1,
  },
};

const weekGraphWindow = {
  initSize: {
    width: 430,
    height: "calc( (100vh - 170px ) / 3 )",
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
    height: "calc( (100vh - 170px ) / 3 )",
  },
  initPosition: {
    x: 900,
    y: 350,
    z: 3,
  },
};

export default Page;
