"use client";
import { Windows } from "@/components/common/Windows";
import React, { useState } from "react";
import { SukipiInfo } from "./SukipiInfo";
import { SukipiLikedAt } from "./SukipiLikedAt";
import { SukipiInfo as SukipiInfoType } from "@/types/sukipiInfo";
import { MobileWindows } from "@/components/common/MobileWindows";
import { SukipiInfoMobile } from "./SukipiInfoMobile";

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

  const mobile = [
    {
      ...sukipiInfoMobile,
      children: (
        <SukipiInfoMobile sukipi={sukipiData} onUpdate={handleUpdate} />
      ),
    },
  ];

  return (
    <div>
      <div className="hidden md:block">
        <Windows windows={windows} />
      </div>
      <div className="block md:hidden " style={{ fontSize: "21px" }}>
        <MobileWindows windows={mobile} />
      </div>
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

const sukipiInfoMobile = {
  initSize: {
    width: 750,
    height: 800,
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
