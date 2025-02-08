"use client";
import { MobileWindows } from "@/components/common/MobileWindows";
import React, { useState } from "react";
import { Score } from "./_components/Score";
import FriendlyGraph from "./_components/FriendlyGraph";
import NextAction from "./_components/NextAction";
import { FileUploader } from "./_components/FileUploader";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Windows } from "@/components/common/Windows";

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
      children: <FriendlyGraph score={uploadResult?.score} isMobile />,
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

  const pcResultWindows = [
    {
      ...pcScoreWindow,
      children: <Score total={uploadResult?.score.total} />,
      title: "スコア",
    },
    {
      ...pcFriendlyGraphWindow,
      children: <FriendlyGraph score={uploadResult?.score} />,
      title: "分析結果",
    },
    {
      ...pcResultWindow,
      title: "次のアクション",
      children: (
        <NextAction
          text={uploadResult?.next_action}
          onClickBack={handleInitResult}
        />
      ),
    },
  ];

  const pcFileuploadWindows = [
    {
      ...pcFileuploadWindow,
      children: (
        <Flex direction={"column"} className="gap-4">
          <FileUploader handleFileChange={handleFileChange} />
          <Button
            onClick={handleUpload}
            disabled={uploading}
            className="text-white"
          >
            {uploading ? "アップロード中..." : "アップロード"}
          </Button>
          {message && <p className="mt-4 text-white">{message}</p>}
        </Flex>
      ),
    },
  ];

  return (
    <div>
      <div className="hidden md:block">
        {uploadResult ? (
          <Windows windows={pcResultWindows} />
        ) : (
          <Windows windows={pcFileuploadWindows} />
        )}
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
    width: 800,
    height: 160,
  },
  initPosition: {
    x: 900,
    y: 350,
    z: 3,
  },
};

const friendlyGraphWindow = {
  initSize: {
    width: 800,
    height: 180,
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
    height: 155,
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
    height: 250,
  },
  initPosition: {
    x: 140,
    y: 80,
    z: 1,
  },
};

const pcFileuploadWindow = {
  initSize: {
    width: 800,
    height: 500,
  },
  initPosition: {
    x: 350,
    y: 150,
    z: 1,
  },
};

const pcScoreWindow = {
  initSize: {
    width: 640,
    height: 350,
  },
  initPosition: {
    x: 750,
    y: 50,
    z: 3,
  },
};

const pcFriendlyGraphWindow = {
  initSize: {
    width: 600,
    height: 400,
  },
  initPosition: {
    x: 180,
    y: 140,
    z: 2,
  },
};

const pcResultWindow = {
  initSize: {
    width: 540,
    height: 280,
  },
  initPosition: {
    x: 750,
    y: 400,
    z: 1,
  },
};
export default Page;
