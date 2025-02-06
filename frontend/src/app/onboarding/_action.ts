'use server'

import { sukipis, users } from '@/db/schema'
import { SukipiInfo } from '@/types/sukipiInfo'
import { db } from '@/utils/db'
import { auth, clerkClient } from '@clerk/nextjs/server'
import { eq } from 'drizzle-orm'

export const completeOnboarding = async (sukipiInfo: SukipiInfo) => {
  const { userId } = await auth()

  if (!userId) {
    return { message: 'No Logged In User' }
  }

  const client = await clerkClient()
  try {
    const user = await db.select().from(users).where(eq(users.clerkId, userId))
    if (!user.length) {
      return { error: 'User Not Found', message: null }
    }
    await db.insert(sukipis).values({
      name: sukipiInfo.name,
      likedAt: sukipiInfo.likedAt || new Date(),
      userId: user[0].id,
      xId: sukipiInfo.twitterId,
      weight: sukipiInfo.weight,
      height: sukipiInfo.height,
      mbti: sukipiInfo.mbti,
      birthday: sukipiInfo.birthday,
      hobby: sukipiInfo.hobby,
      shoesSize: sukipiInfo.shoesSize,
      family: sukipiInfo.family,
    })
    const res = await client.users.updateUser(userId, {
        publicMetadata: {
            onboardingComplete: true,
        },
    })
    return { message: res.publicMetadata, error: null }
  } catch (err) {
    console.error(err)
    return { error: 'There was an error updating the user metadata.', message: null }
  }
}