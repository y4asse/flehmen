'use server'

import { SukipiInfo } from '@/types/sukipiInfo'
import { db } from '@/utils/db'
import { auth, clerkClient } from '@clerk/nextjs/server'
import { sql } from 'drizzle-orm'

export const completeOnboarding = async (sukipiInfo: SukipiInfo) => {
  const { userId } = await auth()

  if (!userId) {
    return { message: 'No Logged In User' }
  }

  const client = await clerkClient()
  try {
    const result = await db.execute(sql`
            INSERT INTO sukipis (name, user_id, x_id, liked_at, weight, height, mbti, birthday, hobby, shoes_size, family, near_station)
            VALUES (${sukipiInfo.name}, ${userId}, ${sukipiInfo.twitterId}, ${sukipiInfo.likedAt}, ${sukipiInfo.weight}, ${sukipiInfo.height}, ${sukipiInfo.mbti}, ${sukipiInfo.birthday}, ${sukipiInfo.hobby}, ${sukipiInfo.shoesSize}, ${sukipiInfo.family}, ${sukipiInfo.nearStation})
    `)
    console.log(result)
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