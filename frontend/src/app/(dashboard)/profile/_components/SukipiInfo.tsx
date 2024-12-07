import React, { useState } from "react";
import { Sukipi } from "../page";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

type Props = {
  sukipi: Sukipi;
  onUpdate: (updatedSukipi: Sukipi) => void;
};

<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
  const infoItems = [
    { key: "name", title: "なまえ", content: sukipi.name },
    {
      key: "height",
      title: "身長",
      content: sukipi.height ? `${sukipi.height} cm` : "?まだわからない?",
    },
    {
      key: "weight",
      title: "体重",
      content: sukipi.weight ? `${sukipi.weight} kg` : "?まだわからない?",
    },
    { key: "mbti", title: "MBTI", content: sukipi.mbti || "?まだわからない?" },
    {
      key: "birthday",
      title: "誕生日",
      content: sukipi.birthday || "?まだわからない?",
    },
<<<<<<< Updated upstream
    { key: "hobby", title: "趣味", content: sukipi.hobby || "?まだわからない?" },
=======
    {
      key: "hobby",
      title: "趣味",
      content: sukipi.hobby || "?まだわからない?",
    },
>>>>>>> Stashed changes
    {
      key: "shoesSize",
      title: "くつのサイズ",
      content: sukipi.shoesSize ? `${sukipi.shoesSize} cm` : "?まだわからない?",
    },
<<<<<<< Updated upstream
    { key: "family", title: "かぞく", content: sukipi.family || "?まだわからない?" },
=======
    {
      key: "family",
      title: "かぞく",
      content: sukipi.family || "?まだわからない?",
    },
>>>>>>> Stashed changes
    {
      key: "nearlyStation",
      title: "最寄駅",
      content: sukipi.nearlyStation || "?まだわからない?",
    },
  ];

  if (isEditing) {
    return (
      <div className="flex flex-col p-4 pt-10 w-4/5 mx-auto mt-[10rem]">
        <div className="grid grid-cols-2 gap-4 mb-6">
          {infoItems.map((item) => (
            <div key={item.key} className="flex flex-col">
              <div className="text-white font-bold">{item.title}</div>
              <div className="text-white text-sm mb-2">
                元の値: <span className="opacity-70">{item.content}</span>
              </div>
              <Input
                name={item.key}
<<<<<<< Updated upstream
                value={(formData as any)[item.key] || ""}
=======
                value={(formData as never)[item.key] || ""}
>>>>>>> Stashed changes
                onChange={handleInputChange}
                className="text-white"
              />
            </div>
          ))}
        </div>
        <div className="flex justify-end gap-4">
          <Button
            className="text-white border border-white hover:bg-white hover:text-black"
            onClick={() => setIsEditing(false)}
          >
            もどる
          </Button>
          <Button
            className="text-white border border-white hover:bg-white hover:text-black"
            onClick={handleSave}
          >
            保存
          </Button>
        </div>
      </div>
    );
  }


  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="flex flex-col gap-4 p-4 pt-10 w-4/5 max-w-lg mx-auto">
        {infoItems.map((item) => (
<<<<<<< Updated upstream
          <div 
            key={item.key} 
            className="grid grid-cols-[6rem_1fr] gap-2"
          >
=======
          <div key={item.key} className="grid grid-cols-[6rem_1fr] gap-2">
>>>>>>> Stashed changes
            <div className="text-white font-bold text-left">{item.title}</div>
            <div className="text-white text-left w-40">{item.content}</div>
          </div>
        ))}
        <div className="flex justify-end">
          <Button
            className="text-white border border-white hover:bg-white hover:text-black"
            onClick={() => setIsEditing(true)}
          >
            編集
          </Button>
        </div>
      </div>
    </div>
  );
<<<<<<< Updated upstream
  
=======
>>>>>>> Stashed changes
};

export default SukipiInfo;
