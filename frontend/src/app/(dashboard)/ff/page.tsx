import React from "react";
import { FollowLog } from "./_components/FollowLog";
import { FollowerLog } from "./_components/FollowerLog";
import { UniversitySearch } from "./_components/UniversitySearch";
import { Windows } from "@/components/common/Windows";
import { MobileWindows } from "@/components/common/MobileWindows";

export type Follow = {
  id: string;
  name: string;
  followed_at: string;
  bio: string;
  icon: string;
};

export type Follower = {
  id: string;
  name: string;
  followed_at: string;
  bio: string;
  icon: string;
};

const Page = async (props: {
  searchParams?: Promise<{
    query?: string;
  }>;
}) => {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";

  const windows = [
    {
      ...followWindow,
      title: "あたらしいフォロー",
      children: <FollowLog followList={followList} />,
    },
    {
      ...followerWindow,
      title: "あたらしいフォロワー",
      children: <FollowerLog followerList={followerList} />,
    },
    {
      ...universityWindow,
      title: "大学どこだろ",
      children: <UniversitySearch query={query} />,
    },
  ];
  return (
    <div>
      <div className="hidden md:block">
        <Windows windows={windows} />;
      </div>

      <div className="block md:hidden ">
        <MobileWindows windows={windows} />
      </div>
    </div>
  );
};

const followerList: Follower[] = [
  {
    id: "1",
    name: "かりんたん",
    followed_at: "2024-12-06:09:40:21",
    bio: "2ch/カフェ巡り☕️",
    icon: "/images/karin.png",
  },
  {
    id: "2",
    name: "ささandりょう",
    followed_at: "2024-12-05:02:30:11",
    bio: "YNU 26卒",
    icon: "https://picsum.photos/400",
  },
  {
    id: "3",
    name: "まえひろ",
    followed_at: "2024-12-03:12:30:11",
    bio: "UT",
    icon: "https://picsum.photos/600",
  },
  {
    id: "4",
    name: "kipeo",
    followed_at: "2024-12-01:11:00:00",
    bio: "動物と話せるペキ！",
    icon: "https://picsum.photos/150",
  },
  {
    id: "5",
    name: "ごろう",
    followed_at: "2024-11-05:13:56:42",
    bio: "ぼくはごろう",
    icon: "https://picsum.photos/300",
  },
];

const followList: Follow[] = [
  {
    id: "1",
    name: "sap",
    followed_at: "2024-12-06:01:30:11",
    bio: "thanks",
    icon: "https://picsum.photos/100",
  },
  {
    id: "2",
    name: "ちょまる",
    followed_at: "2023-11-02:12:24:33",
    bio: "ぼくはちょまる",
    icon: "https://picsum.photos/250",
  },
  {
    id: "3",
    name: "ドメイン工藤",
    followed_at: "2024-11-01:12:30:11",
    bio: "バーロー Run",
    icon: "https://picsum.photos/350",
  },
  {
    id: "4",
    name: "しろう",
    followed_at: "2024-10-05:08:20:42",
    bio: "ぼくはしろう",
    icon: "https://picsum.photos/550",
  },
  {
    id: "5",
    name: "ごろう",
    followed_at: "2021-11-05",
    bio: "ぼくはごろう",
    icon: "https://picsum.photos/200",
  },
];

const followerWindow = {
  initSize: { width: 600, height: 400 },
  initPosition: { x: 200, y: 130, z: 1 },
};

const followWindow = {
  initSize: { width: 700, height: 400 },
  initPosition: { x: 800, y: 50, z: 2 },
};

const universityWindow = {
  initSize: { width: 500, height: 300 },
  initPosition: { x: 700, y: 450, z: 3 },
};
export default Page;
