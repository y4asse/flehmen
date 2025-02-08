"use client";
import React, { useEffect, useState } from "react";
import {
  RadarChart,
  PolarGrid,
  PolarAngleAxis,
  PolarRadiusAxis,
  Radar,
  Tooltip,
} from "recharts";
import { Score } from "../page";

type Props = {
  score?: Score;
  isMobile?: boolean;
};

const FriendlyGraph = (props: Props) => {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  if (!isClient) return null; // サーバー側では何も描画しない

  if (!props.score) {
    return null;
  }
  const { isMobile } = props;
  const { balance, rhythm, time, type, words } = props.score;
  const data = [
    { subject: "時間帯", value: time },
    { subject: "バランス", value: balance },
    { subject: "テンポ", value: rhythm },
    { subject: "タイプ", value: type },
    { subject: "言葉", value: words },
  ];

  return (
    <div className="flex justify-center items-center">
      <RadarChart
        width={isMobile ? 465 : 600}
        height={isMobile ? 140 : 300}
        data={data}
        outerRadius={isMobile ? 50 : 100} // 五角形のサイズ
        margin={{ top: 40 }}
      >
        {/* 背景のグリッド */}
        <PolarGrid stroke="#eee" strokeDasharray="3 3" />

        {/* 軸のラベルの色 */}
        <PolarAngleAxis
          dataKey="subject"
          tick={{ fill: "#eee", fontSize: 14 }}
        />

        {/* 値の目盛り線 */}
        <PolarRadiusAxis tick={false} axisLine={false} tickLine={false} />

        {/* レーダーチャート本体 */}
        <Radar
          name="ステータス"
          dataKey="value"
          stroke="#e4007f" // 線の色
          fill="#e4007f" // 内部の塗りつぶし色
          fillOpacity={0.8} // 透明度
        />

        {/* ツールチップ */}
        <Tooltip
          wrapperStyle={{
            backgroundColor: "rgba(0, 0, 0, 0.7)",
            color: "#ffffff",
            borderRadius: "5px",
            padding: "5px",
          }}
        />
      </RadarChart>
    </div>
  );
};

export default FriendlyGraph;
