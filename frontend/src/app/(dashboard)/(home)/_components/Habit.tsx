import React from "react";

type LogProps = {
  busy_color_index_list: number[][];
};

export const Habit = (props: LogProps) => {
  const { busy_color_index_list } = props;
  const days = ["げつ", "にち", "どー", "きん", "もく", "すい", "かー"];
  const hours = [
    "00:00",
    "01:00",
    "02:00",
    "03:00",
    "04:00",
    "05:00",
    "06:00",
    "07:00",
    "08:00",
    "09:00",
    "10:00",
    "11:00",
    "12:00",
    "13:00",
    "14:00",
    "15:00",
    "16:00",
    "17:00",
    "18:00",
    "19:00",
    "20:00",
    "21:00",
    "22:00",
    "23:00",
    "24:00",
  ];
  const busy_color = [0, 0.2, 0.4, 0.6, 0.8, 1];
  return (
    <div
      style={{ display: "flex", flexDirection: "column", width: "100%" }}
      className="overflow-y-scroll max-h-full"
    >
      {/* Header Row */}
      <div
        style={{
          display: "flex",
          backgroundColor: "#000",
          fontWeight: "bold",
        }}
      >
        <div
          style={{ width: "80px", padding: "8px", textAlign: "center" }}
        ></div>
        {days.map((day) => (
          <div
            key={day}
            style={{
              flex: 1,
              padding: "8px",
              textAlign: "center",
            }}
          >
            <p className="text-primary">{day}</p>
          </div>
        ))}
      </div>
      {/* Data Rows */}
      {hours.map((hour) => (
        <div key={hour} style={{ display: "flex" }}>
          {/* Hour Column */}
          <div
            style={{
              width: "80px",
              padding: "8px",
              textAlign: "center",
              backgroundColor: "#000",
            }}
            className="text-primary"
          >
            {hour}
          </div>
          {/* Data Columns */}
          {busy_color_index_list.map((busy_color_index, i) => (
            <div
              key={i}
              style={{
                flex: 1,
                padding: "8px",
                textAlign: "center",
                backgroundColor: "#E4007F",
                opacity: busy_color[busy_color_index[i]],
              }}
            >
              {busy_color[busy_color_index[i]]}
            </div>
          ))}
        </div>
      ))}
    </div>
  );
};
