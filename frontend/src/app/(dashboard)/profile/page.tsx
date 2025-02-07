import { db } from "@/utils/db";
import SukipiProfile from "./_components/SukipiProfile";
import { sukipis, users } from "@/db/schema";
import { auth } from "@clerk/nextjs/server";
import { eq } from "drizzle-orm";
import { MBTI } from "@/types/mbti";

const Page = async () => {
  const { userId } = await auth();
  if (!userId) {
    return <div>ログインしてください</div>;
  }
  const user = await db.select().from(users).where(eq(users.clerkId, userId)).catch((error) => {
    console.error(error);
    return null;
  });
  if (!user) {
    return <div>エラーが発生しました</div>;
  }
  if (user.length === 0) {
    return <div>ユーザーが見つかりません</div>;
  }
  const sukipi = (await db.select().from(sukipis).where(eq(sukipis.userId, user[0].id))).map((sukipi) => ({
    name: sukipi.name,
    weight: sukipi.weight,
    height: sukipi.height,
    mbti: sukipi.mbti as MBTI,
    birthday: sukipi.birthday,
    hobby: sukipi.hobby,
    shoesSize: sukipi.shoesSize,
    family: sukipi.family,
    nearStation: sukipi.nearlyStation,
    twitterId: sukipi.xId,
    likedAt: sukipi.likedAt,
  }));
  return <SukipiProfile sukipi={sukipi[0]} />;
};

export default Page;
