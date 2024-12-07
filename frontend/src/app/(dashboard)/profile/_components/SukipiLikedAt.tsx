import React from "react";
import dayjs from "dayjs";

// Propsの型定義
type Props = {
  likedAt: string;
};

export const SukipiLikedAt = ({ likedAt }: Props) => {
  // 日付操作を dayjs で行う
  const likedAtDate = dayjs(likedAt);
  const today = dayjs();

  // 経過日数を計算
  const elapsedDays = today.diff(likedAtDate, "day");

  return (
    <div style={styles.container}>
      <div>好きになって</div>
      <div style={styles.text}>{elapsedDays}日</div>
    </div>
  );
};

const styles: { [key: string]: React.CSSProperties } = {
  container: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    height: "100%",
    color: "white",
    fontSize: "1.2em",
  },
  text: {
    fontWeight: "bold",
  },
};

export default SukipiLikedAt;
