import React from "react";

// Propsの型定義
type Sukipi = {
  name: string;
  weight?: number;
  height?: number;
  mbti?: string;
};

type Props = {
  sukipi: Sukipi;
};

// メインコンポーネント
export const SukipiInfo = ({ sukipi }: Props) => {
  // Sukipiオブジェクトから情報を作成
  const infoItems = [
    { title: "なまえ", content: sukipi.name },
    { title: "身長", content: sukipi.height ? `${sukipi.height} cm` : "未設定" },
    { title: "体重", content: sukipi.weight ? `${sukipi.weight} kg` : "未設定" },
    { title: "MBTI", content: sukipi.mbti || "未設定" },
  ];

  return (
    <div>
      {infoItems.map((item) => (
        <InfoItem key={item.title} title={item.title} content={item.content} />
      ))}
    </div>
  );
};

// 情報項目コンポーネント
type InfoItemProps = {
  title: string;
  content: string;
};

const InfoItem = ({ title, content }: InfoItemProps) => {
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
  title: {
    width: "100px",
  },
  content: {
    marginLeft: "10px",
  },
};

export default SukipiInfo;
