import { Windows } from "@/components/common/Windows";
import React from "react";
import { TimeGraph } from "./_components/TimeGraph";
import { WeekGraph } from "./_components/WeekGraph";
import { Log } from "./_components/Log";

const Page = () => {
  const windows = [
    {
      ...timeGraphWindow,
      title: "何曜日の何時がひまかなー?？",
      children: <TimeGraph />,
    },
    {
      ...weekGraphWindow,
      title: "何曜日に一番ついーとしてる？",
      children: <WeekGraph />,
    },
    {
      ...logWindow,
      children: <Log />,
    },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

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
const logWindow = {
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
