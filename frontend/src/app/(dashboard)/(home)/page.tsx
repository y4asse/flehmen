"use client";
import { Windows } from "@/components/common/Windows";
import React, { createContext } from "react";
import { MonthlyGraph } from "./_components/MonthlyGraph";
import { WeeklyGraph } from "./_components/WeeklyGraph";
import { Habit } from "./_components/Habit";
// import { MobileWindows } from "@/components/common/MobileWindows";
import { HomeIcon } from "./_components/HomeIcon";
// import * as $axios from "@/lib/axios";

import { Autoplay, Navigation, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";

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

export type Icon = {
  name: string;
  image: string;
  href: string;
};

const Page = () => {
  // const result = $axios.request({
  //   url: "/pet/findByTags",
  //   method: "get",
  //   params: {
  //     tags: ["available"],
  //   },
  // });
  const windows = [
    {
      ...habitWindow,
      title: "生息時間",
      children: <Habit busy_color_index_list={busy_color_index_list} />,
    },
    {
      ...monthlyGraphWindow,
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
  ];

  const mobile = [
    {
      title: "生息時間",
      children: <Habit busy_color_index_list={busy_color_index_list} />,
    },
    {
      title: "いつひまだったの",
      children: <MonthlyGraph monthlyTweetCounts={monthlyTweetCounts} />,
    },
    {
      title: "なんようにひまなんだろ",
      children: (
        <WeeklyGraph
          weeklyAllTweetCounts={weeklyAllTweetCounts}
          nameKey="day"
          dataKey="count"
        />
      ),
    },
  ];

  return (
    <div>
      {/* PC用 */}
      <div className="hidden md:block">
        <Windows windows={windows} />
      </div>

      {/* スマホ用 */}
      <div className="block md:hidden ">
        {/* アイコン */}
        <HomeIcon Icon={Icon} />
        {/* グラフ */}
        <Swiper
          style={{
            position: "absolute",
            bottom: "100px",
            left: "0",
            width: "100%",
            height: "300px",
          }}
          // modules={[Autoplay, Navigation, Pagination]}
          // navigation
          // pagination
        >
          {mobile.map((mobile, index) => (
            <SwiperSlide
              key={index}
              style={{
                width: "100%",
                height: "300px",
                padding: "0 50px",
              }}
            >
              {mobile.children}
            </SwiperSlide>
          ))}
        </Swiper>
        {/* <div
          className="flex flex-col gap-4"
          style={{ width: "300px", height: "300px" }}
        >
          <WeeklyGraph
            weeklyAllTweetCounts={weeklyAllTweetCounts}
            nameKey="day"
            dataKey="count"
          />
        </div> */}
      </div>
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
    height: 530,
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

const Icon: Icon[] = [
  {
    name: "プロフィール",
    image: "/images/file.svg",
    href: "/profile",
  },
  {
    name: "FF",
    image: "/images/file.svg",
    href: "/ff",
  },
  {
    name: "リプ",
    image: "/images/file.svg",
    href: "/reply",
  },
  {
    name: "ひみつ",
    image: "/images/file.svg",
    href: "/secret",
  },
  {
    name: "会話",
    image: "/images/file.svg",
    href: "/",
  },
];

export default Page;
