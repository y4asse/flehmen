import { int, mysqlTable, varchar, timestamp, boolean, serial, datetime } from 'drizzle-orm/mysql-core';
import { relations } from "drizzle-orm";

// User Table
export const users = mysqlTable("users", {
  id: serial("id").primaryKey(),
  clerkId: varchar("clerk_id", { length: 255 }).notNull(),
  name: varchar("name", { length: 255 }).notNull(),
  weight: int("weight"),
  height: int("height"),
  isMale: boolean("is_male").notNull(),
  createdAt: timestamp("created_at").defaultNow().notNull(),
});

// Sukipi Table
export const sukipis = mysqlTable("sukipis", {
  id: serial("id").primaryKey(),
  name: varchar("name", { length: 255 }).notNull(),
  likedAt: datetime("liked_at"),
  weight: int("weight"),
  height: int("height"),
  xId: varchar("x_id", { length: 255 }),
  hobby: varchar("hobby", { length: 255 }),
  birthday: datetime("birthday"),
  shoesSize: int("shoes_size"),
  family: varchar("family", { length: 255 }),
  nearlyStation: varchar("nearly_station", { length: 255 }),
  mbti: varchar("mbti", { length: 255 }),
  createdAt: timestamp("created_at").notNull().$defaultFn(() => new Date()),
  userId: serial("sukipi_user").references(() => users.id, { onDelete: "cascade" }).unique(),
});

// Relations
export const usersRelations = relations(users, ({ one }) => ({
  sukipi: one(sukipis, {
    fields: [users.id],
    references: [sukipis.userId],
  }),
}));

export const sukipisRelations = relations(sukipis, ({ one }) => ({
  user: one(users, {
    fields: [sukipis.userId],
    references: [users.id],
  }),
}));
