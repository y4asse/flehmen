import React, { useState } from "react";
import { Sukipi } from "../page";
import { Button } from "@/components/ui/button";

type Props = {
  sukipi: Sukipi;
  onUpdate: (updatedSukipi: Sukipi) => void;
};

// メインコンポーネント
export const SukipiInfo = ({ sukipi, onUpdate }: Props) => {
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState<Sukipi>(sukipi);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSave = () => {
    onUpdate(formData);
    setIsEditing(false);
  };

  console.log(formData);

  // Sukipiオブジェクトから情報もらう
  const infoItems = [
    { title: "なまえ", content: sukipi.name },
    {
      title: "身長",
      content: sukipi.height ? `${sukipi.height} cm` : "?まだわからない?",
    },
    {
      title: "体重",
      content: sukipi.weight ? `${sukipi.weight} kg` : "?まだわからない?",
    },
    { title: "MBTI", content: sukipi.mbti || "?まだわからない?" },
    { title: "誕生日", content: sukipi.birthday || "?まだわからない?" },
    { title: "趣味", content: sukipi.hobby || "?まだわからない?" },
    {
      title: "くつのサイズ",
      content: sukipi.shoesSize ? `${sukipi.shoesSize} cm` : "?まだわからない?",
    },
    { title: "かぞく", content: sukipi.family || "?まだわからない?" },
    { title: "最寄駅", content: sukipi.nearlyStation || "?まだわからない?" },
  ];

  if (isEditing) {
    return (
      <div>
        <div style={styles.editor}>
          <div>
            {infoItems.map((item) => (
              <div key={item.title} style={styles.title}>
                {item.title}
              </div>
            ))}
          </div>
          <div>
            {Object.keys(formData).map((key) => (
              <div key={key} style={styles.content}>
                <input
                  type="text"
                  name={key}
                  value={(formData as any)[key] || ""}
                  onChange={handleInputChange}
                  style={styles.editInput}
                />
              </div>
            ))}
          </div>
        </div>
        <button onClick={() => setIsEditing(false)} style={styles.cancel}>
          もどる
        </button>
        <Button onClick={handleSave}>保存</Button>
      </div>
    );
  }

  return (
    <div>
      {infoItems.map((item) => (
        <InfoItem key={item.title} title={item.title} content={item.content} />
      ))}
      <Button onClick={() => setIsEditing(true)}>編集</Button>
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
  },
  editor: {
    display: "flex",
    margin: "10px 0",
    fontSize: "0.8em",
    fontWeight: "bold",
    color: "white",
  },
  title: {
    width: "100px",
  },
  content: {
    marginLeft: "10px",
  },
  cancel: {
    color: "white",
    backgroundColor: "black",
  },
  editInput: {
    background: "transparent",
  },
};

export default SukipiInfo;
