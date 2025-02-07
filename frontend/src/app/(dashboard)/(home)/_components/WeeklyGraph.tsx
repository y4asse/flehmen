"use client";
import React, { useEffect, useState } from "react";
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
  isMobile?: boolean;
};

type WeeklyTweetCountKey = keyof WeeklyTweetCount;

export const WeeklyGraph = (props: Props) => {
  const { weeklyAllTweetCounts, dataKey, nameKey, isMobile } = props;
  const [radius, setRadius] = useState({ inner: 0, outer: 120 });
  const [heartSize, setHeartSize] = useState(50);
  const [fontSize, setFontSize] = useState(12);

  useEffect(() => {
    const updateSizes = () => {
      const width = window.innerWidth;
      const height = window.innerHeight;
      const minSize = Math.min(width, height);

      setRadius((prev) => ({
        ...prev,
        outer: minSize * 0.17,
      }));

      setHeartSize(minSize * 0.08); // ハートのサイズを調整
      setFontSize(Math.max(10, minSize * 0.015)); // 文字サイズを調整（最小10px）
    };

    updateSizes();
    window.addEventListener("resize", updateSizes);
    return () => window.removeEventListener("resize", updateSizes);
  }, []);

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
        fill="#000"
        textAnchor={x > cx ? "start" : "end"}
        dominantBaseline="central"
        fontSize={fontSize}
      >
        <tspan x={x} dy="0">
          {weeklyAllTweetCounts[index].day}
        </tspan>
        <tspan x={x} dy={fontSize * 1.5}>{`${(percent * 100).toFixed(
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
          innerRadius={radius.inner}
          outerRadius={radius.outer}
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
          width={heartSize}
          height={heartSize}
          style={{
            transform: `translate(-${heartSize / 2}px, -${heartSize / 2}px)`,
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="black"
            width={heartSize}
            height={heartSize}
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
