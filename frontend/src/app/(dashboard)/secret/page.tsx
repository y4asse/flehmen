import { Windows } from "@/components/common/Windows";
import React from "react";
import  InputVoice  from "./_components/InputVoice";
import  VoicePlay  from "./_components/VoicePlay";



const Page = () => {
  
  const windows = [
    {
      ...timeGraphWindow,
      title: "聞く❤︎",
      children:<VoicePlay />,
    },
    {
      ...weekGraphWindow,
      title: "作る❤︎",
      children: <InputVoice />,
    },
  ];
  return (
    <div>
      <Windows windows={windows} />
    </div>
  )
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
    height: 330,
  },
  initPosition: {
    x: 950,
    y: 50,
    z: 2,
  },
};

export default Page;
