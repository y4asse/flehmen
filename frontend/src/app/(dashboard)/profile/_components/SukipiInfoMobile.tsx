import React, { useState } from "react";
import { SukipiInfo as SukipiInfoType } from "@/types/sukipiInfo";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import dayjs from "dayjs";

type Props = {
  sukipi: SukipiInfoType;
  onUpdate: (updatedSukipi: SukipiInfoType) => void;
};

export const SukipiInfoMobile = ({ sukipi, onUpdate }: Props) => {
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState<SukipiInfoType>(sukipi);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSave = () => {
    onUpdate(formData);
    setIsEditing(false);
  };

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
    {
      key: "hobby",
      title: "趣味",
      content: sukipi.hobby || "?まだわからない?",
    },
    {
      key: "shoesSize",
      title: "くつのサイズ",
      content: sukipi.shoesSize ? `${sukipi.shoesSize} cm` : "?まだわからない?",
    },
    {
      key: "family",
      title: "かぞく",
      content: sukipi.family || "?まだわからない?",
    },
    {
      key: "nearStation",
      title: "最寄駅",
      content: sukipi.nearStation || "?まだわからない?",
    },
  ];

  if (isEditing) {
    return (
      <div className="flex flex-col mt-12 w-4/5 mx-auto">
        <div className="grid grid-cols-2 mt-5 mb-5  gap-4 ">
          {infoItems.map((item) => (
            <div key={item.key} className="flex flex-col">
              <div className="text-white font-bold">{item.title}</div>
              <div className="text-white text-sm">
                元の値:{" "}
                <span className="opacity-70">
                  {typeof item.content === "string"
                    ? item.content
                    : dayjs(item.content).format("YYYY/MM/DD")}
                </span>
              </div>
              <Input
                name={item.key}
                value={(formData as never)[item.key] || ""}
                onChange={handleInputChange}
                className="text-white"
              />
            </div>
          ))}
        </div>
        <div className="flex justify-end mb-2 gap-4">
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
    <div className="flex items-center justify-center">
      <div className="flex flex-col gap-4 pt-10">
        {infoItems.map((item) => (
          <div key={item.key} className="grid grid-cols-[10rem_1fr] gap-2">
            <div className="text-white font-bold text-left">{item.title}</div>
            <div className="text-white text-left ">
              {typeof item.content === "string"
                ? item.content
                : dayjs(item.content).format("YYYY/MM/DD")}
            </div>
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
};

export default SukipiInfoMobile;
