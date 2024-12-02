import React from "react";

type LogProps = {
  busy_color_index_list: number[][];
};

export const Habit = (props: LogProps) => {
  const { busy_color_index_list } = props;
  const days = ["げつ", "か-", "すい", "もく", "きん", "ど", "にち"];
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
  // 現在時刻の時間帯を取得
  const currentHour = new Date().getHours(); // 0〜23 の時間
  const currentHourIndex = hours.findIndex((hour) =>
    hour.startsWith(currentHour.toString().padStart(2, "0"))
  );
  // 現在の曜日を取得
  const currentDayIndex = (new Date().getDay() + 6) % 7; // 0(日曜)〜6(土曜)

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
              textAlign: "center",
            }}
          >
            <p className="text-primary">{day}</p>
          </div>
        ))}
      </div>
      {/* Data Rows */}
      {hours.map((hour, hourIndex) => (
        <div key={hour} style={{ display: "flex" }}>
          {/* Hour Column */}
          <div
            style={{
              width: "80px",
              textAlign: "center",
              backgroundColor: "#000",
            }}
            className="text-primary"
          >
            {hour}
          </div>
          {/* Data Columns */}
          {days.map((_, dayIndex) => {
            const colorIndex =
              busy_color_index_list[hourIndex]?.[dayIndex] ?? 0; // 安全に取得
            return (
              <div
                key={dayIndex}
                style={{
                  flex: 1,
                  textAlign: "center",
                  backgroundColor: "#E4007F",
                  opacity: busy_color[colorIndex],
                  border:
                    hourIndex === currentHourIndex &&
                    dayIndex === currentDayIndex
                      ? "2px solid #fff"
                      : "none",
                }}
              ></div>
            );
          })}
        </div>
      ))}
    </div>
  );
};
