"use client";
import React, { useRef, useState } from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";

type Props = {
  handleSetLoading: () => void;
  handleStopLoading: () => void;
  setUrl: React.Dispatch<React.SetStateAction<string | null>>
};

export const InputVoice = (props: Props) => {
  const { handleSetLoading, handleStopLoading, setUrl } = props;
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [text, setText] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
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
    setLoading(true);
    handleStartLoading();

    // FormData オブジェクトを作成
    const formData = new FormData();
    formData.append("voice", selectedFile); // "file" はサーバーで受け取るキー名に合わせる
    formData.append("text", text);

    // fetch リクエストを送信
    const response = await fetch(`http://localhost:8080/sukipi_voice`, {
      method: "POST",
      body: formData,
    });
    if (!response.ok) {
      console.error("Error:", response.statusText);
      setLoading(false);
      return;
    }
    const blob = await response.blob();
    console.log(blob);
    const url = URL.createObjectURL(blob);
    setUrl(url);
    stopLoading();
    setLoading(false);
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
          value={text}
          onChange={(event) => setText(event.target.value)}
        />
      </Flex>

      {/* 選択されたファイルの名前を表示 */}

      <p className="text-center mt-2">
        選択されたファイル:
        {selectedFile && <strong>{selectedFile.name}</strong>}
      </p>

      {/* 右下：作成ボタン */}
      <Button className="my-3 float-right" disabled={loading}>作成</Button>
    </form>
  );
};

export default InputVoice;
