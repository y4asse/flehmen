// formひどいので後で直す
"use client";
import React, { useState } from "react";
import { Flex } from "@/components/ui/flex";
import Typewriter from "@/components/common/Typewriter";
import { useRouter } from "next/navigation";
import { fetchPost } from "@/lib/fetcher";

type SukipiInput = {
  title: string;
  type: string;
  name: string;
  placeholder: string;
  required: boolean;
};
const sukipiInput: SukipiInput[] = [
  {
    title: "名前",
    type: "text",
    name: "name",
    placeholder: "",
    required: true,
  },
  {
    title: "TwitterID",
    type: "text",
    name: "twitterId",
    placeholder: "",
    required: true,
  },
  {
    title: "好きになった日",
    type: "text",
    name: "likedAt",
    placeholder: "YYYY-MM-DD",
    required: true,
  },
  {
    title: "体重",
    type: "number",
    name: "weight",
    placeholder: "kg",
    required: false,
  },
  {
    title: "身長",
    type: "number",
    name: "height",
    placeholder: "cm",
    required: false,
  },
  {
    title: "MBTI",
    type: "text",
    name: "mbti",
    placeholder: "",
    required: false,
  },
  {
    title: "誕生日",
    type: "text",
    name: "birthday",
    placeholder: "YYYY-MM-DD",
    required: false,
  },
  {
    title: "趣味",
    type: "text",
    name: "hobby",
    placeholder: "",
    required: false,
  },
  {
    title: "靴のサイズ",
    type: "number",
    name: "shoesSize",
    placeholder: "cm",
    required: false,
  },
  {
    title: "家族構成",
    type: "text",
    name: "family",
    placeholder: "",
    required: false,
  },
  {
    title: "最寄り駅",
    type: "text",
    name: "nearStation",
    placeholder: "",
    required: false,
  },
];
type RequestBody = {
  name: string;
  twitterId: string;
  likedAt: string;
  weight: number | null;
  height: number | null;
  mbti: string | null;
  birthday: Date | null;
  hobby: string | null;
  shoesSize: number | null;
  family: string | null;
  nearStation: string | null;
};

const Page = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [responses, setResponses] = useState(
    Array(sukipiInput.length).fill("")
  );
  const [topIndex, setTopIndex] = useState(0);
  const [checkMessage, setCheckMessage] = useState("Press Enter");
  const router = useRouter();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newResponses = [...responses]; // 配列をコピー
    newResponses[currentIndex] = e.target.value; // 現在のインデックスだけを更新
    setResponses(newResponses); // 状態を更新
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      const nextIndex = currentIndex + 1;
      if (nextIndex < sukipiInput.length) {
        setCurrentIndex(nextIndex);
        if (nextIndex > topIndex) setTopIndex(nextIndex);
      }
    }
    if (e.key === "ArrowUp" && currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
    if (e.key === "ArrowDown" && currentIndex < sukipiInput.length - 1) {
      setCurrentIndex(currentIndex + 1);
    }
  };

  const handleClick = (index: number) => {
    setCurrentIndex(index);
    if (index > topIndex) setTopIndex(index);
  };

  const handleFocus = () => {
    if (responses[0] === "" || responses[1] === "" || responses[2] === "") {
      setCheckMessage("名前とTwitterと好きになった日は知ってるよね");
      return;
    }
    setCheckMessage("Press Enter");
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (currentIndex != sukipiInput.length - 1) {
      return;
    }
    if (responses[0] === "" || responses[1] === "" || responses[2] === "") {
      return;
    }
    const data = sukipiInput.reduce(
      (acc: { [key: string]: string | number | Date | null }, input, index) => {
        const value = responses[index];

        if (value === "") {
          // 空文字列は null に変換
          acc[input.name] = null;
        } else if (input.placeholder === "YYYY-MM-DD" && value) {
          // "birthday" は Date 型に変換
          acc[input.name] = new Date(value);
        } else if (!isNaN(Number(value)) && value.trim() !== "") {
          // 数値として変換可能な場合は Number 型に変換
          acc[input.name] = Number(value);
        } else {
          // それ以外はそのまま文字列として保存
          acc[input.name] = value;
        }

        return acc;
      },
      {} as RequestBody // 型定義を更新
    );
    await fetchPost("/sukipi", data);
    router.push("/loading");
  };

  return (
    <Flex
      className="w-[100%] h-[100vh] bg-[#000]"
      style={{ color: "#E4007F" }}
      direction={"column"}
      align={"center"}
      justify={"center"}
    >
      <div className="text-[5vw] mb-[5%]">
        <Typewriter text="スキピのこと教えて？" speed={100} loop={false} />
      </div>
      <div
        className="console-container"
        style={{
          backgroundColor: "#000",
          color: "#E4007F",
          padding: "20px",
          width: "80%",
          maxWidth: "800px",
          height: "400px", // 高さを固定
          border: "2px solid #E4007F",
          borderRadius: "10px",
          boxShadow: "0 0 20px rgba(228, 0, 127, 0.7)",
          overflowY: "auto", // 縦スクロールを有効に
        }}
      >
        <Flex direction={"column"} align={"start"}>
          <form className="input-section" method="post" onSubmit={handleSubmit}>
            {sukipiInput.map((input, index) => (
              <div
                key={input.name}
                style={{
                  display: index > topIndex ? "none" : "flex", // Flexで横並びに
                  alignItems: "center", // ラベルと入力を中央揃えに
                  marginBottom: "5px", // 質問間の余白を調整
                }}
                onClick={() => handleClick(index)}
              >
                <label
                  htmlFor={input.name}
                  style={{ marginRight: "8px", whiteSpace: "nowrap" }}
                >
                  {input.required ? "★" : ""} {input.title}:
                </label>
                {index === currentIndex ? (
                  <input
                    type={input.type}
                    name={input.name}
                    id={input.name}
                    value={responses[index] || ""}
                    onChange={handleChange}
                    onKeyDown={handleKeyDown}
                    autoFocus
                    placeholder={input.placeholder}
                    className="bg-transparent text-[#E4007F] border-none outline-none flex-[0_0_60%] max-w-[300px] placeholder-[#E4007F] placeholder-opacity-50"
                    onFocus={handleFocus}
                  />
                ) : (
                  <div style={{ flex: "0 0 60%" }}>{responses[index]}</div>
                )}
              </div>
            ))}
            {currentIndex === sukipiInput.length - 1 && <p>{checkMessage}</p>}
          </form>
        </Flex>
      </div>
    </Flex>
  );
};

export default Page;
