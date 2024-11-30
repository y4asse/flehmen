import { title } from "process";
import React from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";

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
    name: "shoeSize",
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
  }
]

const Page = () => {
  return (
    <Flex direction={"column"}>
    <Flex className="flex-col">
        {sukipiInput.map((input) => (
          <div key={input.name}>
            <label htmlFor={input.name}>{input.title}</label>
            <input type={input.type} name={input.name} id={input.name} />
          </div>
        ))}
      </Flex>
      <Button>登録</Button>
    </Flex>
  );
};


export default Page;
