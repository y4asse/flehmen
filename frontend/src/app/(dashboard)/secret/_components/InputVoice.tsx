"use client";
import React, { useRef, useState } from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";

type Props = {
  handleSetLoading: () => void;
  handleStopLoading: () => void;
};

export const InputVoice = (props: Props) => {
  const { handleSetLoading, handleStopLoading } = props;
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleStartLoading = () => {
    if (selectedFile) {
      handleSetLoading();
    }
  };

  const stopLoading = () => {
    setSelectedFile(null);
    handleStopLoading();
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files && files.length > 0) {
      const file = files[0];
      setSelectedFile(file);
    }
  };

  const openFileSelector = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const onSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!selectedFile) {
      return;
    }
    handleStartLoading();
    setTimeout(() => {
      // ローディング終了処理
      stopLoading(); // setIsLoading(false)を含む関数
    }, 5000);

    // FormData オブジェクトを作成
    // const formData = new FormData();
    // formData.append("voice", selectedFile); // "file" はサーバーで受け取るキー名に合わせる

    // fetch リクエストを送信
    // const response = await fetch(`http://localhost:8080/sukipi_voice`, {
    //   method: "POST",
    //   body: formData,
    // });

    // const result = await response.json();
    // console.log(result);
  };

  return (
    <form
      className="text-[#FFF] w-full h-full relative p-4"
      method="POST"
      encType="multipart/form-data"
      onSubmit={onSubmit}
    >
      {/* 上部：質問とサンプルボイス入力ボタン */}
      <Flex
        direction="row"
        align="center"
        justify="center"
        className="w-full mb-4"
      >
        <span className="mr-4">今日は何で寝落ちする？</span>
        <Button onClick={openFileSelector}>サンプルボイス入力</Button>
        {/* ファイル選択用の非表示input */}
        <input
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

      <p className="text-center mt-2">
        選択されたファイル:
        {selectedFile && <strong>{selectedFile.name}</strong>}
      </p>

      {/* 右下：作成ボタン */}
      <Button className="my-3 float-right">作成</Button>
    </form>
  );
};

export default InputVoice;
