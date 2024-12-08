"use client";
import React, { useRef, useState } from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Input } from "@/components/ui/input";

export const InputVoice = () => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      const file = files[0];
      setSelectedFile(file);
      alert(`選択したファイル: ${file.name}`);
    }
  };

  const openFileSelector = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  return (
    <div className="text-[#FFF] w-full h-full relative p-4">
      {/* 上部：質問とサンプルボイス入力ボタン */}
      <Flex direction="row" align="center" justify="center" className="w-full mb-4">
        <span className="mr-4">今日は何で寝落ちする？</span>
        <Button onClick={openFileSelector}>サンプルボイス入力</Button>
        {/* ファイル選択用の非表示input */}
        <Input
          ref={fileInputRef}
          type="file"
          accept="audio/*"
          style={{ display: "none" }}
          onChange={handleFileChange}
        />
      </Flex>

      {/* テキスト入力フィールド */}
      <Flex direction="column" align="center" className="w-full">
        <Textarea
          placeholder="ここに文章を入力..."
          className={`w-[80%] h-[10vw] p-2 rounded-md border `}
        />
      </Flex>

      {/* 選択されたファイルの名前を表示 */}
      {selectedFile && (
        <p className="text-center mt-2">
          選択されたファイル: <strong>{selectedFile.name}</strong>
        </p>
      )}

      {/* 右下：作成ボタン */}
      <Button className="absolute bottom-4 right-4">作成</Button>
    </div>
  );
};

export default InputVoice;
