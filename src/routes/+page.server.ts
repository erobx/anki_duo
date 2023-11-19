import { db } from '$lib/server/db';
import { usersTable, usersWordsTable, wordsTable } from '../../drizzle/schema';

export const actions = {
    createUser: async () => {
        type NewUser = typeof usersTable.$inferInsert;

        const insertUser = async (user: NewUser) => {
            return db.insert(usersTable).values(user);
        }

        const newUser: NewUser = { name: "Evan" };
        await insertUser(newUser);

        return { success: true };
    },
    createWord: async () => {
        type NewWord = typeof wordsTable.$inferInsert;

        const insertWord = async (word: NewWord) => {
            return db.insert(wordsTable).values(word);
        }

        const newWord: NewWord = {
            id: "poner",
            infinitive: "poner",
            normalized_string: "poner",
            pos: null,
            word_string: "poner",
            gender: null
        };
        await insertWord(newWord);

        return { success: true };
    },
    createManyToMany: async () => {
        const newUsers = await db
            .insert(usersTable)
            .values([
                {
                    name: "Test",
                },
                {
                    name: "Test2",
                }
            ])
            .returning({userId: usersTable.id});
        
        // const userId = newUsers[0].userId;

        const newWords = await db.insert(wordsTable).values([
            {
                id: "7e89asdf",
                infinitive: "saber",
                normalized_string: "saber",
                pos: undefined,
                word_string: "saber",
                gender: undefined,
            },
            {
                id: "eyJ23Lad",
                infinitive: "haber",
                normalized_string: "haber",
                pos: undefined,
                word_string: "haber",
                gender: undefined,
            }
        ]).returning({wordId: wordsTable.id});

        await db.insert(usersWordsTable).values([
            {
                userId: newUsers[0].userId,
                wordId: newWords[0].wordId
            },
            {
                userId: newUsers[0].userId,
                wordId: newWords[1].wordId
            },
            {
                userId: newUsers[1].userId,
                wordId: newWords[0].wordId
            }
        ]).execute();

        return { success: true };
    }
};