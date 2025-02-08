"use client";
import React from "react";
import {
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Bar,
  ComposedChart,
} from "recharts";
import { monthly } from "../page";

type Props = {
  GraphInfo: monthly[];
  isMobile?: boolean;
};

export const MonthlyRep = (props: Props) => {
  const { GraphInfo, isMobile } = props;
  return (
    <ResponsiveContainer
      width="100%"
      height={isMobile ? 130 : 300}
      className={"relative right-7"}
    >
      <ComposedChart data={GraphInfo}>
        <CartesianGrid strokeDasharray="5 5" stroke="#E4007F" />
        <XAxis dataKey="date" stroke="#E4007F" />
        <YAxis stroke="#E4007F" />
        <Tooltip />
        <Bar dataKey="count" fill="#E4007F"></Bar>
      </ComposedChart>
    </ResponsiveContainer>
  );
};
