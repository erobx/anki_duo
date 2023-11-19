import { drizzle } from 'drizzle-orm/libsql';
import { createClient } from '@libsql/client';
import { DATABASE_URL, AUTH_TOKEN } from '$env/static/private';
import * as schema from '../../lib/server/index';

const client = createClient({ url: DATABASE_URL, authToken: AUTH_TOKEN });
export const db = drizzle(client, { schema });