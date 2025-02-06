import { MBTI } from "./mbti"

export type SukipiInfo = {
  name: string;
  likedAt: Date | null;
  twitterId: string | null;
  weight: number | null;
  height: number | null;
  mbti: MBTI | null;
  birthday: Date | null;
  hobby: string | null;
  shoesSize: number | null;
  family: string | null;
  nearStation: string | null;
};
