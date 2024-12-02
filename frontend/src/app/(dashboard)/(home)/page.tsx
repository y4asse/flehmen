import { Windows } from "@/components/common/Windows";
import React from "react";
import { TimeGraph } from "./_components/TimeGraph";
import { WeekGraph } from "./_components/WeekGraph";
import { Habit } from "./_components/Habit";

const Page = () => {
  const windows = [
    {
      ...timeGraphWindow,
      title: "何の何時がひまかなー?？",
      children: <TimeGraph />,
    },
    {
      ...weekGraphWindow,
      title: "何に一番ついーとしてる？",
      children: <WeekGraph />,
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
