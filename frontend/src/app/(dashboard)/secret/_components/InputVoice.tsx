"use client";
import React from 'react';
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";

export const InputVoice = ({  }) => {  // isActiveを受け取る
  return (
    <div className="text-[#FFF] w-full h-full relative p-4">
      {/* 上部：質問とサンプルボイス入力ボタン */}
      <Flex direction="row" align="center" justify="center" className="w-full mb-4">
        <span className="mr-4">今日は何で寝落ちする？</span>
        <Button >
          サンプルボイス入力
        </Button>
      </Flex>
      
      {/* テキスト入力フィールド */}
      <Flex direction="column" align="center" className="w-full">
        <Textarea
          placeholder="ここに文章を入力..." 
          className={`w-[80%] h-[10vw] p-2 rounded-md border `}
        />
      </Flex>

      {/* 右下：作成ボタン */}
      <Button className="absolute bottom-4 right-4">
        作成
      </Button>
    </div>
  );
};
