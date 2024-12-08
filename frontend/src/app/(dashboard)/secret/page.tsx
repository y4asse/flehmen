"use client";
import { Windows } from "@/components/common/Windows";
import React, { useState } from "react";
import InputVoice from "./_components/InputVoice";
import VoicePlay from "./_components/VoicePlay";

const Page = () => {
  const [isLoading, setIsLoading] = React.useState(true);
  const handleSetLoading = () => {
    setIsLoading(true);
  };
  const handleStopLoading = () => {
    setIsLoading(false);
  };
  const [url, setUrl] = useState<string | null>(null);

  const windows = [
    {
      ...timeGraphWindow,
      title: "聞く❤︎",
      children: <VoicePlay isLoading={isLoading} url={url} />,
    },
    {
      ...weekGraphWindow,
      title: "作る❤︎",
      children: (
        <InputVoice
          handleSetLoading={handleSetLoading}
          handleStopLoading={handleStopLoading}
          setUrl={setUrl}
        />
      ),
    },
  ];
  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

const timeGraphWindow = {
  initSize: {
    width: 800,
    height: 530,
  },
  initPosition: {
    x: 170,
    y: 80,
    z: 1,
  },
};
const weekGraphWindow = {
  initSize: {
    width: 430,
    height: 350,
  },
  initPosition: {
    x: 950,
    y: 50,
    z: 2,
  },
};

export default Page;
