import React from "react";
import { title } from "process";

// 情報アイテムのリスト
const infoItems = [
  { title: "なまえ", content: "" },
  { title: "身長", content: "" },
  { title: "体重", content: "" },
  { title: "学校", content: "" },
  { title: "クラス", content: "" },
  { title: "ねんれい", content: "" },
  { title: "たんじょうび", content: "" },
  { title: "しゅみ", content: "" },
  { title: "すきなたべもの", content: "" },
  { title: "すきなうた", content: "" },
  { title: "すきなほん", content: "" },
  { title: "すきなえいが", content: "" },
  { title: "すきなアニメ", content: "" },
  { title: "すきなゲーム", content: "" }
];

// Propsの型定義
type Props = { 
  title: string;
  content: string;
};

// メインコンポーネント
export const SukipiInfo = () => {
  return (
    <div>
      {infoItems.map((item) => (
        <InfoItem key={item.title} title={item.title} content={item.content} />
      ))}
    </div>
  );
};

// 情報項目コンポーネント
const InfoItem = ({ title, content }: Props) => {
  return (
    <div style={styles.infoItem}>
      <div style={styles.title}>{title}</div>
      <div style={styles.content}>{content}</div>
    </div>
  );
};

// スタイルの定義
const styles: { [key: string]: React.CSSProperties } = {
  infoItem: {
    display: "flex",
    margin: "10px 0",
    fontSize: "0.8em",
    fontWeight: "bold",
    color: "white",
    fontFamily: '"PixelMplus10"',
  },
};

export default SukipiInfo;
