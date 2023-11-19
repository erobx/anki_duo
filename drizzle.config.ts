import type { Config } from 'drizzle-kit';
import * as dotenv from 'dotenv';
dotenv.config({ path: ".env" });

export default {
    schema: './drizzle/schema.ts',
    out: './drizzle/migrations',
    driver: 'turso',
    dbCredentials: {
        url: process.env.DATABASE_URL as string,
        authToken: process.env.AUTH_TOKEN as string
    }
} satisfies Config;