import { drizzle } from 'drizzle-orm/libsql';
import { migrate } from 'drizzle-orm/libsql/migrator';
import { createClient } from '@libsql/client';
import * as dotenv from 'dotenv';
dotenv.config({ path: './.env' });

export const client = createClient({
    url: process.env.DATABASE_URL as string,
    authToken: process.env.AUTH_TOKEN as string
});

export const db = drizzle(client);

async function main() {
    try {
        await migrate(db, {
            migrationsFolder: 'drizzle/migrations'
        });
        console.log("Tables migrated!");
        process.exit(0);
    } catch (error) {
        console.error("Error performing migration", error);
        process.exit(1);
    }
}

main();