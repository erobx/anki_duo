import { relations } from 'drizzle-orm';
import { integer, text, sqliteTable, unique } from 'drizzle-orm/sqlite-core';

export const usersTable = sqliteTable('users', {
    id: integer('id').primaryKey({ autoIncrement: true }),
    name: text('name'),
});

export const usersWordsTable = sqliteTable('usersWords', {
    userId: integer('user_id').notNull().references(() => usersTable.id),
    wordId: text('word_id').notNull().references(() => wordsTable.id),
}, (t) => ({
    pk: unique().on(t.userId, t.wordId)
}));

export const wordsTable = sqliteTable('words', {
    id: text('id').primaryKey(),
    infinitive: text('infinitive'),
    normalized_string: text('normalized_word'),
    pos: text('pos'),
    word_string: text('word').notNull(),
    gender: text('gender'),
});

export const usersRelations = relations(usersTable, ({ many }) => ({
    words: many(usersWordsTable),
}));

export const wordsRelations = relations(wordsTable, ({ many }) => ({
    users: many(usersWordsTable),
}));

export const usersWordsRelations = relations(usersWordsTable, ({ one }) => ({
    users: one(usersTable, {
        fields: [usersWordsTable.userId],
        references: [usersTable.id]
    }),
    words: one(wordsTable, {
        fields: [usersWordsTable.wordId],
        references: [wordsTable.id]
    })
}));