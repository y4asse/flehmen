"use client";
import React, { useState } from "react";
import { Flex } from "@/components/ui/flex";
import Typewriter from "@/components/common/Typewriter";

const sukipiInput = [

  {
    title: "名前",
    type: "text",
    name: "name",
  },
  {
    title: "体重",
    type: "number",
    name: "weight",
  },
  {
    title: "身長",
    type: "number",
    name: "height",
  },
  {
    title: "MBTI",
    type: "text",
    name: "mbti",
  },
  {
    title: "誕生日",
    type: "date",
    name: "birthday",
  },
  {
    title: "趣味",
    type: "text",
    name: "hobby",
  },
  {
    title: "靴のサイズ",
    type: "number",
    name: "shoesSize",
  },
  {
    title: "家族",
    type: "text",
    name: "family",
  },
  {
    title: "最寄り駅",
    type: "text",
    name: "nearStation",
  },
];

const Page = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [responses, setResponses] = useState(Array(sukipiInput.length).fill(""));
  const [topIndex, setTopIndex] = useState(0);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newResponses = [...responses];
    newResponses[currentIndex] = e.target.value;
    setResponses(newResponses);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && responses[currentIndex] !== "") {
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

  return (

    <Flex className="w-[100%] h-[100vh] bg-[#000]" style={{ color: "#E4007F" }} direction={"column"} align={"center"} justify={"center"}>
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
      <div className="input-section">
        {sukipiInput.map((input, index) => (
      <div
        key={input.name}
        style={{
          display: index > topIndex ? "none" : "flex", // Flexで横並びに
          alignItems: "center",  // ラベルと入力を中央揃えに
          marginBottom: "5px",   // 質問間の余白を調整
        }}
        onClick={() => handleClick(index)}
      >
        <label htmlFor={input.name} style={{ marginRight: "8px", whiteSpace: "nowrap" }}>
          {input.title}:
        </label>
        {index === currentIndex ? (
          <input
          type={input.type}
          name={input.name}
          id={input.name}
          value={responses[index]}
          onChange={handleChange}
          onKeyDown={handleKeyDown}
          autoFocus
          placeholder={input.name === "birthday" ? "YYYY-MM-DD" : ""} // 誕生日のヒント
          className="bg-transparent text-[#E4007F] border-none outline-none flex-[0_0_60%] max-w-[300px] placeholder-[#E4007F] placeholder-opacity-50"
        />
        
          ) : (
            <div style={{ flex: "0 0 60%" }}>{responses[index]}</div>
          )}
            </div>
          ))}
          </div>
        </Flex>
      </div>
    </Flex>
  );
};

export default Page;
