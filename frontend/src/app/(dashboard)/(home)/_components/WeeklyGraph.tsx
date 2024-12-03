"use client";
import React from "react";

import { PieChart, Pie, Cell, Tooltip, ResponsiveContainer } from "recharts";
import { WeeklyTweetCount } from "../page";

const COLORS = [
  "#E4007F",
  "#FFC2E2",
  "#FAA2D1",
  "#F681C1",
  "#F161B0",
  "#ED40A0",
  "#E81F8F",
];

type Props = {
  weeklyAllTweetCounts: WeeklyTweetCount[];
  dataKey: WeeklyTweetCountKey;
  nameKey: WeeklyTweetCountKey;
};

type WeeklyTweetCountKey = keyof WeeklyTweetCount;

export const WeeklyGraph = (props: Props) => {
  const { weeklyAllTweetCounts, dataKey, nameKey } = props;

  const renderCustomLabel = ({
    cx,
    cy,
    midAngle,
    innerRadius,
    outerRadius,
    percent,
    index,
  }: {
    cx: number;
    cy: number;
    midAngle: number;
    innerRadius: number;
    outerRadius: number;
    percent: number;
    index: number;
  }) => {
    // ラベルの表示位置を計算
    const RADIAN = Math.PI / 180;
    const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
    const x = cx + radius * Math.cos(-midAngle * RADIAN);
    const y = cy + radius * Math.sin(-midAngle * RADIAN);

    return (
      <text
        x={x}
        y={y}
        fill="#fff"
        textAnchor={x > cx ? "start" : "end"}
        dominantBaseline="central"
        fontSize={12}
      >
        {/* 改行する部分 */}
        <tspan x={x} dy="0" fill="#000">
          {weeklyAllTweetCounts[index].day}
        </tspan>
        <tspan fill="#000" x={x} dy="15">{`${(percent * 100).toFixed(
          0
        )}%`}</tspan>
      </text>
    );
  };

  return (
    <ResponsiveContainer width="100%" height="100%">
      <PieChart>
        {/* 円グラフ */}
        <Pie
          data={weeklyAllTweetCounts}
          dataKey={dataKey} // データ値に対応
          nameKey={nameKey} // 名前キー
          cx="50%" // 円のX位置
          cy="50%" // 円のY位置
          innerRadius={20} // 内側半径
          outerRadius={120} // 外側半径
          labelLine={false} // ラベルラインを非表示
          label={renderCustomLabel} // カスタムラベルを指定
        >
          {weeklyAllTweetCounts.map((_, index) => (
            <Cell
              key={`cell-${index}`}
              fill={COLORS[index % COLORS.length]}
              stroke="#000000"
            />
          ))}
        </Pie>

        {/* SVGでハート型を描画 */}
        <foreignObject
          x="50%" // 円の中心
          y="50%"
          width={75}
          height={75}
          style={{
            transform: "translate(-35px, -34px)",
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="black"
            width="75"
            height="75"
          >
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
          </svg>
        </foreignObject>

        {/* ツールチップ */}
        <Tooltip />
      </PieChart>
    </ResponsiveContainer>
  );
};
