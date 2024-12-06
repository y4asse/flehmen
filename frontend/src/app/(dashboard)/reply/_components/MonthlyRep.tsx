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
};

export const MonthlyRep = (props: Props) => {
  const { GraphInfo } = props;
  return (
    <ResponsiveContainer width="100%" height={300}>
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
