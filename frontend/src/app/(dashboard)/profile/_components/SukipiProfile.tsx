"use client";
import { Windows } from "@/components/common/Windows";
import React, { useState } from "react";
import { SukipiInfo } from "./SukipiInfo";
import { SukipiLikedAt } from "./SukipiLikedAt";
import { SukipiInfo as SukipiInfoType } from "@/types/sukipiInfo";

const SukipiProfile = ({ sukipi }: { sukipi: SukipiInfoType }) => {
  const [sukipiData, setSukipiData] = useState<SukipiInfoType>(sukipi);

  const handleUpdate = (updatedSukipi: SukipiInfoType) => {
    setSukipiData(updatedSukipi);
  };

  const windows = [
    {
      ...sukipiInfoWindow,
      title: `${sukipi.name}くんのこと`,
      children: <SukipiInfo sukipi={sukipiData} onUpdate={handleUpdate} />,
    },
    {
      ...sukipiLikeAtWindow,
      title: "好きになって何日？",
      children: <SukipiLikedAt likedAt={sukipi.likedAt} />,
    },
  ];

  return (
    <div>
      <Windows windows={windows} />
    </div>
  );
};

const sukipiInfoWindow = {
  initSize: {
    width: 750,
    height: 500,
  },
  initPosition: {
    x: 650,
    y: 50,
    z: 2,
  },
};

const sukipiLikeAtWindow = {
  initSize: {
    width: 500,
    height: 400,
  },
  initPosition: {
    x: 300,
    y: 250,
    z: 1,
  },
};

export default SukipiProfile;
