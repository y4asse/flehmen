"use client";
import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { DateTweetCount } from "../page";

interface CustomDotProps {
  cx?: number; // ドットの中心 x 座標
  cy?: number; // ドットの中心 y 座標
  stroke?: string; // 線の色
}

type Props = {
  monthlyTweetCounts: DateTweetCount[];
  isMobile?: boolean;
};

const HeartDot: React.FC<CustomDotProps> = ({ cx, cy, stroke }) => {
  if (!cx || !cy) return null; // cx, cy が undefined の場合は何も描画しない

  return (
    <svg
      x={cx - 10}
      y={cy - 10}
      width={20}
      height={20}
      viewBox="0 0 24 24"
      fill="white"
      stroke={stroke || "pink"}
      strokeWidth={2}
    >
      <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
    </svg>
  );
};

export const MonthlyGraph = (props: Props) => {
  const { monthlyTweetCounts, isMobile } = props;
  return (
    <ResponsiveContainer
      width="100%"
      height={isMobile ? 130 : 300}
      className={"relative right-7"}
    >
      <LineChart data={monthlyTweetCounts}>
        <CartesianGrid strokeDasharray="5 5" stroke="#E4007F" />
        <XAxis dataKey="date" stroke="#E4007F" />
        <YAxis stroke="#E4007F" />
        <Tooltip />
        <Line
          type="monotone"
          strokeWidth={3}
          dataKey="count"
          stroke="#E4007F"
          dot={<HeartDot />}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};
