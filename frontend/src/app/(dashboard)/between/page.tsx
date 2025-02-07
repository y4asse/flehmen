"use client";
import { MobileWindows } from "@/components/common/MobileWindows";
import React, { useState } from "react";
import { Score } from "./_components/Score";
import FriendlyGraph from "./_components/FriendlyGraph";
import NextAction from "./_components/NextAction";
import { FileUploader } from "./_components/FileUploader";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";

type UploadResult = {
  next_action: string;
  score: Score;
};

export type Score = {
  balance: number;
  rhythm: number;
  time: number;
  total: number;
  type: number;
  words: number;
};

const Page = () => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState("");
  const [uploadResult, setUploadResult] = useState<UploadResult | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setSelectedFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      setMessage("ファイルを選択してください");
      return;
    }

    setUploading(true);
    setMessage("");

    const formData = new FormData();
    formData.append("text", selectedFile);

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/between`,
        {
          method: "POST",
          body: formData,
        }
      );

      const result = (await response.json()) as UploadResult;

      if (result) {
        setMessage("アップロードが完了しました");
        setUploadResult(result);
      } else {
        setMessage("アップロードに失敗しました");
      }
    } catch (error) {
      console.error(error);
      setMessage("アップロードに失敗しました");
    } finally {
      setUploading(false);
    }
  };

  const handleInitResult = () => {
    setUploadResult(null);
    setSelectedFile(null);
  };

  const resultMobile = [
    {
      ...scoreWindow,
      children: <Score total={uploadResult?.score.total} />,
    },
    {
      ...friendlyGraphWindow,
      children: <FriendlyGraph score={uploadResult?.score} />,
    },
    {
      ...resultWindow,
      children: (
        <NextAction
          text={uploadResult?.next_action}
          onClickBack={handleInitResult}
        />
      ),
    },
  ];

  const fileuploadMobile = [
    {
      ...fileuploadWindow,
      children: <FileUploader handleFileChange={handleFileChange} />,
    },
  ];

  return (
    <div>
      <div className="hidden md:block">
        {/* <Windows windows={windows} /> */}
      </div>
      <div className="block md:hidden ">
        {uploadResult ? (
          <MobileWindows windows={resultMobile} />
        ) : (
          <Flex direction={"column"}>
            <MobileWindows windows={fileuploadMobile} />
            <Button
              onClick={handleUpload}
              disabled={uploading}
              className="text-white absolute bottom-[40vh]"
            >
              {uploading ? "アップロード中..." : "アップロード"}
            </Button>
            {message && (
              <p className="mt-4 text-white absolute bottom-[50vh]">
                {message}
              </p>
            )}
          </Flex>
        )}
      </div>
    </div>
  );
};

const scoreWindow = {
  initSize: {
    width: 460,
    height: 330,
  },
  initPosition: {
    x: 900,
    y: 350,
    z: 3,
  },
};

const friendlyGraphWindow = {
  initSize: {
    width: 430,
    height: 330,
  },
  initPosition: {
    x: 950,
    y: 50,
    z: 2,
  },
};

const resultWindow = {
  initSize: {
    width: 800,
    height: 530,
  },
  initPosition: {
    x: 140,
    y: 80,
    z: 1,
  },
};

const fileuploadWindow = {
  initSize: {
    width: 800,
    height: 800,
  },
  initPosition: {
    x: 140,
    y: 80,
    z: 1,
  },
};
export default Page;
