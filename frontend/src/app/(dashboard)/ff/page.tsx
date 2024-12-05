import React from "react";
import { FollowLog } from "./_components/FollowLog";
import { FollowerLog } from "./_components/FollowerLog";
import { UniversitySearch } from "./_components/UniversitySearch";
import { Windows } from "@/components/common/Windows";

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
      ...universityWindow,
      title: "大学どこだろ",
      children: <UniversitySearch query={query} />,
    },
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
  ];
  return <Windows windows={windows} />;
};

const followerList: Follower[] = [
  {
    id: "1",
    name: "たろう",
    followed_at: "2021-11-01",
    bio: "ぼくはたろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "2",
    name: "じろう",
    followed_at: "2021-11-02",
    bio: "ぼくはじろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "3",
    name: "さぶろう",
    followed_at: "2021-11-03",
    bio: "ぼくはさぶろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "4",
    name: "しろう",
    followed_at: "2021-11-04",
    bio: "ぼくはしろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "5",
    name: "ごろう",
    followed_at: "2021-11-05",
    bio: "ぼくはごろう",
    icon: "https://picsum.photos/200",
  },
];

const followList: Follow[] = [
  {
    id: "1",
    name: "たろう",
    followed_at: "2021-11-01",
    bio: "ぼくはたろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "2",
    name: "じろう",
    followed_at: "2021-11-02",
    bio: "ぼくはじろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "3",
    name: "さぶろう",
    followed_at: "2021-11-03",
    bio: "ぼくはさぶろう",
    icon: "https://picsum.photos/200",
  },
  {
    id: "4",
    name: "しろう",
    followed_at: "2021-11-04",
    bio: "ぼくはしろう",
    icon: "https://picsum.photos/200",
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
  initSize: { width: 550, height: 350 },
  initPosition: { x: 200, y: 100, z: 1 },
};

const followWindow = {
  initSize: { width: 600, height: 350 },
  initPosition: { x: 800, y: 55, z: 2 },
};

const universityWindow = {
  initSize: { width: 500, height: 300 },
  initPosition: { x: 780, y: 350, z: 3 },
};
export default Page;
