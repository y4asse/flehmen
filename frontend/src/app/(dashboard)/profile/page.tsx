import { db } from "@/utils/db";
import SukipiProfile from "./_components/SukipiProfile";
import { sukipis, users } from "@/db/schema";
import { auth } from "@clerk/nextjs/server";
import { eq } from "drizzle-orm";

const Page = async () => {
  const { userId } = await auth();
  const user = await db.select().from(users).where(eq(users.clerkId, userId!));
  const sukipi = (await db.select().from(sukipis).where(eq(sukipis.userId, user[0].id))).map((sukipi) => ({
    name: sukipi.name,
    weight: sukipi.weight,
    height: sukipi.height,
    mbti: sukipi.mbti,
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
